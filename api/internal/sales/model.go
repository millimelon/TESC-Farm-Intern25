package sales

import (
	"gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Name     string
    Type     string
    Unit     string
    Price    int
    Quantity int
    Weight   int
    Length   int
    Width    int
    Height   int
}

type Sale struct {
    gorm.Model
    Type      string
    Quantity  int
    Price     int
    Tax       int
    Total     int
    Product   Product
    ProductID uint
}
