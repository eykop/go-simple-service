package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationErrorMissingProductName(t *testing.T) {
	p := Product{}

	p.SKU = "as-qwe-rty"
	p.Price = 1.00
	err := p.Validate()

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Product.Name")
	assert.ErrorContains(t, err, "required")
}

func TestValidationErrorMissingProducPrice(t *testing.T) {
	p := Product{}

	p.SKU = "as-qwe-rty"
	p.Name = "prod1"
	err := p.Validate()

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Product.Price")
	assert.ErrorContains(t, err, "validation")
}

func TestValidationErrorBadProducPrice(t *testing.T) {
	p := Product{}

	p.SKU = "as-qwe-rty"
	p.Name = "prod1"
	p.Price = -1
	err := p.Validate()

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Product.Price")
	assert.ErrorContains(t, err, "validation")
}

func TestValidationErrorMissingProductSKU(t *testing.T) {
	p := &Product{}

	p.Name = "prod1"
	p.Price = 1.00
	err := p.Validate()

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Product.SKU")
	assert.ErrorContains(t, err, "validation")
}

func TestValidationErrorBadProductSKU(t *testing.T) {
	p := &Product{}

	p.Name = "prod1"
	p.SKU = "badvalue"
	p.Price = 1.00

	err := p.Validate()

	assert.Error(t, err)
	assert.ErrorContains(t, err, "Product.SKU")
	assert.ErrorContains(t, err, "validation")
}

func TestValidationProductOK(t *testing.T) {
	p := &Product{}

	p.Name = "prod1"
	p.SKU = "bad-value-lal"
	p.Price = 1.00

	err := p.Validate()
	assert.Nil(t, err)
}
