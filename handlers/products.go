package handlers

import (
	"log"
	"net/http"
	"simplems/data"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ListProducts(rw http.ResponseWriter, r *http.Request) {

	pl := data.GetProductsList()
	err := pl.ToJson(rw)
	if err != nil {
		p.l.Println(err)
		http.Error(rw, "Failed to list products", http.StatusInternalServerError)
	}
	p.l.Printf("%s %s %s %d\n", r.RemoteAddr, r.Method, r.URL, http.StatusOK)
}

func (p *Products) CreateProduct(rw http.ResponseWriter, r *http.Request) {

	p1 := JsonToProduct(r, p)
	if p1 == nil {
		http.Error(rw, "Failed to deserialize product from json", http.StatusBadRequest)
		return
	}
	data.AppnedPorduct(p1)
	p.l.Printf("%s %s %s %d\n", r.RemoteAddr, r.Method, r.URL, http.StatusOK)
}

func JsonToProduct(r *http.Request, prods *Products) *data.Product {
	p := &data.Product{}

	errorToJson := p.FromJson(r.Body)
	if errorToJson != nil {
		prods.l.Println(errorToJson)
		return nil
	}
	return p
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p.l.Println("bad request no numeric product vars: ", vars)
	id := vars["id"]

	productId, err := strconv.Atoi(id)
	if err != nil {
		p.l.Println("bad request no numeric product id in route ,id is: ", id)
		p.l.Println("bad request no numeric product id in route , route is: ", r.URL.Path)
		http.Error(rw, "Bad request, could not get product id", http.StatusBadRequest)
		return
	}
	p.l.Printf("updating product with id %d", productId)

	p1 := JsonToProduct(r, p)
	if p1 == nil {
		http.Error(rw, "Failed to deserialize product from json", http.StatusBadRequest)
		return
	}
	if data.GetProductIndexById(productId) == -1 {
		p.l.Println("bad request invalid product id in route ,id is: ", id)
		http.Error(rw, "Failed to update new product invlaid product id", http.StatusBadRequest)
		return

	}
	data.UpdateProduct(p1, productId)
	p.l.Printf("%s %s %s %d\n", r.RemoteAddr, r.Method, r.URL, http.StatusOK)
}
