package handlers

import (
	"log"
	"net/http"
	"simplems/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.listProducts(rw, r)
		return
	} else if r.Method == http.MethodPost {
		p.createProduct(rw, r)
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
