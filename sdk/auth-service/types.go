package authservice

import "github.com/dgrijalva/jwt-go"

// PayloadGenerateToken payload
type PayloadGenerateToken struct {
	UserID   string
	Username string
}

// ResponseClaim for token claim data
type ResponseClaim struct {
	jwt.StandardClaims
	DeviceID string `json:"did"`
	User     struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
	Alg string `json:"-"`
}

// ResponseGenerateToken model
type ResponseGenerateToken struct {
	Token        string        `json:"token"`
	RefreshToken string        `json:"refresh_token"`
	Claim        ResponseClaim `json:"claim"`
}
