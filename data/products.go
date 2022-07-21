package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
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

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", ValidateSKU)
	return validate.Struct(p)
}

func ValidateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")
	return len(re.FindAllString(fl.Field().String(), -1)) == 1
}

func GetProductsList() Products {
	return productsList
}

func AppnedPorduct(p *Product) {
	p.ID = GetNextProductId()
	productsList = append(productsList, p)
}

func UpdateProduct(updatedProductData *Product, id int) bool {
	pi := GetProductIndexById(id)
	if pi == -1 {
		return false
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
	return true
}

func GetNextProductId() int {
	lastProduct := productsList[len(productsList)-1]
	return lastProduct.ID + 1
}

func GetProductIndexById(id int) int {
	for index, product := range productsList {
		if product.ID == id {
			return index
		}
	}
	return -1
}

func DeleteProductByID(id int) error {
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
