package domain

import "pkg.agungdp.dev/candi/candishared"

// RoleListFilter filter
type RoleListFilter struct {
	candishared.Filter
	AppsID  string
	RoleIDs []string
}

// ACLFilter filter
type ACLFilter struct {
	candishared.Filter
	AppsID string
	UserID string
}
