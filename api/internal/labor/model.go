package labor

import (
	"gorm.io/gorm"
    "time"
)

type Worker struct {
	gorm.Model
    Name     string
    Barcode  string
    Position string
}

type Hours struct {
	gorm.Model
    Start      time.Time
    Duration   time.Duration
    Department string
    Task       string
    Worker     Worker
    WorkerID   uint
}
