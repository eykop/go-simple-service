package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"simplems/handlers"

	"go.uber.org/zap"
)

func DecodeProductMiddleware(l *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost || r.Method == http.MethodPut {
				l.Debug("Request", zap.String("url", r.RequestURI), zap.String("method", r.Method), zap.Bool("decoding", true))
				prod := handlers.JsonToProduct(r, l)
				if prod == nil {
					l.Error("Failed to read produc, could not decode product from json.")
					http.Error(w, "Failed to deserialize product from json.", http.StatusBadRequest)
					return
				}

				validationErr := prod.Validate()
				if validationErr != nil {
					l.Error("Failed to validate produc.", zap.Error(validationErr))
					http.Error(w, fmt.Sprintf("Failed to validate product %s.", validationErr), http.StatusBadRequest)
					return
				}

				l.Debug("decoded product info", zap.Int("name", prod.GetID()))
				// add the product to the context
				ctx := context.WithValue(r.Context(), handlers.ProductKey{}, prod)
				r = r.WithContext(ctx)
			}
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
}
