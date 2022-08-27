package data

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ProductTestSuite struct {
	suite.Suite
	logger             *zap.Logger
	initialProductSize int
	products           *Products
}

func (suite *ProductTestSuite) SetupSuite() {
	suite.logger, _ = zap.NewDevelopment()
	suite.products = &Products{logger: suite.logger}

}

func (suite *ProductTestSuite) SetupTest() {

	suite.products.SetProducts(&ProductsList{
		&Product{
			ID:          0,
			Name:        "Espresso",
			Description: "Lite coffe drink...",
			Price:       1.49,
			SKU:         "5faf1ada-5d01-4831-aa0c-8f93eec9d86e",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
		&Product{
			ID:          2,
			Name:        "Arabica",
			Description: "Arabica coffe drink without milk...",
			Price:       5.99,
			SKU:         "a345d9d6-0c08-45a2-887a-4c2259f4i7n3",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
		&Product{
			ID:          3,
			Name:        "Latte",
			Description: "Lite coffe drink with milk...",
			Price:       2.49,
			SKU:         "a345d9d6-0c08-45a2-887a-4c22594737b3",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
	})
	suite.initialProductSize = suite.products.Count()
	assert.Equal(suite.T(), 3, suite.initialProductSize)

}

func (suite *ProductTestSuite) TearDownTest() {
}

func (suite *ProductTestSuite) TearDownSuite() {
	defer suite.logger.Sync()
	suite.products.SetProducts(&ProductsList{
		&Product{
			ID:          0,
			Name:        "Espresso",
			Description: "Lite coffe drink...",
			Price:       1.49,
			SKU:         "5faf1ada-5d01-4831-aa0c-8f93eec9d86e",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
		&Product{
			ID:          3,
			Name:        "Latte",
			Description: "Lite coffe drink with milk...",
			Price:       2.49,
			SKU:         "a345d9d6-0c08-45a2-887a-4c22594737b3",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
	})
}

func TestProductSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}

func (suite *ProductTestSuite) TestGetProductList() {
	assert.Equal(suite.T(), suite.initialProductSize, suite.products.Count())
}

func (suite *ProductTestSuite) TestAddPorduct() {
	prod := &Product{Name: "No 7", Price: 7}
	suite.products.AddPorduct(prod)
	assert.Equal(suite.T(), suite.initialProductSize+1, suite.products.Count())
}

func (suite *ProductTestSuite) TestUpdateProductAllFields() {
	prod := &Product{Name: "No 7", Price: 7, Description: "New Desciption", SKU: "sss-sss-sss"}
	err := suite.products.UpdateProduct(prod, 0)
	assert.NoError(suite.T(), err)

	upProd := (*suite.products.GetProductsList())[0].(*Product)

	//check the product return by update
	assert.Equal(suite.T(), prod.Name, upProd.Name)
	assert.Equal(suite.T(), prod.Price, upProd.Price)
	assert.Equal(suite.T(), prod.Description, upProd.Description)
	assert.Equal(suite.T(), prod.SKU, upProd.SKU)
	assert.NotEqual(suite.T(), prod.UpdatedOn, upProd.UpdatedOn)

}

func (suite *ProductTestSuite) TestGetProductByIndexOutOfRange() {
	//index is larger than list items...
	assert.Nil(suite.T(), suite.products.GetProductByIndex(suite.products.Count()))

	//index is less than zero
	assert.Nil(suite.T(), suite.products.GetProductByIndex(-1))
}

func (suite *ProductTestSuite) TestUpdateProductPartialFields() {
	prod := &Product{Name: "No 7", SKU: "sss-sss-sss"}
	orgProdPrice := suite.products.GetProductByIndex(0).(*Product).Price
	orgProdDesc := suite.products.GetProductByIndex(0).(*Product).Description
	err := suite.products.UpdateProduct(prod, 0)
	assert.NoError(suite.T(), err)

	upProd := suite.products.GetProductByIndex(0).(*Product)

	//check the product return by update
	assert.Equal(suite.T(), prod.Name, upProd.Name)
	assert.Equal(suite.T(), prod.SKU, upProd.SKU)
	assert.NotEqual(suite.T(), prod.UpdatedOn, upProd.UpdatedOn)

	//not updated fields
	assert.Equal(suite.T(), orgProdPrice, upProd.Price)
	assert.Equal(suite.T(), orgProdDesc, upProd.Description)
}

func (suite *ProductTestSuite) TestUpdateProductBadIndex() {
	prod := &Product{Name: "No 7", SKU: "sss-sss-sss"}
	notFoundIndex := 7
	err := suite.products.UpdateProduct(prod, notFoundIndex)
	assert.Error(suite.T(), err)
}

func (suite *ProductTestSuite) TestDeleteProduct() {
	assert.NoError(suite.T(), suite.products.DeleteProduct(0))
	assert.Equal(suite.T(), suite.initialProductSize-1, suite.products.Count())
	assert.Equal(suite.T(), 2, suite.products.GetProductByIndex(0).(*Product).ID)
}

func (suite *ProductTestSuite) TestDeleteLastProduct() {
	assert.NoError(suite.T(), suite.products.DeleteProduct(suite.products.Count()-1))
	assert.Equal(suite.T(), suite.initialProductSize-1, suite.products.Count())
	assert.Equal(suite.T(), 0, suite.products.GetProductByIndex(0).(*Product).ID)
}

func (suite *ProductTestSuite) TestDeleteMiddleProduct() {
	assert.NoError(suite.T(), suite.products.DeleteProduct(suite.products.Count()/2))
	assert.Equal(suite.T(), suite.initialProductSize-1, suite.products.Count())
	assert.Equal(suite.T(), 0, suite.products.GetProductByIndex(0).(*Product).ID)
	assert.Equal(suite.T(), 3, suite.products.GetProductByIndex(len(*suite.products.GetProductsList())-1).(*Product).ID)
}

func (suite *ProductTestSuite) TestDeleteProductBadIndex() {
	assert.Error(suite.T(), suite.products.DeleteProduct(5))
}

func (suite *ProductTestSuite) TestNextProductsList() {

	for suite.products.Count() > 0 {
		suite.products.DeleteProduct(0)
	}
	assert.Equal(suite.T(), 0, suite.products.GetNextProductId())
}

func (suite *ProductTestSuite) TestProductsToJson() {
	strBuffer := bytes.Buffer{}
	assert.NoError(suite.T(), suite.products.ToJson(&strBuffer))
}
