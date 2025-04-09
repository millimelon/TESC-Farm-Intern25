package util

import (
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var WorkerToken string
var WorkerHash string

// Create a shorthand function to check for errors
func Check(e error, m string) {
	if e != nil {
		log.Panic(m+": ", e.Error())
	}
}
