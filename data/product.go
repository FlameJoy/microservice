package data

import "time"

type Product struct {
	ID        int       `json:"id"`
	SKU       int       `json:"sku"`
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
