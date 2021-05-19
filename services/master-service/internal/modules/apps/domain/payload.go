package domain

// AppsRequest model
type AppsRequest struct {
	Code        string              `json:"code"`
	Name        string              `json:"name"`
	Permissions []PermissionRequest `json:"permissions"`
}

// PermissionRequest model
type PermissionRequest struct {
	ID       string              `json:"-"`
	ParentID string              `json:"-"`
	Code     string              `json:"code"`
	Name     string              `json:"name"`
	Icon     string              `json:"icon"`
	URL      string              `json:"url"`
	NewPage  bool                `json:"new_page"`
	Childs   []PermissionRequest `json:"childs"`
}
