package handlers

import (
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

func (p *Products) ListProducts(rw http.ResponseWriter, r *http.Request) {

	pl := data.GetProductsList()
	err := pl.ToJson(rw)
	if err != nil {
		p.l.Error("Failed to List product", zap.Error(err))
		http.Error(rw, "Failed to list products", http.StatusInternalServerError)
	}
	p.l.Info("List Products Response", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}

func (p *Products) CreateProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Info("Will create a new product")
	p1 := JsonToProduct(r, p)
	if p1 == nil {
		p.l.Error("create product", zap.String("reason", "failed to decode json"))
		http.Error(rw, "Failed to deserialize product from json", http.StatusBadRequest)
		return
	}
	data.AppnedPorduct(p1)
	p.l.Info("Create Product Response: ", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}

func JsonToProduct(r *http.Request, prods *Products) *data.Product {
	p := &data.Product{}

	errorToJson := p.FromJson(r.Body)
	if errorToJson != nil {
		prods.l.Error("Failed to encode product", zap.Error(errorToJson))
		return nil
	}
	return p
}

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

	p1 := JsonToProduct(r, p)
	if p1 == nil {
		p.l.Error("Failed to update produc, could not decode product from json.", zap.Int("id", productId))
		http.Error(rw, "Failed to deserialize product from json.", http.StatusBadRequest)
		return
	}

	data.UpdateProduct(p1, productId)
	p.l.Info("Update Products Response: ", zap.String("remoteAddr", r.RemoteAddr), zap.String("method", r.Method), zap.String("url", r.URL.Path), zap.Int("status", http.StatusOK))
}
