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

func (suite *DeleteTestSuite) SetupSuite() {
	suite.logger, _ = zap.NewDevelopment()
}

func (suite *DeleteTestSuite) SetupTest() {
	for _, product := range *data.GetProductsList() {
		foo := product.(*data.Product)
		suite.logger.Debug("found product in list: ", zap.Int("id", foo.ID), zap.String("name", foo.Name))
	}
}

func (suite *DeleteTestSuite) TearDownTest() {
}

func (suite *DeleteTestSuite) TearDownSuite() {
	defer suite.logger.Sync()
}

func TestDeleteProductSuite(t *testing.T) {
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
