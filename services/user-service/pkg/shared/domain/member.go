package domain

import "time"

// Member model
type Member struct {
	ID           string    `json:"id" bson:"_id"`
	Username     string    `json:"username" bson:"username"`
	Password     string    `json:"-" bson:"password"`
	PasswordSalt string    `json:"-" bson:"password_salt"`
	Fullname     string    `json:"fullname" bson:"fullname"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	ModifiedAt   time.Time `json:"modifiedAt" bson:"modifiedAt"`
}

// CollectionName for member model
func (Member) CollectionName() string {
	return "members"
}
