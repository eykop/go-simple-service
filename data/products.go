package data

import (
	"fmt"
	"io"

	"go.uber.org/zap"
)

type ProductInterface interface {
	ToJson(w io.Writer) error
	FromJson(r io.Reader) error
	GetID() int
	SetID(id int)
	UpdateProduct(updated ProductInterface)
	Validate() error
}

type ProductsList []ProductInterface

type ProductsInterface interface {
	Count() int
	GetProductByIndex(index int) ProductInterface
	GetProductsList() *ProductsList
	AddPorduct(p ProductInterface)
	UpdateProduct(incommingProd ProductInterface, index int) error
	DeleteProduct(index int) error
	GetProductIndexById(id int) int
	GetNextProductId() int
	ToJson(w io.Writer) error
}

type Products struct {
	productsList *ProductsList
	logger       *zap.Logger
}

func NewProducts() ProductsInterface {
	return &Products{logger: &zap.Logger{}, productsList: &ProductsList{}}
}

func (products *Products) SetProducts(pl *ProductsList) {

	products.productsList = pl
}

func (products *Products) GetProductsList() *ProductsList {
	return products.productsList
}

func (products *Products) Count() int {
	return len(*products.productsList)
}

func (products *Products) GetProductByIndex(index int) ProductInterface {
	if index < 0 || index > products.Count()-1 {
		products.logger.Error("update error, product index out of range", zap.Int("index", index))
		return nil
	}
	return (*products.productsList)[index]
}

func (products *Products) AddPorduct(p ProductInterface) {
	p.SetID(products.GetNextProductId())
	*products.productsList = append((*products.productsList), p)
}

func (products *Products) UpdateProduct(incommingProd ProductInterface, index int) error {
	//add logger
	if index < 0 || index > products.Count()-1 {

		return fmt.Errorf("update error, product index %d out of range", index)
	}
	p := (*products.productsList)[index]
	p.UpdateProduct(incommingProd)
	return nil
}

func (products *Products) DeleteProduct(index int) error {
	// todo add logger l.Info("Delete Procut", zap.Int("id", productId))

	if index < 0 || index > products.Count()-1 {
		return fmt.Errorf("deletion error, product index %d out of range", index)
	}
	if index == products.Count()-1 {
		*products.productsList = (*products.productsList)[:index]
	} else {
		*products.productsList = append((*products.productsList)[:index], (*products.productsList)[index+1:]...)
	}
	return nil

}

func (products *Products) GetNextProductId() int {
	if products.Count() == 0 {
		return 0
	}
	lastProduct := (*products.productsList)[products.Count()-1]
	return lastProduct.GetID() + 1
}

func (products *Products) GetProductIndexById(id int) int {
	for index, product := range *products.productsList {
		if product.GetID() == id {
			return index
		}

	}
	return -1
}

func (products *Products) ToJson(w io.Writer) error {
	return toJson(products.productsList, w)
}

var productsInstance = NewProducts()

func ProductsInstance() ProductsInterface {
	return productsInstance
}

func InitProducts(products ProductsInterface) {
	productsInstance = products
}
