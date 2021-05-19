package domain

import "time"

// ACL model
type ACL struct {
	ID         string    `json:"id" bson:"_id"`
	UserID     string    `json:"userId" bson:"userId"`
	AppsID     string    `json:"appsId" bson:"appsId"`
	RoleID     string    `json:"roleId" bson:"roleId"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt" bson:"modifiedAt"`
}

// CollectionName for model
func (ACL) CollectionName() string {
	return "access_control_list"
}

// TableName for model
func (ACL) TableName() string {
	return "access_control_list"
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

// CollectionName for model
func (Role) CollectionName() string {
	return "roles"
}

// TableName for model
func (Role) TableName() string {
	return "roles"
}
