package handlers_test

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"simplems/app"
	"simplems/data"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type DummyProduct struct {
	ID int
}

func (p *DummyProduct) ToJson(w io.Writer) error {
	return fmt.Errorf("Failed to encode product.")
}

func (p *DummyProduct) FromJson(r io.Reader) error {
	return fmt.Errorf("Failed to decode product.")
}

func (p *DummyProduct) GetID() int {
	return p.ID
}

func (p *DummyProduct) GetName() string {
	return fmt.Sprintf("DummyProduct_%d", rand.Int())
}

func (p *DummyProduct) SetID(id int) {
	p.ID = id
}

func (p *DummyProduct) UpdateProduct(updated data.ProductInterface) {
}

func (p *DummyProduct) Validate() error {
	return fmt.Errorf("dummy validation error")
}

type GetTestSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (suite *GetTestSuite) SetupSuite() {
	suite.logger, _ = zap.NewDevelopment()
}

func (suite *GetTestSuite) SetupTest() {
	suite.logProducts()
}

func (suite *GetTestSuite) logProducts() {
	for i, p := range *data.GetProductsList() {
		if _, ok := p.(*data.Product); ok {
			suite.logger.Debug("product", zap.Int("id", p.GetID()), zap.Int("index", i), zap.String("type", "prod"))
		}
		if _, ok := p.(*DummyProduct); ok {
			suite.logger.Debug("product", zap.Int("id", p.GetID()), zap.Int("index", i), zap.String("type", "dummyprod"))
		}
	}
}

func (suite *GetTestSuite) TearDownTest() {
	suite.logProducts()
}

func (suite *GetTestSuite) TearDownSuite() {
	defer suite.logger.Sync()
}

func (suite *GetTestSuite) TestGetPass() {

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Get("/products/1/").
		Expect(suite.T()).
		Body(`{"descripton":"Lite coffe drink with milk...", "id":1, "name":"Latte", "price":2.49, "sku":"a345d9d6-0c08-45a2-887a-4c22594737b3"}`).
		Status(http.StatusOK).
		End()
}

func (suite *GetTestSuite) TestGetFail() {
	assert.Equal(suite.T(), 2, len(*data.GetProductsList()))

	data.AddPorduct(&DummyProduct{})
	assert.Equal(suite.T(), 3, len(*data.GetProductsList()))
	suite.logProducts()
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Get("/products/3/").
		Expect(suite.T()).
		Body("Failed to encode product.\n").
		Status(http.StatusInternalServerError).
		End()

	data.DeleteProduct(2)
	assert.NotEmpty(suite.T(), data.GetProductsList())
	assert.Equal(suite.T(), 2, len(*data.GetProductsList()))
}

func (suite *GetTestSuite) TestGetProductList() {

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Get("/products/").
		Expect(suite.T()).
		Body(`[
			{
				"descripton":"Lite coffe drink with milk...", 
				"id":1, 
				"name":"Latte", 
				"price":2.49, 
				"sku":"a345d9d6-0c08-45a2-887a-4c22594737b3"
			},
			{
				"descripton":"", 
				"id":2, 
				"name":"Despresso3", 
				"price":1, 
				"sku":"abs-nfg-poe"
			}
				]`).
		Status(http.StatusOK).
		End()
}

func TestGetProductSuite(t *testing.T) {
	suite.Run(t, new(GetTestSuite))
}
