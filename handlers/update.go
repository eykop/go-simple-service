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

	productId := getProductId(r, p.l)
	if productId == -1 {
		http.Error(rw, "Bad request, could not get product id", http.StatusBadRequest)
		return
	}
	p.l.Info("Update Procut", zap.Int("id", productId))

	if data.GetProductIndexById(productId) == -1 {
		p.l.Error("Failed to update produc, invalid product id.", zap.Int("id", productId))
		http.Error(rw, "Failed to update new product invlaid product id", http.StatusBadRequest)
		return

	}
	prod := r.Context().Value(ProductKey{}).(*data.Product)
	//we already validated the product id to be valie so we ignore the returned error here!
	// TODO: maybe we don't need the bool return of the update product!
	updatedProd, _ := data.UpdateProduct(prod, productId)

	rw.Header().Add("Content-Type", "application/json")
	data.ToJson(updatedProd, rw)
	p.l.Info("Update Products Response: ", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}
