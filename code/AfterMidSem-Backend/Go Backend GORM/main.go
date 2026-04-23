package main

import (
	"gorm/config"
	"gorm/models"
	"gorm/routes"

	"github.com/gin-gonic/gin"
)


func main(){
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})
	r:= gin.Default()
	// setup
	routes.SetupRoutes(r)
	r.Run(":9092")

	
}