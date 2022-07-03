package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"simplems/handlers"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func loggingMiddleware(l *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Debug("Request", zap.String("url", r.RequestURI), zap.String("method", r.Method))
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
}

func decodeProductMiddleware(l *zap.Logger) func(http.Handler) http.Handler {
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

				l.Debug("decoded product info", zap.String("name", prod.Name))
				// add the product to the context
				ctx := context.WithValue(r.Context(), handlers.ProductKey{}, prod)
				r = r.WithContext(ctx)
			}
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	logger, _ := zap.NewDevelopment()

	defer logger.Sync()
	healthCheckHandler := handlers.NewHealthCheck(logger)
	productsHandler := handlers.NewProducts(logger)
	router := mux.NewRouter()
	psr := router.PathPrefix("/products").Subrouter()
	psr.Use(loggingMiddleware(logger), decodeProductMiddleware(logger))
	psr.HandleFunc("/", productsHandler.ListProducts).Methods(http.MethodGet)
	psr.HandleFunc("/", productsHandler.CreateProduct).Methods(http.MethodPost)
	psr.HandleFunc("/{id:[0-9]+}/", productsHandler.UpdateProduct).Methods(http.MethodPut)
	router.Handle("/ping", healthCheckHandler)

	server := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// anonymouse go routine which runs concurrently in the backgroun!
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal("Fatal error occured", zap.Error(err))
		}
	}()

	notifyChannel := make(chan os.Signal, 1)
	signal.Notify(notifyChannel, os.Interrupt, syscall.SIGTERM)

	// block untill signal is recieved
	signal := <-notifyChannel
	logger.Info("Gracefully Shutting down...", zap.String("reason", signal.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}
