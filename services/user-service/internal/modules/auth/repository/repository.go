// Code generated by candi v1.3.1.

package repository

import (
	"context"
)

// AuthRepository abstract interface
type AuthRepository interface {
	// add method
	FindHello(ctx context.Context) (string, error)
}