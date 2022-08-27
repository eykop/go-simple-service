package middlewares_test

import (
	"fmt"
	"math"
	"net/http"
	"simplems/app"
	"simplems/data"
	"testing"
	"time"

	"github.com/steinfletcher/apitest"
	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type MiddlewarseTestSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (suite *MiddlewarseTestSuite) SetupSuite() {
	suite.logger, _ = zap.NewDevelopment()
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

func (suite *MiddlewarseTestSuite) SetupTest() {

	for _, product := range *data.ProductsInstance().GetProductsList() {
		foo := product.(*data.Product)
		suite.logger.Debug("found product in list: ", zap.Int("id", foo.ID), zap.String("name", foo.Name))
	}

}

func (suite *MiddlewarseTestSuite) TearDownTest() {
}

func (suite *MiddlewarseTestSuite) TearDownSuite() {
	defer suite.logger.Sync()
}

func TestMiddlewarseSuite(t *testing.T) {
	suite.Run(t, new(MiddlewarseTestSuite))
}

func (suite *MiddlewarseTestSuite) TestValidProductIdWithDeletePass() {

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Delete("/products/0/").
		Expect(suite.T()).
		Body("").
		Status(http.StatusNoContent).
		End()
}

func (suite *MiddlewarseTestSuite) TestBadProductIdWithDelete() {

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Delete("/products/200/").
		Expect(suite.T()).
		Body("Failed to `DELETE` product, invalid product id.\n").
		Status(http.StatusBadRequest).
		End()
}

func (suite *MiddlewarseTestSuite) TestProductIdExceedsLimitWithDelete() {
	const exceeded_id = math.MaxInt64 + 1
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Delete(fmt.Sprintf("/products/%d/", uint64(exceeded_id))).
		Expect(suite.T()).
		Body("Invalid product id.\n").
		Status(http.StatusBadRequest).
		End()
}

func (suite *MiddlewarseTestSuite) TestStringProductIdWithDelete() {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Delete(fmt.Sprintf("/products/%s/", "strid")).
		Expect(suite.T()).
		Body("404 page not found\n").
		Status(http.StatusNotFound).
		End()
}

func (suite *MiddlewarseTestSuite) TestGetProductBadProductId() {

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Get("/products/200/").
		Expect(suite.T()).
		Body("Failed to `GET` product, invalid product id.\n").
		Status(http.StatusBadRequest).
		End()
}

func (suite *MiddlewarseTestSuite) TestGetProductProductIdExceedsLimit() {

	const exceeded_id = math.MaxInt64 + 1
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Get(fmt.Sprintf("/products/%d/", uint64(exceeded_id))).
		Expect(suite.T()).
		Body("Invalid product id.\n").
		Status(http.StatusBadRequest).
		End()
}

func (suite *MiddlewarseTestSuite) TestGetProductStringId() {

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Get(fmt.Sprintf("/products/%s/", "strid")).
		Expect(suite.T()).
		Body("404 page not found\n").
		Status(http.StatusNotFound).
		End()
}

func (suite *MiddlewarseTestSuite) TestGetProductList() {

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

func TestCreateProductPass(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(logger).Router).
		Post("/products/").
		Body(`{"name":"Despresso3","price":1.0, "sku": "abs-nfg-poe"}`).
		Expect(t).
		Body(`{
			"id": 2,
			"name": "Despresso3",
			"descripton": "",
			"price": 1,
			"sku": "abs-nfg-poe"
		}
		`).
		Status(http.StatusOK).
		End()
}
