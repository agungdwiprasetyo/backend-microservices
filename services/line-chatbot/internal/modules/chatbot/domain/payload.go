package domain

// PushMessagePayload model
type PushMessagePayload struct {
	To      string `json:"to"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
