package domain

// PushRequest model
type PushRequest struct {
	To           string                 `json:"to"`
	Notification *Notification          `json:"notification"`
	Data         map[string]interface{} `json:"data"`
}

// Notification model
type Notification struct {
	Title          string `json:"title"`
	Body           string `json:"body"`
	Image          string `json:"image"`
	Sound          string `json:"sound"`
	MutableContent bool   `json:"mutable-content"`
	ResourceID     string `json:"resourceId"`
	ResourceName   string `json:"resoureceName"`
}

// PushResponse response data
type PushResponse struct {
	Error bool        `json:"error"`
	Body  interface{} `json:"body"`
}

// Event domain
type Event struct {
	ID        string `json:"id"`
	Topic     string `json:"topic"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}
