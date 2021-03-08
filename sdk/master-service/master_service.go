package masterservice

import (
	"context"
)

// MasterService abstract interface
type MasterService interface {
	CheckPermission(ctx context.Context, userID string, permissionCode string) (role string, err error)
}
