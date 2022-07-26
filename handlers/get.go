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

	pl := data.GetProductsList()
	rw.Header().Add("Content-Type", "application/json")
	// rw.WriteHeader(http.StatusOK)
	err := data.ToJson(pl, rw)
	if err != nil {
		p.l.Error("Failed to List product", zap.Error(err))
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

	productId := getProductId(r, p.l)
	if productId == -1 {
		http.Error(rw, "Bad request, could not get product id", http.StatusBadRequest)
		return
	}
	p.l.Info("Get Procut", zap.Int("id", productId))
	pi := data.GetProductIndexById(productId)
	if pi == -1 {
		http.Error(rw, "Failed to get new product invlaid product id", http.StatusBadRequest)
		return
	}

	pl := data.GetProductsList()
	prod := pl[pi]
	rw.Header().Add("Content-Type", "application/json")
	prodErr := data.ToJson(prod, rw)
	if prodErr != nil {
		p.l.Error("Failed to get product", zap.Error(prodErr))
		http.Error(rw, "Failed to get product", http.StatusInternalServerError)
	}
	p.l.Info("Get Product Response", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}
