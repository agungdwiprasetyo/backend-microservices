package domain

import "pkg.agungdp.dev/candi/candishared"

type FilterModule struct {
	candishared.Filter
	AppsID string
}

type FilterPermission struct {
	candishared.Filter
	AppID         string `json:"app_id"`
	PermissionIDs []string
	Codes         []string
}
