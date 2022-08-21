package data

import (
	"io"
	"time"

	"github.com/go-playground/validator/v10"
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

func (p *Product) GetID() int {
	return p.ID
}

func (p *Product) SetID(id int) {
	p.ID = id
}

func (p *Product) ToJson(w io.Writer) error {
	return toJson(p, w)
}

func (p *Product) FromJson(r io.Reader) error {
	return fromJson(p, r)
}

func (p *Product) UpdateProduct(updated ProductInterface) {

	updatedProduct := updated.(*Product)
	if updatedProduct.Name != "" {
		p.Name = updatedProduct.Name
	}
	if updatedProduct.Description != "" {
		p.Description = updatedProduct.Description
	}
	if updatedProduct.SKU != "" {
		p.SKU = updatedProduct.SKU
	}
	if updatedProduct.Price > 0 {
		p.Price = updatedProduct.Price
	}
	p.UpdatedOn = time.Now().UTC().String()
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}
