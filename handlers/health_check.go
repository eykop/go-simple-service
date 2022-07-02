package handlers

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type HealthCheck struct {
	l *zap.Logger
}

func NewHealthCheck(l *zap.Logger) *HealthCheck {
	return &HealthCheck{l}
}

func (h *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Health Check: ", zap.String("status", "healthy"))
	// h.l.Println("User Agent:", r.UserAgent())
	// h.l.Printf("%s %s %s %d\n", r.RemoteAddr, r.Method, r.URL, http.StatusOK)
	h.l.Info("Response: ", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
	//write to response ...
	fmt.Fprintf(w, "pong")
}
