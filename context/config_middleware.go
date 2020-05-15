package context

import (
	"context"
	"errors"
	"net/http"
)

// ConfigMiddleware is a middleware to to inject the app config object
func ConfigMiddleware(cfg *Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				ctx := context.WithValue(r.Context(), KeyAppConfig, cfg)
				next.ServeHTTP(w, r.WithContext(ctx))
			})
	}
}

// GetConfigFromContext returns the config from the current context
func GetConfigFromContext(ctx context.Context) (*Config, error) {
	errMsg := errors.New("No config in context")
	if ctx.Value(KeyAppConfig) == nil {
		return nil, errMsg
	}

	cfg, ok := ctx.Value(KeyAppConfig).(*Config)

	if !ok {
		return nil, errMsg
	}

	return cfg, nil
}
