package domain

import (
	"time"
)

// Apps model
type Apps struct {
	ID           string       `json:"id" bson:"_id"`
	Code         string       `json:"code" bson:"code"`
	Name         string       `json:"name" bson:"name"`
	Icon         string       `json:"icon" bson:"icon"`
	URL          string       `json:"url" bson:"url"`
	CreatedAt    string       `json:"-" bson:"-"`
	ModifiedAt   string       `json:"-" bson:"-"`
	CreatedAtDB  time.Time    `json:"-" bson:"createdAt"`
	ModifiedAtDB time.Time    `json:"-" bson:"modifiedAt"`
	Permissions  []Permission `json:"permission" bson:"-"`
}

// Module model
type Module struct {
	ID           string       `json:"id" bson:"_id"`
	AppsID       string       `json:"appsId" bson:"appsId"`
	Code         string       `json:"code" bson:"code"`
	Name         string       `json:"name" bson:"name"`
	Icon         string       `json:"icon" bson:"icon"`
	CreatedAt    string       `json:"createdAt" bson:"-"`
	ModifiedAt   string       `json:"modifiedAt" bson:"-"`
	CreatedAtDB  time.Time    `json:"-" bson:"createdAt"`
	ModifiedAtDB time.Time    `json:"-" bson:"modifiedAt"`
	Permissions  []Permission `json:"permission" bson:"-"`
}

// Permission model
type Permission struct {
	ID           string       `json:"id" bson:"_id"`
	AppsID       string       `json:"appsId,omitempty" bson:"appsId"`
	Code         string       `json:"code" bson:"code"`
	Name         string       `json:"name" bson:"name"`
	Icon         string       `json:"icon" bson:"icon"`
	ParentID     string       `json:"parentId,omitempty" bson:"parentId"`
	Childs       []Permission `json:"childs" bson:"-"`
	CreatedAt    string       `json:"-" bson:"-"`
	ModifiedAt   string       `json:"-" bson:"-"`
	CreatedAtDB  time.Time    `json:"-" bson:"createdAt"`
	ModifiedAtDB time.Time    `json:"-" bson:"modifiedAt"`
}
