package sales

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Type     string `json:"type"`
	Unit     string `json:"unit"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Weight   int    `json:"weight"`
	Length   int    `json:"length"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

type Sale struct {
	gorm.Model
	Type      string   `json:"type,omitempty"`
	Quantity  int      `json:"quantity"`
	Price     int      `json:"price"`
	Tax       int      `json:"tax,omitempty"`
	Total     int      `json:"total"`
	Product   *Product `json:"product,omitempty"`
	ProductID uint     `json:"product_id"`
}
