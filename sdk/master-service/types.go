package masterservice

// PayloadCheckPermission payload
type PayloadCheckPermission struct {
	UserID         string
	PermissionCode string
}

// UserApps response data
type UserApps struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	FrontendURL string `json:"frontendUrl"`
	BackendURL  string `json:"backendUrl"`
	Role        struct {
		ID   string `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"role"`
}
