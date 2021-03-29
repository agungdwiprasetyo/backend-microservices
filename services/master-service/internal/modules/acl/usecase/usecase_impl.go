// Code generated by candi v1.4.0.

package usecase

import (
	"context"
	"errors"
	"fmt"

	"monorepo/sdk"
	"monorepo/services/master-service/internal/modules/acl/domain"
	appsdomain "monorepo/services/master-service/internal/modules/apps/domain"
	shareddomain "monorepo/services/master-service/pkg/shared/domain"
	"monorepo/services/master-service/pkg/shared/repository"

	"pkg.agungdp.dev/candi/candishared"
	"pkg.agungdp.dev/candi/codebase/factory/dependency"
	"pkg.agungdp.dev/candi/codebase/interfaces"
	"pkg.agungdp.dev/candi/tracer"
)

type aclUsecaseImpl struct {
	cache     interfaces.Cache
	repoMongo *repository.RepoMongo
	sdk       sdk.SDK
	// commonUC  common.Usecase
}

// NewACLUsecase usecase impl constructor
func NewACLUsecase(deps dependency.Dependency) ACLUsecase {
	return &aclUsecaseImpl{
		cache:     deps.GetRedisPool().Cache(),
		repoMongo: repository.GetSharedRepoMongo(),
		sdk:       sdk.GetSDK(),
		// commonUC:  common.GetCommonUsecase(),
	}
}

func (uc *aclUsecaseImpl) Hello(ctx context.Context) (msg string) {
	trace := tracer.StartTrace(ctx, "AclUsecase:Hello")
	defer trace.Finish()
	ctx = trace.Context()

	return
}

func (uc *aclUsecaseImpl) SaveRole(ctx context.Context, payload domain.AddRoleRequest) (resp domain.RoleResponse, err error) {
	trace := tracer.StartTrace(ctx, "AclUsecase:SaveRole")
	defer trace.Finish()
	ctx = trace.Context()

	apps := shareddomain.Apps{Code: payload.AppsCode}
	if err := uc.repoMongo.AppsRepo.Find(ctx, &apps); err != nil {
		return resp, errors.New("Apps not found")
	}

	currentRole := shareddomain.Role{AppsID: apps.ID, Code: payload.Code}
	if err := uc.repoMongo.RoleRepo.Find(ctx, &currentRole); err != nil {
		currentRole.AppsID = apps.ID
		currentRole.Code = payload.Code
		currentRole.Name = payload.Name
	}

	permList, err := uc.repoMongo.PermissionRepo.FetchAll(ctx, appsdomain.FilterPermission{
		Filter: candishared.Filter{ShowAll: true}, AppID: apps.ID,
	})
	rootPermission := shareddomain.Permission{
		Code: apps.Code, Childs: shareddomain.MakeTreePermission(permList),
	}
	allVisitedPath := rootPermission.GetAllVisitedPath()

	currentRole.Permissions = map[string]string{}
	for _, permCode := range payload.Permissions {
		perm := shareddomain.Permission{Code: permCode}
		if err := uc.repoMongo.PermissionRepo.Find(ctx, &perm); err != nil {
			return resp, fmt.Errorf("Permission data '%s' not found", permCode)
		}
		if perm.AppsID != apps.ID {
			return resp, fmt.Errorf("Permission data '%s' invalid", permCode)
		}
		currentRole.Permissions[perm.Code] = perm.ID

		fullParentPath := allVisitedPath[perm.Code]
		for _, path := range fullParentPath {
			if path.Code == apps.Code {
				continue
			}
			currentRole.Permissions[path.Code] = path.ID
		}
	}

	err = uc.repoMongo.RoleRepo.Save(ctx, &currentRole)
	detailRole, _ := uc.GetDetailRole(ctx, currentRole.ID)

	resp.ID = currentRole.ID
	resp.Code = currentRole.Code
	resp.Name = currentRole.Name
	resp.Apps.ID = apps.ID
	resp.Apps.Code = apps.Code
	resp.Apps.Name = apps.Name
	resp.Permissions = detailRole.Permissions
	return
}

func (uc *aclUsecaseImpl) GrantUser(ctx context.Context, payload domain.GrantUserRequest) (err error) {
	trace := tracer.StartTrace(ctx, "AclUsecase:GrantUser")
	defer trace.Finish()
	ctx = trace.Context()

	_, err = uc.sdk.UserService().GetMember(ctx, payload.UserID)
	if err != nil {
		return errors.New("user not found")
	}

	roleData := shareddomain.Role{ID: payload.RoleID}
	if err := uc.repoMongo.RoleRepo.Find(ctx, &roleData); err != nil {
		return errors.New("role not found")
	}

	// handle multiple acl role for one apps
	aclList, _ := uc.repoMongo.AclRepo.FindByUserID(ctx, payload.UserID)
	var roles []string
	aclMap := make(map[string]string, len(aclList))
	for _, acl := range aclList {
		roles = append(roles, acl.RoleID)
		aclMap[acl.RoleID] = acl.ID
	}

	aclData := shareddomain.ACL{
		UserID: payload.UserID, RoleID: payload.RoleID, AppsID: roleData.AppsID,
	}
	roleGroup := uc.repoMongo.RoleRepo.GroupByID(ctx, roles...)
	for _, role := range roleGroup {
		if role.AppsID == roleData.AppsID {
			aclData.ID = aclMap[role.ID]
			break
		}
	}

	aclData.AdditionalPermissions = map[string]string{}
	if len(payload.AdditionalPermissions) > 0 {
		apps := shareddomain.Apps{ID: roleData.AppsID}
		if err := uc.repoMongo.AppsRepo.Find(ctx, &apps); err != nil {
			return errors.New("Apps not found")
		}
		permList, _ := uc.repoMongo.PermissionRepo.FetchAll(ctx, appsdomain.FilterPermission{
			Filter: candishared.Filter{ShowAll: true}, AppID: roleData.AppsID,
		})
		rootPermission := shareddomain.Permission{
			Code: apps.Code, Childs: shareddomain.MakeTreePermission(permList),
		}
		allVisitedPath := rootPermission.GetAllVisitedPath()
		for _, permCode := range payload.AdditionalPermissions {
			if _, ok := roleData.Permissions[permCode]; ok {
				return fmt.Errorf("Permission data '%s' is exist in role '%s'", permCode, roleData.Name)
			}

			perm := shareddomain.Permission{Code: permCode}
			if err := uc.repoMongo.PermissionRepo.Find(ctx, &perm); err != nil {
				return fmt.Errorf("Permission data '%s' not found", permCode)
			}
			if perm.AppsID != apps.ID {
				return fmt.Errorf("Permission data '%s' invalid", permCode)
			}
			aclData.AdditionalPermissions[perm.Code] = perm.ID

			fullParentPath := allVisitedPath[perm.Code]
			for _, path := range fullParentPath {
				if path.Code == apps.Code {
					continue
				}
				aclData.AdditionalPermissions[path.Code] = path.ID
			}
		}
	}

	return uc.repoMongo.AclRepo.Save(ctx, &aclData)
}

