package domain

import "time"

// Token model
type Token struct {
	ID           string                 `bson:"_id" json:"id"`
	UserID       string                 `bson:"user_id" json:"user_id"`
	DeviceID     string                 `bson:"device_id" json:"device_id"`
	Token        string                 `bson:"token" json:"token"`
	RefreshToken string                 `bson:"refresh_token" json:"refresh_token"`
	IsActive     *bool                  `bson:"is_active" json:"is_active"`
	Claims       map[string]interface{} `bson:"claims" json:"claims"`
	ExpiredAt    time.Time              `bson:"expired_at" json:"expired_at"`
	CreatedAt    time.Time              `bson:"created_at" json:"created_at"`
	ModifiedAt   time.Time              `bson:"modified_at" json:"modified_at"`
}

// CollectionName for token model
func (Token) CollectionName() string {
	return "tokens"
}
