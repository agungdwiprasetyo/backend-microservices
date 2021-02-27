package userservice

import (
	"context"
)

// UserService abstract interface
type UserService interface {
	GetMember(ctx context.Context, id string) (Member, error)
}
