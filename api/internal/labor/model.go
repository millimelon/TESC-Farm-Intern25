package labor

import (
	"github.com/absentbird/TESC-Farm/internal/harvest"
	"gorm.io/gorm"
	"time"
)

type Worker struct {
	gorm.Model
	Name     string `json:"name"`
	Barcode  string `json:"barcode"`
	Position string `json:"position,omitempty"`
}

type Hours struct {
	gorm.Model
	Start      time.Time        `json:"start"`
	Duration   float64          `json:"duration"`
	Department string           `json:"department,omitempty"`
	Task       string           `json:"task,omitempty"`
	Worker     *Worker          `json:"worker,omitempty"`
	WorkerID   uint             `json:"worker_id"`
	Harvest    *harvest.Harvest `json:"harvest,omitempty"`
	HarvestID  uint             `json:"harvest_id,omitempty"`
	Process    *harvest.Process `json:"process,omitempty"`
	ProcessID  uint             `json:"process_id,omitempty"`
}
