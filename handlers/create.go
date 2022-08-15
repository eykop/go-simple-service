package handlers

import (
	"net/http"
	"simplems/data"

	"go.uber.org/zap"
)

// AddProduct godoc
// @Summary      Add a product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      data.Product  true  "Add product"
// @Success      200      {object}  data.Product
// @Failure      400      {object}  HTTPError
// @Failure      404      {object}  HTTPError
// @Failure      500      {object}  HTTPError
// @Router       /products [post]
func (p *Products) CreateProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Info("Will create a new product")
	prod := r.Context().Value(ProductKey{}).(data.ProductInterface)
	data.AddPorduct(prod)
	rw.Header().Add("Content-Type", "application/json")
	prod.ToJson(rw)
	p.l.Info("Create Product Response: ", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}
