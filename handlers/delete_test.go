package handlers_test

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

type DeleteTestSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (suite *DeleteTestSuite) SetupSuite() {
	suite.logger, _ = zap.NewDevelopment()
}

func (suite *DeleteTestSuite) SetupTest() {
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

	for _, product := range *data.ProductsInstance().GetProductsList() {
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

func (suite *DeleteTestSuite) TestDeleteLastIndex() {
	count := data.ProductsInstance().Count()
	for i := count - 1; i >= 0; i-- {
		apitest.New().
			Report(apitest.SequenceDiagram()).
			Handler(app.NewApplication(suite.logger).Router).
			Delete(fmt.Sprintf("/products/%d/", i)).
			Expect(suite.T()).
			Status(http.StatusNoContent).
			End()
	}
}
