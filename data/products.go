package data

import (
	"fmt"
	"time"
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

func AddPorduct(p *Product) {
	p.ID = getNextProductId()
	productsList = append(productsList, p)
}

func UpdateProduct(incommingProd *Product, index int) error {
	//add logger
	if index < 0 || index > len(productsList)-1 {

		return fmt.Errorf("update error, product index %d out of range", index)
	}

	p := productsList[index]
	if incommingProd.Name != "" {
		p.Name = incommingProd.Name
	}
	if incommingProd.Description != "" {
		p.Description = incommingProd.Description
	}
	if incommingProd.SKU != "" {
		p.SKU = incommingProd.SKU
	}
	if incommingProd.Price > 0 {
		p.Price = incommingProd.Price
	}
	p.UpdatedOn = time.Now().UTC().String()
	return nil
}

func DeleteProduct(index int) error {
	// todo add logger l.Info("Delete Procut", zap.Int("id", productId))

	if index < 0 || index > len(productsList)-1 {
		return fmt.Errorf("deletion error, product index %d out of range", index)
	}
	if index == len(productsList) {
		productsList = productsList[:index]
	} else {
		productsList = append(productsList[:index], productsList[index+1:]...)
	}
	return nil

}

var productsList = []*Product{
	{
		ID:          0,
		Name:        "Espresso",
		Description: "Lite coffe drink...",
		Price:       1.49,
		SKU:         "5faf1ada-5d01-4831-aa0c-8f93eec9d86e",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          1,
		Name:        "Latte",
		Description: "Lite coffe drink with milk...",
		Price:       2.49,
		SKU:         "a345d9d6-0c08-45a2-887a-4c22594737b3",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
