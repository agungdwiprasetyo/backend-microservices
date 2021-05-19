package domain

import (
	masterservice "monorepo/sdk/master-service"
	shareddomain "monorepo/services/user-service/pkg/shared/domain"
)

// LoginRequest request payload
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse request payload
type LoginResponse struct {
	Token        string                   `json:"token"`
	RefreshToken string                   `json:"refresh_token"`
	Profile      shareddomain.Member      `json:"profile"`
	UserApps     []masterservice.UserApps `json:"userApps"`
}
