package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simplems/handlers"
	"syscall"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "healtheCheck ", log.LstdFlags)
	healthCheckHandler := handlers.NewHealthCheck(logger)
	sm := http.NewServeMux()
	sm.Handle("/ping", healthCheckHandler)

	server := &http.Server{
		Addr:         ":3000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// anonymouse go routine which runs concurrently in the backgroun!
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	notifyChannel := make(chan os.Signal, 1)
	signal.Notify(notifyChannel, os.Interrupt, syscall.SIGTERM)

	// block untill signal is recived
	signal := <-notifyChannel
	logger.Println("Gracefully Shutting down...", signal)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}
