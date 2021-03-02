package domain

import shareddomain "monorepo/services/user-service/pkg/shared/domain"

// LoginRequest request payload
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse request payload
type LoginResponse struct {
	Detail shareddomain.Member `json:"detail"`
	Token  string              `json:"token"`
}
