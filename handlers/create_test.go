package handlers_test

import (
	"net/http"
	"simplems/app"
	"testing"

	"github.com/steinfletcher/apitest"
	_ "github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateProductFailedEmptyJson(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(app.NewApplication(logger).Router).
		Post("/products/").
		Expect(t).
		Body("Failed to deserialize product from json.\n").
		Status(http.StatusBadRequest).
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
			"id": 0,
			"name": "Despresso3",
			"descripton": "",
			"price": 1,
			"sku": "abs-nfg-poe"
		}
		`).
		Status(http.StatusOK).
		End()
}
