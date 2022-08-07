package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ValidationsTestSuite struct {
	suite.Suite
	logger *zap.Logger
}

func (suite *ValidationsTestSuite) SetupSuite() {
	suite.logger, _ = zap.NewDevelopment()
}

func (suite *ValidationsTestSuite) SetupTest() {

}

func (suite *ValidationsTestSuite) TearDownTest() {
}

func (suite *ValidationsTestSuite) TearDownSuite() {
	defer suite.logger.Sync()
}

func TestValidationsSuite(t *testing.T) {
	suite.Run(t, new(ValidationsTestSuite))
}

func (suite *ValidationsTestSuite) TestValidationMissingName() {
	p := Product{}

	p.SKU = "as-qwe-rty"
	p.Price = 1.00
	err := p.Validate()

	assert.Error(suite.T(), err)
	assert.ErrorContains(suite.T(), err, "Product.Name")
	assert.ErrorContains(suite.T(), err, "required")
}

func (suite *ValidationsTestSuite) TestValidationMissingPrice() {
	p := Product{}

	p.SKU = "as-qwe-rty"
	p.Name = "prod1"
	err := p.Validate()

	assert.Error(suite.T(), err)
	assert.ErrorContains(suite.T(), err, "Product.Price")
	assert.ErrorContains(suite.T(), err, "validation")
}

func (suite *ValidationsTestSuite) TestValidationBadPrice() {
	p := Product{}

	p.SKU = "as-qwe-rty"
	p.Name = "prod1"
	p.Price = -1
	err := p.Validate()

	assert.Error(suite.T(), err)
	assert.ErrorContains(suite.T(), err, "Product.Price")
	assert.ErrorContains(suite.T(), err, "validation")
}

func (suite *ValidationsTestSuite) TestValidationMissingSku() {
	p := &Product{}

	p.Name = "prod1"
	p.Price = 1.00
	err := p.Validate()

	assert.Error(suite.T(), err)
	assert.ErrorContains(suite.T(), err, "Product.SKU")
	assert.ErrorContains(suite.T(), err, "validation")
}

func (suite *ValidationsTestSuite) TestValidationBadSku() {
	p := &Product{}

	p.Name = "prod1"
	p.SKU = "badvalue"
	p.Price = 1.00

	err := p.Validate()

	assert.Error(suite.T(), err)
	assert.ErrorContains(suite.T(), err, "Product.SKU")
	assert.ErrorContains(suite.T(), err, "validation")
}

func (suite *ValidationsTestSuite) TestValidationOK() {
	p := &Product{}

	p.Name = "prod1"
	p.SKU = "bad-value-lal"
	p.Price = 1.00

	err := p.Validate()
	assert.Nil(suite.T(), err)
}
