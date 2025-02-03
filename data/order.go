package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     int       `json:"price"`
	TotalSum  int       `json:"total_sum"`
	Status    string    `json:"status"`
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

func (o *Order) Validate() error {
	if o.UserID == 0 {
		return errors.New("missing user_id")
	}
	if o.ProductID == 0 {
		return errors.New("missing product_id")
	}
	if o.Quantity == 0 {
		return errors.New("quantity is 0")
	}

	return nil
}
