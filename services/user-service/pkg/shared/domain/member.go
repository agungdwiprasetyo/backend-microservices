package domain

import "time"

// Member model
type Member struct {
	ID         string    `json:"id" bson:"_id"`
	Username   string    `json:"username" bson:"username"`
	Password   string    `json:"password" bson:"password"`
	Fullname   string    `json:"fullname" bson:"fullname"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt" bson:"modifiedAt"`
}
