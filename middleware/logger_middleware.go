package middleware

import (
	"context"
	"net/http"

	gcontext "github.com/fusion44/couchless-backend/context"
	"github.com/op/go-logging"
	"github.com/pkg/errors"
)

// LoggerMiddleware is a middleware to inject the logger to a context
func LoggerMiddleware(l *logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				ctx := context.WithValue(r.Context(), gcontext.KeyLogger, l)
				next.ServeHTTP(w, r.WithContext(ctx))
			})
	}
}

// GetLoggerFromContext returns the logger from the current context
func GetLoggerFromContext(ctx context.Context) (*logging.Logger, error) {
	errMsg := errors.New("No logger in context")
	if ctx.Value(gcontext.KeyLogger) == nil {
		return nil, errMsg
	}

	l, ok := ctx.Value(gcontext.KeyLogger).(*logging.Logger)

	if !ok {
		return nil, errMsg
	}

	return l, nil
}
