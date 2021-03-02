package domain

import (
	shareddomain "monorepo/services/master-service/pkg/shared/domain"
)

// GrantUserRequest payload
type GrantUserRequest struct {
	UserID string `json:"userId"`
	RoleID string `json:"roleId"`
}

// AddRoleRequest model
type AddRoleRequest struct {
	AppsCode    string   `json:"appsCode"`
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// AddRoleResponse model
type AddRoleResponse struct {
	ID          string                    `json:"id"`
	AppsID      string                    `json:"appsId"`
	Code        string                    `json:"code"`
	Name        string                    `json:"name"`
	Permissions []shareddomain.Permission `json:"permissions"`
}

// Permissions model
type Permissions struct {
	ID string `json:"id"`
}

// CheckPermissionRequest payload
type CheckPermissionRequest struct {
	UserID         string `json:"userId"`
	PermissionCode string `json:"permissionCode"`
}

// CheckPermissionResponse payload
type CheckPermissionResponse struct {
	RoleID string                    `json:"userId"`
	Access []shareddomain.Permission `json:"access"`
}

// RoleResponse response payload
type RoleResponse struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Apps struct {
		ID   string `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"apps"`
	Permissions []shareddomain.Permission `json:"permissions,omitempty"`
}
