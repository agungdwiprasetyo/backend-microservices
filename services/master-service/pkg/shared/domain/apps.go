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

// MakeTreePermission construct tree permission data  (parent-child)
func MakeTreePermission(permissions []Permission) (results []Permission) {
	permParentGroups := make(map[string][]Permission)
	for _, perm := range permissions {
		permParentGroups[perm.ParentID] = append(permParentGroups[perm.ParentID], perm)
		if perm.ParentID == "" {
			results = append(results, perm)
		}
	}

	var findAllChild func(parentID string) []Permission
	findAllChild = func(parentID string) (childs []Permission) {
		childs = permParentGroups[parentID]
		for i, c := range childs {
			c.AppsID = ""
			c.Childs = append(c.Childs, findAllChild(c.ID)...)
			childs[i] = c
		}
		return childs
	}

	for i, perm := range results {
		perm.AppsID = ""
		perm.Childs = append(perm.Childs, findAllChild(perm.ID)...)
		results[i] = perm
	}

	return
}
