package domain

import (
	shareddomain "monorepo/services/master-service/pkg/shared/domain"
)

// AppDetail response data
type AppDetail struct {
	ID          string                    `json:"id"`
	Code        string                    `json:"code"`
	Name        string                    `json:"name"`
	Icon        string                    `json:"icon"`
	URL         string                    `json:"url"`
	Permissions []shareddomain.Permission `json:"permissions"`
}

// Permission response data
type Permission struct {
	FullCode string `json:"fullCode"`
}
