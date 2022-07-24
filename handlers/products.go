package handlers

import (
	"fmt"
	"net/http"
	"simplems/data"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Products struct {
	l *zap.Logger
}

func NewProducts(l *zap.Logger) *Products {
	return &Products{l}
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

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
	err := pl.ToJson(rw)
	if err != nil {
		p.l.Error("Failed to List product", zap.Error(err))
		http.Error(rw, "Failed to list products", http.StatusInternalServerError)
	}
	p.l.Info("List Products Response", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}

type ProductKey struct{}

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
	prod := r.Context().Value(ProductKey{}).(*data.Product)
	data.AppnedPorduct(prod)
	rw.Header().Add("Content-Type", "application/json")
	prod.ToJson(rw)
	p.l.Info("Create Product Response: ", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}

func JsonToProduct(r *http.Request, l *zap.Logger) *data.Product {
	p := &data.Product{}
	errorToJson := p.FromJson(r.Body)
	if errorToJson != nil {
		l.Error("Failed to encode product", zap.Error(errorToJson))
		return nil
	}
	return p
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
	vars := mux.Vars(r)
	id := vars["id"]

	productId, err := strconv.Atoi(id)
	if err != nil {
		p.l.Error("Failed to get product", zap.Error(err))
		http.Error(rw, "Bad request, could not get product id", http.StatusBadRequest)
		return
	}
	p.l.Info("Get Procut", zap.Int("id", productId))
	pi := data.GetProductIndexById(productId)
	if pi == -1 {
		p.l.Error("Failed to update produc, invalid product id.", zap.Int("id", productId))
		http.Error(rw, "Failed to update new product invlaid product id", http.StatusBadRequest)
		return
	}

	pl := data.GetProductsList()
	prod := pl[pi]
	rw.Header().Add("Content-Type", "application/json")
	prodErr := prod.ToJson(rw)
	if prodErr != nil {
		p.l.Error("Failed to get product", zap.Error(err))
		http.Error(rw, "Failed to get product", http.StatusInternalServerError)
	}
	p.l.Info("Get Product Response", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete by product ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"  Format(int64)
// @Success      204
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /products/{id} [delete]
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	productId, err := strconv.Atoi(id)
	if err != nil {
		p.l.Error("Failed to delete product", zap.Error(err))
		http.Error(rw, "Bad request, could not get product id", http.StatusBadRequest)
		return
	}
	p.l.Info("Delete Procut", zap.Int("id", productId))
	delErr := data.DeleteProductByID(productId)
	if delErr != nil {
		p.l.Error("Failed to delete product", zap.Error(delErr))
		http.Error(rw, fmt.Sprintf("Bad request, could not delete product %v", delErr), http.StatusBadRequest)
		return
	}

}

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
	vars := mux.Vars(r)
	id := vars["id"]

	productId, err := strconv.Atoi(id)
	if err != nil {
		p.l.Error("Failed to update product", zap.Error(err))
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
	pp, _ := data.UpdateProduct(prod, productId)

	rw.Header().Add("Content-Type", "application/json")
	pp.ToJson(rw)
	p.l.Info("Update Products Response: ", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}
