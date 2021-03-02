package acl

import (
	"context"
	"monorepo/services/master-service/internal/modules/acl/domain"
)

// ACLUsecase abstraction
type ACLUsecase interface {
	// add method
	Hello(ctx context.Context) string
	SaveRole(ctx context.Context, payload domain.AddRoleRequest) (err error)
	GrantUser(ctx context.Context, payload domain.GrantUserRequest) (err error)
	CheckPermission(ctx context.Context, payload domain.CheckPermissionRequest) (err error)
}
