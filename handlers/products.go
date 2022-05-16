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
