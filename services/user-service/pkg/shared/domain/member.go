package domain

import "time"

// Member model
type Member struct {
	ID           string    `json:"id" gorm:"column:id" bson:"_id"`
	Username     string    `json:"username" gorm:"column:username" bson:"username"`
	Password     string    `json:"-" gorm:"column:password" bson:"password"`
	PasswordSalt string    `json:"-" gorm:"column:password_salt" bson:"password_salt"`
	Fullname     string    `json:"fullname" gorm:"column:fullname" bson:"fullname"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at" bson:"createdAt"`
	ModifiedAt   time.Time `json:"modifiedAt" gorm:"column:modified_at" bson:"modifiedAt"`
}

// CollectionName for member model
func (Member) CollectionName() string {
	return "members"
}
