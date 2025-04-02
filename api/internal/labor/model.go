package labor

import (
	"github.com/absentbird/TESC-Farm/internal/harvest"
	"gorm.io/gorm"
	"time"
)

type Worker struct {
	gorm.Model
	Barcode string `json:"barcode"`
}

type Hours struct {
	gorm.Model
	Start      time.Time `json:"start"`
	Duration   float64   `json:"duration"`
	Department string    `json:"department,omitempty"`
	Task       *Task     `json:"task,omitempty"`
	TaskID     uint      `json:"task_id"`
	Worker     *Worker   `json:"worker,omitempty"`
	WorkerID   uint      `json:"worker_id"`
}

type Task struct {
	gorm.Model
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Barcode      string              `json:"barcode"`
	Type         string              `json:"type"`
	Area         *harvest.Area       `json:"area,omitempty"`
	AreaID       uint                `json:"area_id"`
	Preharvest   *harvest.Preharvest `json:"preharvest,omitempty"`
	PreharvestID uint                `json:"preharvest_id,omitempty"`
	Harvest      *harvest.Harvest    `json:"harvest,omitempty"`
	HarvestID    uint                `json:"harvest_id,omitempty"`
	Process      *harvest.Process    `json:"process,omitempty"`
	ProcessID    uint                `json:"process_id,omitempty"`
}
