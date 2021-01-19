package domain

// LineMessage domain model
type LineMessage struct {
	To       string               `json:"to"`
	Messages []LineContentMessage `json:"messages"`
}

// LineContentMessage model
type LineContentMessage struct {
	Type     string            `json:"type"`
	AltText  string            `json:"altText"`
	Contents LineContentFormat `json:"contents"`
}

// LineContentFormat model
type LineContentFormat struct {
	Type string          `json:"type"`
	Body LineContentBody `json:"body"`
}

// LineContentBody model
type LineContentBody struct {
	Type     string        `json:"type"`
	Layout   string        `json:"layout"`
	Contents []LineContent `json:"contents"`
}

// LineContent model
type LineContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
