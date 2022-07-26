package data

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"descripton"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func GetProductsList() Products {
	return productsList
}

func AppnedPorduct(p *Product) {
	p.ID = getNextProductId()
	productsList = append(productsList, p)
}

func UpdateProduct(updatedProductData *Product, id int) (*Product, bool) {
	pi := GetProductIndexById(id)
	if pi == -1 {
		return nil, false
	}
	p := productsList[pi]
	if updatedProductData.Name != "" {
		p.Name = updatedProductData.Name
	}
	if updatedProductData.Description != "" {
		p.Description = updatedProductData.Description
	}
	if updatedProductData.SKU != "" {
		p.SKU = updatedProductData.SKU
	}
	if updatedProductData.Price > 0 {
		p.Price = updatedProductData.Price
	}
	p.UpdatedOn = time.Now().UTC().String()
	return p, true
}

func DeleteProduct(id int) error {
	if id < 0 || id > len(productsList)-1 {
		return fmt.Errorf("Product Id %d not found", id)
	}
	if id == len(productsList) {
		productsList = productsList[:id]
	} else {
		productsList = append(productsList[:id], productsList[id+1:]...)
	}
	return nil

}

var productsList = []*Product{
	{
		ID:          0,
		Name:        "Espresso",
		Description: "Lite coffe drink...",
		Price:       1.49,
		SKU:         uuid.New().String(),
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          1,
		Name:        "Latte",
		Description: "Lite coffe drink with milk...",
		Price:       2.49,
		SKU:         uuid.New().String(),
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
