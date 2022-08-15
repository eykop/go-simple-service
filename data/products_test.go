package data

import (
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
}

func (suite *ProductTestSuite) SetupSuite() {
	suite.logger, _ = zap.NewDevelopment()

}

func (suite *ProductTestSuite) SetupTest() {

	productsList = &Products{
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
	}
	suite.initialProductSize = len(*GetProductsList())
	assert.Equal(suite.T(), 3, suite.initialProductSize)

}

func (suite *ProductTestSuite) TearDownTest() {
}

func (suite *ProductTestSuite) TearDownSuite() {
	defer suite.logger.Sync()
	productsList = &Products{
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
	}
}

func TestProductSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}

func (suite *ProductTestSuite) TestGetProductList() {
	assert.Equal(suite.T(), suite.initialProductSize, ProductsCount())
}

func (suite *ProductTestSuite) TestAddPorduct() {
	prod := &Product{Name: "No 7", Price: 7}
	AddPorduct(prod)
	assert.Equal(suite.T(), suite.initialProductSize+1, ProductsCount())
}

func (suite *ProductTestSuite) TestUpdateProductAllFields() {
	prod := &Product{Name: "No 7", Price: 7, Description: "New Desciption", SKU: "sss-sss-sss"}
	err := UpdateProduct(prod, 0)
	assert.NoError(suite.T(), err)

	upProd := (*GetProductsList())[0].(*Product)

	//check the product return by update
	assert.Equal(suite.T(), prod.Name, upProd.Name)
	assert.Equal(suite.T(), prod.Price, upProd.Price)
	assert.Equal(suite.T(), prod.Description, upProd.Description)
	assert.Equal(suite.T(), prod.SKU, upProd.SKU)
	assert.NotEqual(suite.T(), prod.UpdatedOn, upProd.UpdatedOn)

}

func (suite *ProductTestSuite) TestUpdateProductPartialFields() {
	prod := &Product{Name: "No 7", SKU: "sss-sss-sss"}
	orgProdPrice := GetProductByIndex(0).(*Product).Price
	orgProdDesc := GetProductByIndex(0).(*Product).Description
	err := UpdateProduct(prod, 0)
	assert.NoError(suite.T(), err)

	upProd := GetProductByIndex(0).(*Product)

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
	err := UpdateProduct(prod, notFoundIndex)
	assert.Error(suite.T(), err)
}

func (suite *ProductTestSuite) TestDeleteProduct() {
	assert.NoError(suite.T(), DeleteProduct(0))
	assert.Equal(suite.T(), suite.initialProductSize-1, ProductsCount())
	assert.Equal(suite.T(), 2, GetProductByIndex(0).(*Product).ID)
}

func (suite *ProductTestSuite) TestDeleteLastProduct() {
	assert.NoError(suite.T(), DeleteProduct(ProductsCount()-1))
	assert.Equal(suite.T(), suite.initialProductSize-1, ProductsCount())
	assert.Equal(suite.T(), 0, GetProductByIndex(0).(*Product).ID)
}

func (suite *ProductTestSuite) TestDeleteMiddleProduct() {
	assert.NoError(suite.T(), DeleteProduct(ProductsCount()/2))
	assert.Equal(suite.T(), suite.initialProductSize-1, ProductsCount())
	assert.Equal(suite.T(), 0, GetProductByIndex(0).(*Product).ID)
	assert.Equal(suite.T(), 3, GetProductByIndex(len(*GetProductsList())-1).(*Product).ID)
}

func (suite *ProductTestSuite) TestDeleteProductBadIndex() {
	assert.Error(suite.T(), DeleteProduct(5))
}
