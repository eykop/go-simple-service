package data

import (
	"fmt"
	"io"
	"time"
)

type ProductInterface interface {
	ToJson(w io.Writer) error
	FromJson(r io.Reader) error
	GetID() int
	GetName() string
	SetID(id int)
	UpdateProduct(updated ProductInterface)
	Validate() error
}

type Products []ProductInterface

func ProductsCount() int {
	return len(*productsList)
}

func GetProductByIndex(index int) ProductInterface {
	if index < 0 || index > len(*productsList)-1 {
		//fmt.Errorf("update error, product index %d out of range", index)
		return nil
	}
	return (*productsList)[index]

}
func GetProductsList() *Products {
	return productsList
}

func AddPorduct(p ProductInterface) {
	p.SetID(getNextProductId())
	*productsList = append(*productsList, p)
}

func UpdateProduct(incommingProd ProductInterface, index int) error {
	//add logger
	if index < 0 || index > len(*productsList)-1 {

		return fmt.Errorf("update error, product index %d out of range", index)
	}
	p := (*productsList)[index]
	p.UpdateProduct(incommingProd)
	return nil
}

func DeleteProduct(index int) error {
	// todo add logger l.Info("Delete Procut", zap.Int("id", productId))

	if index < 0 || index > len(*productsList)-1 {
		return fmt.Errorf("deletion error, product index %d out of range", index)
	}
	if index == len(*productsList) {
		*productsList = (*productsList)[:index]
	} else {
		*productsList = append((*productsList)[:index], (*productsList)[index+1:]...)
	}
	return nil

}

func getNextProductId() int {
	if len(*productsList) == 0 {
		return 0
	}
	lastProduct := (*productsList)[len(*productsList)-1]
	return lastProduct.GetID() + 1
}

func GetProductIndexById(id int) int {
	for index, product := range *productsList {
		if product.GetID() == id {
			return index
		}

	}
	return -1
}

func (pl *Products) ToJson(w io.Writer) error {
	return toJson(pl, w)
}

var productsList = &Products{
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
		ID:          1,
		Name:        "Latte",
		Description: "Lite coffe drink with milk...",
		Price:       2.49,
		SKU:         "a345d9d6-0c08-45a2-887a-4c22594737b3",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
