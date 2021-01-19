package translator

import "context"

// Translator abstract interface
type Translator interface {
	Translate(ctx context.Context, from, to, text string) (result string)
}
