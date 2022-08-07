package main

import (
	"simplems/client/client"
	"simplems/client/client/products"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestListProducts(t *testing.T) {
	logger, _ := zap.NewDevelopment()

	defer logger.Sync()

	dc := client.Default
	prods, err := dc.Products.GetProducts(products.NewGetProductsParams())
	if err != nil {
		logger.Error("Failed to List product", zap.Error(err))
	}
	assert.Nil(t, err)
	assert.Equal(t, 2, len(prods.Payload))

	logger.Info(*prods.Payload[0].Name)
	logger.Info(*prods.Payload[1].Name)
	// assert.Error(t, err)
	// assert.ErrorContains(t, err, "Product.Name")
	// assert.ErrorContains(t, err, "required")
}
