package handlers

import "go.uber.org/zap"

type ProductKey struct{}

type ValidatedProductIndexKey struct{}

type Products struct {
	l *zap.Logger
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
