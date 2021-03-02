package masterservice

import (
	"context"
)

// MasterService abstract interface
type MasterService interface {
	CheckPermission(ctx context.Context, req PayloadCheckPermission) (isAllowed bool, err error)
}
