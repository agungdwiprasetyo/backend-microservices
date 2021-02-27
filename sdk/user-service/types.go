package userservice

// Member model
type Member struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Fullname   string `json:"fullname"`
	CreatedAt  string `json:"createdAt"`
	ModifiedAt string `json:"modifiedAt"`
}
