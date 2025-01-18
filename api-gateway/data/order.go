package data

import (
	"encoding/json"
	"io"
)

type Order struct {
	ItemID   int    `json:"item_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,gte=0"`
	Price    int    `json:"price" validate:"required,gte=0"`
}

func (o *Order) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Order) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}
