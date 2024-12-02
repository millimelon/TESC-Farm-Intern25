package util

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// Create a shorthand function to check for errors
func Check(e error, m string) {
	if e != nil {
		log.Panic(m+": ", e.Error())
	}
}

func ConnectDB(dbconn string) *gorm.DB {
	database, err := gorm.Open(sqlite.Open(dbconn), &gorm.Config{})
	Check(err, "Database connection error")

	err = database.AutoMigrate(&labor.Hours{})
	Check(err, "Hours migration error")

	err = database.AutoMigrate(&labor.Worker{})
	Check(err, "Worker migration error")

	err = database.AutoMigrate(&harvest.Product{})
	Check(err, "Product migration error")

	err = database.AutoMigrate(&harvest.Harvest{})
	Check(err, "Harvest migration error")

	err = database.AutoMigrate(&sales.Sale{})
	Check(err, "Sales migration error")

	// Set global database variable
	return database
}
