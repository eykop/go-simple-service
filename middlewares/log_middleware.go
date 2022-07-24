package middlewares

import (
	"net/http"

	"go.uber.org/zap"
)

func LoggingMiddleware(l *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Debug("Request", zap.String("url", r.RequestURI), zap.String("method", r.Method))
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
}
