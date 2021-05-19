package domain

const (
	// RedisTokenExpiredKeyConst const
	RedisTokenExpiredKeyConst = "expiredtoken"
)

// RedisTokenExpiredKey for redis pubsub model
type RedisTokenExpiredKey struct {
	DeviceID string `json:"deviceid"`
	UserID   string `json:"userid"`
}

// ResponseGenerateToken model
type ResponseGenerateToken struct {
	Token        string                 `json:"token"`
	RefreshToken string                 `json:"refresh_token"`
	Claim        map[string]interface{} `json:"claim"`
}
