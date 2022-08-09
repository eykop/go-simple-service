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
func GetProductId(r *http.Request, l *zap.Logger) (int, error) {
	vars := mux.Vars(r)
	id := vars["id"]

	// we now have validator in the path param of a regexp of [0-9]+, so wif a product id of not number is sent
	// we will get here  when the user send a very large number (upper limit of int)
	productId, err := strconv.Atoi(id)
	if err != nil {
		l.Error("Failed to parse product id", zap.String("string", id), zap.Error(err))
	}
	return productId, err
}
