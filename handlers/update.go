package handlers

import (
	"net/http"
	"simplems/data"

	"go.uber.org/zap"
)

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      int                  true  "Product ID"
// @Param        product  body      data.Product true  "Update product"
// @Success      200      {object}  data.Product
// @Failure      400      {object}  HTTPError
// @Failure      404      {object}  HTTPError
// @Failure      500      {object}  HTTPError
// @Router       /products/{id} [patch]
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	incommingProd := r.Context().Value(ProductKey{}).(*data.Product)
	prodIndex := r.Context().Value(ProductIndexKey{}).(int)
	updatedProd, _ := data.UpdateProduct(incommingProd, prodIndex)

	rw.Header().Add("Content-Type", "application/json")
	data.ToJson(updatedProd, rw)
	p.l.Info("Update Products Response: ", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}
