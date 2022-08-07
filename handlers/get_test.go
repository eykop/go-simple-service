package handlers_test

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

type GetTestSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (g *GetTestSuite) SetupSuite() {
	g.logger, _ = zap.NewDevelopment()
}

func (suite *GetTestSuite) SetupTest() {
	for _, product := range data.GetProductsList() {
		suite.logger.Debug("found product in list: ", zap.Int("id", product.ID), zap.String("name", product.Name))
	}
}

func (s *GetTestSuite) TearDownTest() {
}

func (g *GetTestSuite) TearDownSuite() {
	defer g.logger.Sync()
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

func TestGetProductSuite(t *testing.T) {
	suite.Run(t, new(GetTestSuite))
}

func (suite *GetTestSuite) TestGetProductBadProductId() {

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Get("/products/200/").
		Expect(suite.T()).
		Body("Failed to `GET` product, invalid product id.\n").
		Status(http.StatusBadRequest).
		End()
}

func (suite *GetTestSuite) TestGetProductProductIdExceedsLimit() {

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

func (suite *GetTestSuite) TestGetProductStringId() {

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Get(fmt.Sprintf("/products/%s/", "strid")).
		Expect(suite.T()).
		Body("404 page not found\n").
		Status(http.StatusNotFound).
		End()
}
