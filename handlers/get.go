package handlers

import (
	"net/http"
	"simplems/data"

	"go.uber.org/zap"
)

// ListProducts godoc
// @Summary     list products
// @Description get all producs currently available
// @Tags        products
// @Accept       json
// @Produce     json
// @Success     200 {array} data.Product "ok"
// @Failure      500  {object}  HTTPError
// @Router      /products/ [get]
func (p *Products) ListProducts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	err := data.GetProductsList().ToJson(rw)
	if err != nil {
		p.l.Error("Failed to List products", zap.Error(err))
		http.Error(rw, "Failed to list products", http.StatusInternalServerError)
	}
	p.l.Info("List Products Response", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}

// GetProduct godoc
// @Summary      Gets a product
// @Description  get product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  data.Product{name=string}
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /products/{id} [get]
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {

	prodIndex := r.Context().Value(ValidatedProductIndexKey{}).(int)
	rw.Header().Add("Content-Type", "application/json")
	prodErr := data.GetProductByIndex(prodIndex).ToJson(rw)
	if prodErr != nil {
		p.l.Error("Failed to encode product.", zap.Error(prodErr))
		http.Error(rw, "Failed to encode product.", http.StatusInternalServerError)
	}
	p.l.Info("Get Product Response", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}
