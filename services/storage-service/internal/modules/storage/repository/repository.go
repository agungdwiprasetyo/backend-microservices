// Code generated by candi v1.3.1.

package repository

import (
	"context"
)

// StorageRepository abstract interface
type StorageRepository interface {
	// add method
	FindHello(ctx context.Context) (string, error)
}