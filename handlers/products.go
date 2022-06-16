package handlers

import (
	"log"
	"net/http"
	"simplems/data"
	"strconv"
	"strings"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("serving request method ", r.Method)
	if r.Method == http.MethodGet {
		p.listProducts(rw, r)
		return
	} else if r.Method == http.MethodPost {
		p.createProduct(rw, r)
		return
	} else if r.Method == http.MethodPut {
		// updates and existing product with the id from the url
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 {
			p.l.Println("bad request no unique product id in route ", parts, r.URL.Path)
			http.Error(rw, "Bad request, could not get product id", http.StatusBadRequest)
			return
		}
		id := parts[len(parts)-1]
		productId, err := strconv.Atoi(id)
		if err != nil {
			p.l.Println("bad request no numeric product id in route ", parts, r.URL.Path)
			http.Error(rw, "Bad request, could not get product id", http.StatusBadRequest)
			return
		}
		p.l.Printf("updating product with id %d", productId)
		p.updateProduct(rw, r)
		return
	}
	// fallback
	// rw.WriteHeader(http.StatusMethodNotAllowed) <-- does not work...
	http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
}

func (p *Products) listProducts(rw http.ResponseWriter, r *http.Request) {
	pl := data.GetProductsList()
	err := pl.ToJson(rw)
	if err != nil {
		p.l.Println(err)
		http.Error(rw, "Failed to list products", http.StatusInternalServerError)
	}
	p.l.Printf("%s %s %s %d\n", r.RemoteAddr, r.Method, r.URL, http.StatusOK)
}

func (p *Products) createProduct(rw http.ResponseWriter, r *http.Request) {

	p1 := &data.Product{}

	errorToJson := p1.FromJson(r.Body)
	if errorToJson != nil {
		p.l.Println(errorToJson)
		http.Error(rw, "Failed to add new product could not decode json", http.StatusBadRequest)
	}
	data.AppnedPorduct(p1)
	p.l.Printf("%s %s %s %d\n", r.RemoteAddr, r.Method, r.URL, http.StatusOK)
	// pl := data.GetProductsList()
	// pl.ToJson(rw)
}

func (p *Products) updateProduct(rw http.ResponseWriter, r *http.Request) {
	p1 := &data.Product{}

	errorToJson := p1.FromJson(r.Body)
	if errorToJson != nil {
		p.l.Println(errorToJson)
		http.Error(rw, "Failed to add new product could not decode json", http.StatusBadRequest)
	}
	data.UpdateProduct(p1)
	p.l.Printf("%s %s %s %d\n", r.RemoteAddr, r.Method, r.URL, http.StatusOK)
}
