package util

import (
	"github.com/absentbird/TESC-Farm/internal/harvest"
	"github.com/absentbird/TESC-Farm/internal/labor"
	"github.com/absentbird/TESC-Farm/internal/sales"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// Create a shorthand function to check for errors
func Check(e error, m string) {
	if e != nil {
		log.Panic(m+": ", e.Error())
	}
}

func ConnectDB(dbconn string) {
	database, err := gorm.Open(sqlite.Open(dbconn), &gorm.Config{})
	Check(err, "Database connection error")

	err = database.AutoMigrate(&labor.Hours{})
	Check(err, "Hours migration error")

	err = database.AutoMigrate(&labor.Worker{})
	Check(err, "Worker migration error")

	err = database.AutoMigrate(&harvest.Crop{})
	Check(err, "Crop migration error")

	err = database.AutoMigrate(&harvest.Harvest{})
	Check(err, "Harvest migration error")

	err = database.AutoMigrate(&harvest.Process{})
	Check(err, "Process migration error")

	err = database.AutoMigrate(&sales.Product{})
	Check(err, "Product migration error")

	err = database.AutoMigrate(&sales.Sale{})
	Check(err, "Sales migration error")

	// Set global database variable
    DB = database
}
