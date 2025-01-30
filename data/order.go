package data

import (
	"encoding/json"
	"io"
	"time"
)

type Order struct {
	ItemID    int       `json:"item_id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,gte=0"`
	TotalSum  int       `json:"total_sum" validate:"required,gte=0"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func (o *Order) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

func (o *Order) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(o)
}