func (uc *aclUsecaseImpl) GetPermission(ctx context.Context, userID, appsID string) (data domain.CheckPermissionResponse, err error) {
	trace := tracer.StartTrace(ctx, "AclUsecase:GetPermission")
	defer trace.Finish()
	ctx = trace.Context()

	return
}

func (uc *aclUsecaseImpl) CheckPermission(ctx context.Context, userID string, permissionCode string) (roleID string, err error) {
	trace := tracer.StartTrace(ctx, "AclUsecase:CheckPermission")
	defer trace.Finish()
	ctx = trace.Context()
	trace.SetTag("userId", userID)

	acl, err := uc.repoMongo.AclRepo.FindByUserID(ctx, userID)
	if err != nil || len(acl) == 0 {
		return roleID, errors.New("ACL not found for this user")
	}

	var roles []string
	for _, a := range acl {
		if _, ok := a.AdditionalPermissions[permissionCode]; ok {
			return a.RoleID, nil
		}
		roles = append(roles, a.RoleID)
	}

	roleGroup := uc.repoMongo.RoleRepo.GroupByID(ctx, roles...)
	for _, role := range roleGroup {
		if _, ok := role.Permissions[permissionCode]; ok {
			return role.ID, nil
		}
	}
	return roleID, errors.New("Access not found")
}

func (uc *aclUsecaseImpl) GetAllRole(ctx context.Context, filter domain.RoleListFilter) (data []domain.RoleResponse, meta candishared.Meta, err error) {
	trace := tracer.StartTrace(ctx, "AclUsecase:GetAllRole")
	defer trace.Finish()
	ctx = trace.Context()

	filter.CalculateOffset()
	count := uc.repoMongo.RoleRepo.Count(ctx, filter)

	roles, err := uc.repoMongo.RoleRepo.FetchAll(ctx, filter)
	if err != nil {
		return data, meta, err
	}

	for _, role := range roles {
		apps := shareddomain.Apps{ID: role.AppsID}
		uc.repoMongo.AppsRepo.Find(ctx, &apps)
		roleDetail := domain.RoleResponse{
			ID:   role.ID,
			Code: role.Code,
			Name: role.Name,
		}
		roleDetail.Apps.ID = apps.ID
		roleDetail.Apps.Code = apps.Code
		roleDetail.Apps.Name = apps.Name
		data = append(data, roleDetail)
	}

	meta = candishared.NewMeta(filter.Page, filter.Limit, int(count))
	return
}

func (uc *aclUsecaseImpl) GetDetailRole(ctx context.Context, roleID string) (data domain.RoleResponse, err error) {
	trace := tracer.StartTrace(ctx, "AclUsecase:GetDetailRole")
	defer trace.Finish()
	ctx = trace.Context()

	role := shareddomain.Role{ID: roleID}
	if err = uc.repoMongo.RoleRepo.Find(ctx, &role); err != nil {
		return data, err
	}

	permFilter := appsdomain.FilterPermission{
		Filter: candishared.Filter{ShowAll: true},
	}
	for _, perm := range role.Permissions {
		permFilter.PermissionIDs = append(permFilter.PermissionIDs, perm)
	}

	var permissions []shareddomain.Permission
	if len(permFilter.PermissionIDs) > 0 {
		permissions, err = uc.repoMongo.PermissionRepo.FetchAll(ctx, permFilter)
		if err != nil || len(permissions) == 0 {
			return data, errors.New("Data not found")
		}
	}

	apps := shareddomain.Apps{ID: role.AppsID}
	uc.repoMongo.AppsRepo.Find(ctx, &apps)

	data.ID = role.ID
	data.Code = role.Code
	data.Name = role.Name
	data.Apps.ID = apps.ID
	data.Apps.Code = apps.Code
	data.Apps.Name = apps.Name
	data.Permissions = shareddomain.MakeTreePermission(permissions)
	return
}

func (uc *aclUsecaseImpl) RevokeUserRole(ctx context.Context, userID, roleID string) (err error) {
	trace := tracer.StartTrace(ctx, "AclUsecase:RevokeUserRole")
	defer trace.Finish()
	ctx = trace.Context()

	acl := shareddomain.ACL{UserID: userID, RoleID: roleID}
	if err := uc.repoMongo.AclRepo.Find(ctx, &acl); err != nil {
		return err
	}

	return uc.repoMongo.AclRepo.Delete(ctx, &acl)
}
