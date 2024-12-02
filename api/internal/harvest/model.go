package harvest

import (
	"github.com/absentbird/TESC-Farm/internal/sales"
	"gorm.io/gorm"
    "time"
)

type Crop struct {
    gorm.Model
    Name   string
    Season string
    Type   string
}

type Harvest struct {
	gorm.Model
    Weight   float64
    Bin      string
    Start    time.Time
    Duration time.Duration
    Crop     Crop
    CropID   uint
}

type Process struct {
	gorm.Model
    Unit       string
    Quantity   int
    Weight     float64
    Cull       float64
    StudentUse float64
    ValueAdded float64
    Harvest    Harvest
    HarvestID  uint
    Product    sales.Product
    ProductID  uint
}
