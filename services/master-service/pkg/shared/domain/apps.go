package domain

import (
	"container/list"
	"time"
)

// Apps model
type Apps struct {
	ID           string       `json:"id" bson:"_id"`
	Code         string       `json:"code" bson:"code"`
	Name         string       `json:"name" bson:"name"`
	Icon         string       `json:"icon" bson:"icon"`
	FrontendURL  string       `json:"frontendUrl" bson:"frontendUrl"`
	BackendURL   string       `json:"backendUrl" bson:"backendUrl"`
	CreatedAt    string       `json:"-" bson:"-"`
	ModifiedAt   string       `json:"-" bson:"-"`
	CreatedAtDB  time.Time    `json:"-" bson:"createdAt"`
	ModifiedAtDB time.Time    `json:"-" bson:"modifiedAt"`
	Permissions  []Permission `json:"permission,omitempty" bson:"-"`
}

// CollectionName for model
func (Apps) CollectionName() string {
	return "apps"
}

// TableName for model
func (Apps) TableName() string {
	return "apps"
}

// Permission model
type Permission struct {
	ID           string       `json:"id" bson:"_id"`
	AppsID       string       `json:"apps_id,omitempty" bson:"appsId"`
	Code         string       `json:"code" bson:"code"`
	Name         string       `json:"name" bson:"name"`
	Icon         string       `json:"icon" bson:"icon"`
	URL          string       `json:"url" bson:"url"`
	NewPage      bool         `json:"new_page" bson:"new_page"`
	ParentID     string       `json:"parent_id,omitempty" bson:"parentId"`
	Childs       []Permission `json:"childs" bson:"-"`
	CreatedAt    string       `json:"-" bson:"-"`
	ModifiedAt   string       `json:"-" bson:"-"`
	CreatedAtDB  time.Time    `json:"-" bson:"createdAt"`
	ModifiedAtDB time.Time    `json:"-" bson:"modifiedAt"`
}

// CollectionName for model
func (Permission) CollectionName() string {
	return "apps_permissions"
}

// TableName for model
func (Permission) TableName() string {
	return "apps_permissions"
}

// GetAllVisitedPath func using BFS
func (p Permission) GetAllVisitedPath() (nodeVisitedPaths map[string][]Permission) {
	visited := make(map[string]struct{})
	nodeVisitedPaths = make(map[string][]Permission)
	queue := list.New()
	queue.PushBack(p)
	visited[p.Code] = struct{}{}

	for queue.Len() > 0 {
		qNode := queue.Front()
		valNode := qNode.Value.(Permission)
		for _, node := range valNode.Childs {
			if _, ok := visited[node.Code]; !ok {
				visited[node.Code] = struct{}{}
				queue.PushBack(node)
				nodeVisitedPaths[node.Code] = nodeVisitedPaths[valNode.Code]
				nodeVisitedPaths[node.Code] = append(nodeVisitedPaths[node.Code], valNode)
			}
		}
		queue.Remove(qNode)
	}

	return nodeVisitedPaths
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
