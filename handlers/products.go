package handlers

import (
	"go.uber.org/zap"
)

func NewProducts(l *zap.Logger) *Products {
	return &Products{l}
}
