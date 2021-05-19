package domain

// RegisterPayload payload
type RegisterPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}
