package handlers_test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"simplems/app"
	"simplems/data"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
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
	count := data.ProductsInstance().Count()
	for i := count - 1; i >= 0; i-- {
		data.ProductsInstance().DeleteProduct(i)
	}

	data.ProductsInstance().AddPorduct(&data.Product{
		ID:          0,
		Name:        "Espresso",
		Description: "Lite coffe drink...",
		Price:       1.49,
		SKU:         "5faf1ada-5d01-4831-aa0c-8f93eec9d86e",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	})
	data.ProductsInstance().AddPorduct(&data.Product{
		ID:          1,
		Name:        "Latte",
		Description: "Lite coffe drink with milk...",
		Price:       2.49,
		SKU:         "a345d9d6-0c08-45a2-887a-4c22594737b3",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	})
}

func (suite *GetTestSuite) logProducts() {
	ls := data.ProductsInstance().GetProductsList()
	for i, p := range *ls {
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
	assert.Equal(suite.T(), 2, len(*data.ProductsInstance().GetProductsList()))

	data.ProductsInstance().AddPorduct(&DummyProduct{})
	assert.Equal(suite.T(), 3, len(*data.ProductsInstance().GetProductsList()))
	suite.logProducts()
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Get("/products/2/").
		Expect(suite.T()).
		Body("Failed to encode product.\n").
		Status(http.StatusInternalServerError).
		End()

	data.ProductsInstance().DeleteProduct(2)
	assert.NotEmpty(suite.T(), data.ProductsInstance().GetProductsList())
	assert.Equal(suite.T(), 2, len(*data.ProductsInstance().GetProductsList()))
}

func (suite *GetTestSuite) TestGetProductList() {

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Get("/products/").
		Expect(suite.T()).
		Body(`[
			{
				"descripton":"Lite coffe drink...", 
				"id":0, "name":"Espresso", 
				"price":1.49, 
				"sku":"5faf1ada-5d01-4831-aa0c-8f93eec9d86e"
			},
			{
				"descripton":"Lite coffe drink with milk...", 
				"id":1, "name":"Latte",
				 "price":2.49, 
				 "sku":"a345d9d6-0c08-45a2-887a-4c22594737b3"
			}
				]`).
		Status(http.StatusOK).
		End()
}

func (suite *GetTestSuite) TestGetProductListFail() {
	count := data.ProductsInstance().Count()
	for index := 0; index < count; index++ {
		data.ProductsInstance().DeleteProduct(0)
	}
	count = data.ProductsInstance().Count()
	assert.Equal(suite.T(), 0, count)
	ctrlr := gomock.NewController(suite.T())
	defer ctrlr.Finish()

	mock := &MockProductsInterface{ctrl: ctrlr}
	mock.recorder = &MockProductsInterfaceMockRecorder{mock}
	data.InitProducts(mock)

	mock.
		EXPECT().
		ToJson(gomock.Any()).
		Return(errors.New("Error tojson")).
		AnyTimes()

	mock.
		EXPECT().
		GetProductsList().Return(&data.ProductsList{}).AnyTimes()

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Get("/products/").
		Expect(suite.T()).
		Body("Failed to list products\n").
		Status(http.StatusInternalServerError).
		End()
}

func TestGetProductSuite(t *testing.T) {
	suite.Run(t, new(GetTestSuite))
}
