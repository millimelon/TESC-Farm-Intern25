package harvest

import (
	"github.com/absentbird/TESC-Farm/internal/sales"
	"gorm.io/gorm"
)

type Crop struct {
	gorm.Model
    Name   string `json:"name"`
	Season string `json:"season"`
	Type   string `json:"type"`
}

type Harvest struct {
	gorm.Model
    Weight   float64       `json:"weight"`
    Bin      string        `json:"bin"`
    Crop     Crop          `json:"crop"`
    CropID   uint          `json:"crop_id"`
}

type Process struct {
	gorm.Model
    Unit       string        `json:"unit"`
    Quantity   int           `json:"qty"`
    Weight     float64       `json:"weight"`
    Cull       float64       `json:"cull"`
    StudentUse float64       `json:"student_use"`
    ValueAdded float64       `json:"value_added"`
    Harvest    Harvest       `json:"harvest"`
    HarvestID  uint          `json:"harvest_id"`
    Product    sales.Product `json:"product"`
    ProductID  uint          `json:"product_id"`
}
