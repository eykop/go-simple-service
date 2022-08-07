package handlers_test

import (
	"fmt"
	"math"
	"net/http"
	"simplems/app"
	"testing"

	"github.com/steinfletcher/apitest"
	_ "github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestDeleteProductPass(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(logger).Router).
		Delete("/products/0/").
		Expect(t).
		Body("").
		Status(http.StatusNoContent).
		End()
}

func TestDeleteProductBadProductId(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(logger).Router).
		Delete("/products/200/").
		Expect(t).
		Body("Failed to `DELETE` product, invalid product id.\n").
		Status(http.StatusBadRequest).
		End()
}

func TestDeleteProductProductIdExceedsLimit(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	const exceeded_id = math.MaxInt64 + 1
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(logger).Router).
		Delete(fmt.Sprintf("/products/%d/", uint64(exceeded_id))).
		Expect(t).
		Body("Invalid product id.\n").
		Status(http.StatusBadRequest).
		End()
}

func TestDeleteProductStringId(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(logger).Router).
		Delete(fmt.Sprintf("/products/%s/", "strid")).
		Expect(t).
		Body("404 page not found\n").
		Status(http.StatusNotFound).
		End()
}
