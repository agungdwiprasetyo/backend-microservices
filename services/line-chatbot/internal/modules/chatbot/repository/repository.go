// Code generated by candi v1.3.1.

package repository

import (
	"context"
)

// ChatbotRepository abstract interface
type ChatbotRepository interface {
	// add method
	FindHello(ctx context.Context) (string, error)
}
