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
    Position string `json:"position"`
}

type Hours struct {
	gorm.Model
	Start      time.Time     `json:"start"`
    Duration   time.Duration `json:"duration"`
    Department string        `json:"department"`
    Task       string        `json:"task"`
    Worker     Worker        `json:"worker"`
    WorkerID   uint          `json:"worker_id"`
    Harvest    harvest.Harvest `json:"harvest"`
    HarvestID  uint          `json:"harvest_id"`
    Process    harvest.Process `json:"process"`
    ProcessID  uint          `json:"process_id"`
}
