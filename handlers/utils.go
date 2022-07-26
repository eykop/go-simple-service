package handlers

import (
	"net/http"
	"simplems/data"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Encodes a json data into a Prodct object
func JsonToProduct(r *http.Request, l *zap.Logger) *data.Product {
	p := &data.Product{}
	errorToJson := data.FromJson(p, r.Body)
	if errorToJson != nil {
		l.Error("Failed to encode product", zap.Error(errorToJson))
		return nil
	}
	return p
}

// Parses product ID from request path parameter.
// returns -1 if failed to parse a valid product id.
func getProductId(r *http.Request, l *zap.Logger) int {
	vars := mux.Vars(r)
	id := vars["id"]

	productId, err := strconv.Atoi(id)
	if err != nil {
		l.Error("Failed to parse product id", zap.Error(err))
		return -1
	}
	return productId
}
