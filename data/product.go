package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
	"unicode"
)

type Product struct {
	ID        int       `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Category  string    `json:"category"`
	UOM       string    `json:"uom"` // Unit of Measurement
	Brand     string    `json:"brand"`
	Stock     int       `json:"stock"` // Количество на складе
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Product) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) Validate() error {

	// if len(p.SKU) != 10 {
	// 	return errors.New("SKU must consist 10 digits")
	// }

	if p.Price <= 0 {
		return errors.New("incorrect price")
	}

	if len(p.Category) == 0 || len(p.Category) > 50 {
		return errors.New("category too long or absent")
	}

	for _, char := range p.Category {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return errors.New("category name must contain only letters and numbers")
		}
	}

	if len(p.UOM) == 0 {
		return errors.New("UOM must contain data")
	}

	if len(p.Brand) == 0 || len(p.Brand) > 50 {
		return errors.New("brand too long or absent")
	}

	for _, char := range p.Brand {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return errors.New("brand name must contain only letters and numbers")
		}
	}

	if p.Stock < 0 {
		return errors.New("negative stock")
	}

	return nil
}
