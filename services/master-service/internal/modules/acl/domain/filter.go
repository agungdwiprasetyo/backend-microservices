package domain

import "pkg.agungdp.dev/candi/candishared"

// RoleListFilter filter
type RoleListFilter struct {
	candishared.Filter
	AppsID  string   `json:"apps_id"`
	RoleIDs []string `json:"role_ids"`
}

// ACLFilter filter
type ACLFilter struct {
	candishared.Filter
	AppsID string
	UserID string
}
