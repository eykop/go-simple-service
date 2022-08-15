package middlewares_test

import (
	"fmt"
	"math"
	"net/http"
	"simplems/app"
	"simplems/data"
	"testing"

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
}

func (suite *MiddlewarseTestSuite) SetupTest() {
	for _, product := range *data.GetProductsList() {
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
