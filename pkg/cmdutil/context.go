package cmdutil

import (
	"context"
	"fmt"

	"github.com/bnassif/anetgo/pkg/api"
)

type contextKey string

const clientKey = contextKey("client")
const rawResponseKey = contextKey("raw_response")

// WithClient returns a new context containing the API client
func WithClient(ctx context.Context, client *api.Client) context.Context {
	return context.WithValue(ctx, clientKey, client)
}

// GetClient extracts the API client from context
func GetClient(ctx context.Context) *api.Client {
	if v := ctx.Value(clientKey); v != nil {
		if c, ok := v.(*api.Client); ok {
			return c
		}
	}
	return nil
}

// WithBool returns a new context containing a Bool value
func WithBool(ctx context.Context, value *bool) context.Context {
	return context.WithValue(ctx, rawResponseKey, value)
}

func GetRawFlagValue(ctx context.Context) *bool {
	defaultValue := bool(false)
	if v := ctx.Value(rawResponseKey); v != nil {
		if c, ok := v.(*bool); ok {
			return c
		}
	}
	fmt.Print("Return default false\n")
	return &defaultValue
}
