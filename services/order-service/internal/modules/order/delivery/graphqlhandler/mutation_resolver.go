// Code generated by candi v1.3.3.

package graphqlhandler

import "context"

type mutationResolver struct {
	root *GraphQLHandler
}

// Hello resolver
func (m *mutationResolver) Hello(ctx context.Context) (string, error) {
	return "Hello", nil
}	
