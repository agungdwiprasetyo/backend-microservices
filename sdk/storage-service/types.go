package storageservice

// Header model
type Header struct {
	ContentType string
	Folder      string
	Filename    string
	Size        int64
}

// Response model
type Response struct {
	Location string `json:"file"`
	Size     int64  `json:"size"`
}
