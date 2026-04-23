package routes

import (
	"gorm/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine){

	r.GET("/health", controller.HealthCheck)

	users:= r.Group("/api")
	users.POST("/users", controller.CreateUser)
	users.GET("/users", controller.GetUsers)
}