package storageservice

import (
	"context"
	"io"
)

// StorageService abstraction
type StorageService interface {
	Upload(ctx context.Context, file io.Reader, metadata Header) (Response, error)
}
