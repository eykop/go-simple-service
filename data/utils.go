package data

import (
	"encoding/json"
	"io"
)

func toJson(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

func fromJson(i interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(i)
}
