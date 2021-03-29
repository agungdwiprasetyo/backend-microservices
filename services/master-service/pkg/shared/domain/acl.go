package domain

import "time"

// ACL model
type ACL struct {
	ID                    string            `json:"id" bson:"_id"`
	UserID                string            `json:"userId" bson:"userId"`
	AppsID                string            `json:"appsId" bson:"appsId"`
	RoleID                string            `json:"roleId" bson:"roleId"`
	AdditionalPermissions map[string]string `json:"additionalPermissions" bson:"additionalPermissions"`
	CreatedAt             time.Time         `json:"createdAt" bson:"createdAt"`
	ModifiedAt            time.Time         `json:"modifiedAt" bson:"modifiedAt"`
}

// Role model
type Role struct {
	ID          string            `json:"id" bson:"_id"`
	AppsID      string            `json:"appsId" bson:"appsId"`
	Code        string            `json:"code" bson:"code"`
	Name        string            `json:"name" bson:"name"`
	Permissions map[string]string `json:"permissions" bson:"permissions"`
	CreatedAt   time.Time         `json:"createdAt" bson:"createdAt"`
	ModifiedAt  time.Time         `json:"modifiedAt" bson:"modifiedAt"`
}
