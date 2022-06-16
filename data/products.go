package data

import (
	"encoding/json"
	"io"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"descripton"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func GetProductsList() Products {
	return productsList
}

func AppnedPorduct(p *Product) {
	p.ID = GetNextProductId()
	productsList = append(productsList, p)
}

func UpdateProduct(newPorcut *Product) bool {
	pi := getProductIndexById(newPorcut.ID)
	if pi == -1 {
		return false
	}
	p := productsList[pi]
	if newPorcut.Name != "" {
		p.Name = newPorcut.Name
	}
	if newPorcut.Description != "" {
		p.Description = newPorcut.Description
	}
	if newPorcut.SKU != "" {
		p.SKU = newPorcut.SKU
	}
	if newPorcut.Price > 0 {
		p.Price = newPorcut.Price
	}
	p.UpdatedOn = time.Now().UTC().String()
	return true
}

func GetNextProductId() int {
	lastProduct := productsList[len(productsList)-1]
	return lastProduct.ID + 1
}

func getProductIndexById(id int) int {
	for index, product := range productsList {
		if product.ID == id {
			return index
		}
	}
	return -1
}

var productsList = []*Product{
	&Product{
		ID:          0,
		Name:        "Espresso",
		Description: "Lite coffe drink...",
		Price:       1.49,
		SKU:         uuid.New().String(),
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Lite coffe drink with milk...",
		Price:       2.49,
		SKU:         uuid.New().String(),
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
