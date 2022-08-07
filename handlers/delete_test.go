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

type DeleteTestSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (g *DeleteTestSuite) SetupSuite() {
	g.logger, _ = zap.NewDevelopment()
}

func (suite *DeleteTestSuite) SetupTest() {
	for _, product := range data.GetProductsList() {
		suite.logger.Debug("found product in list: ", zap.Int("id", product.ID), zap.String("name", product.Name))
	}
	//suite.Equal(2, len(data.GetProductsList()), "list of product not 2.")
}

func (s *DeleteTestSuite) TearDownTest() {
}

func (g *DeleteTestSuite) TearDownSuite() {
	defer g.logger.Sync()
}

func TestGDeleteProductSuite(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}

func (suite *DeleteTestSuite) TestDeletePass() {

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Delete("/products/0/").
		Expect(suite.T()).
		Body("").
		Status(http.StatusNoContent).
		End()
}

func (suite *DeleteTestSuite) TestDeleteBadId() {

	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Delete("/products/200/").
		Expect(suite.T()).
		Body("Failed to `DELETE` product, invalid product id.\n").
		Status(http.StatusBadRequest).
		End()
}

func (suite *DeleteTestSuite) TestDeleteIdExceedsLimit() {
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

func (suite *DeleteTestSuite) TestDeleteStringId() {
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(suite.logger).Router).
		Delete(fmt.Sprintf("/products/%s/", "strid")).
		Expect(suite.T()).
		Body("404 page not found\n").
		Status(http.StatusNotFound).
		End()
}
