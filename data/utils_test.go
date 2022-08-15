package data

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type UtilsTestSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (suite *UtilsTestSuite) SetupSuite() {
	suite.logger, _ = zap.NewDevelopment()
	productsList = &Products{
		&Product{
			ID:          0,
			Name:        "Espresso",
			Description: "Lite coffe drink...",
			Price:       1.49,
			SKU:         "5faf1ada-5d01-4831-aa0c-8f93eec9d86e",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
		&Product{
			ID:          3,
			Name:        "Latte",
			Description: "Lite coffe drink with milk...",
			Price:       2.49,
			SKU:         "a345d9d6-0c08-45a2-887a-4c22594737b3",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
	}
}

func (suite *UtilsTestSuite) SetupTest() {

}

func (suite *UtilsTestSuite) TearDownTest() {
}

func (suite *UtilsTestSuite) TearDownSuite() {
	defer suite.logger.Sync()
	productsList = &Products{
		&Product{
			ID:          0,
			Name:        "Espresso",
			Description: "Lite coffe drink...",
			Price:       1.49,
			SKU:         "5faf1ada-5d01-4831-aa0c-8f93eec9d86e",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
		&Product{
			ID:          3,
			Name:        "Latte",
			Description: "Lite coffe drink with milk...",
			Price:       2.49,
			SKU:         "a345d9d6-0c08-45a2-887a-4c22594737b3",
			CreatedOn:   time.Now().UTC().String(),
			UpdatedOn:   time.Now().UTC().String(),
		},
	}
}

func TestUtilsSuite(t *testing.T) {
	suite.Run(t, new(UtilsTestSuite))
}

func (suite *UtilsTestSuite) TestNextProductId() {
	assert.Equal(suite.T(), 4, getNextProductId())
}

func (suite *UtilsTestSuite) TestGetProductIndexById() {
	assert.Equal(suite.T(), 0, GetProductIndexById(0))
	assert.Equal(suite.T(), 1, GetProductIndexById(3))
}

func (suite *UtilsTestSuite) TestFailGetProductIndexById() {
	assert.Equal(suite.T(), -1, GetProductIndexById(1))
}

func (suite *UtilsTestSuite) TestToJson() {
	var strBuffer bytes.Buffer
	(*productsList)[0].(*Product).ToJson(&strBuffer)
	assert.Equal(suite.T(), `{"id":0,"name":"Espresso","descripton":"Lite coffe drink...","price":1.49,"sku":"5faf1ada-5d01-4831-aa0c-8f93eec9d86e"}`+"\n", strBuffer.String())
}

func (suite *UtilsTestSuite) TestFromJson() {
	prod := &Product{}
	var strBuffer bytes.Buffer
	strBuffer.WriteString(`{"id":0,"name":"Espresso","descripton":"Lite coffe drink...","price":1.49,"sku":"5faf1ada-5d01-4831-aa0c-8f93eec9d86e"}`)
	// fromJson(prod, &strBuffer)
	// assert.Equal(suite.T(), "Espresso", prod.Name)
	err := prod.FromJson(&strBuffer)
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Espresso", prod.Name)
}

func (suite *UtilsTestSuite) TestFromJsonFail() {
	prod := &Product{}
	var strBuffer bytes.Buffer
	strBuffer.WriteString(`not json`)
	assert.Error(suite.T(), prod.FromJson(&strBuffer))
}
