package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB(){
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err!=nil{
		// error
		log.Fatal("Failed to connect DB", err)
	}
	log.Println("DB connection successful")
	// log print DB connected successful
}