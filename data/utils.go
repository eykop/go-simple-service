package data

import (
	"encoding/json"
	"io"
)

func ToJson(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

func FromJson(i interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(i)
}

func getNextProductId() int {
	lastProduct := productsList[len(productsList)-1]
	return lastProduct.ID + 1
}

func GetProductIndexById(id int) int {
	for index, product := range productsList {
		if product.ID == id {
			return index
		}
	}
	return -1
}
