package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"simplems/data"
	"simplems/handlers"

	"go.uber.org/zap"
)

func ValidateProductIdMiddleware(l *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//Skip midddleware for list products Get requests
			re := regexp.MustCompile("[0-9]+")
			hasIdInPath := len(re.FindAllString(r.URL.Path, -1)) == 1
			if !hasIdInPath && r.Method == http.MethodGet {
				l.Debug("List Products Request", zap.String("url", r.RequestURI), zap.String("method", r.Method))
			} else if r.Method == http.MethodGet || r.Method == http.MethodPut || r.Method == http.MethodDelete {
				l.Debug("Request", zap.String("url", r.RequestURI), zap.String("method", r.Method), zap.Bool("decoding", true))

				pid, err := handlers.GetProductId(r, l)
				if err != nil {
					http.Error(w, "Invalid product id.", http.StatusBadRequest)
					return
				}

				l.Info("Parsed product id", zap.Int("id", pid))
				pIndex := data.GetProductIndexById(pid)
				if pIndex == -1 {
					l.Error(fmt.Sprintf("Failed to `%s` product, invalid product id.", r.Method), zap.Int("id", pid))
					http.Error(w, fmt.Sprintf("Failed to `%s` product, invalid product id.", r.Method), http.StatusBadRequest)
					return
				}
				l.Info("Got product index from request:", zap.Int("index", pIndex))
				// add the product to the context
				ctx := context.WithValue(r.Context(), handlers.ValidatedProductIndexKey{}, pIndex)
				r = r.WithContext(ctx)
			}
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
}
