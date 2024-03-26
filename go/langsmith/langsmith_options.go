package langsmith

import (
	"context"
	"net/http"
)

// WithAPIKey is a ClientOption that sets the Authorization header to the given API key.
func WithAPIKey(apiKey string) ClientOption {
	return WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("x-api-key", apiKey)
		return nil
	})
}
