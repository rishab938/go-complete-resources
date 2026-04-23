package controller

import (
	"gorm/config"
	"gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck (ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{"result": "Server is up and running"})
}

func GetUsers(ctx *gin.Context){
	//
	var users []models.User
	result:= config.DB.Find(&users)
	if result.Error!=nil{
		ctx.JSON(http.StatusAccepted, gin.H{"result": "Couldn't find"})
		return
	}
	ctx.JSON(http.StatusAccepted, users)
}


// handler func api("/user") POST
func CreateUser(ctx *gin.Context){

	var user models.User

	// ctx.ShouldBindJSON = binding payload with User model
	if err:= ctx.ShouldBindJSON(&user); err!=nil{
		// binding failed

		// ctx.JSON = response (aka MERN res.send())
		ctx.JSON(http.StatusNotAcceptable, gin.H{"result": "JSON Binding Failed"})
		return
	}

	// user payload
	// json -> struct
	result:= config.DB.Create(&user)

	if result.Error != nil{
		ctx.JSON(http.StatusNotAcceptable, gin.H{"result": "User Creation Failed"})
		return
	}

	// db insert

	// 
	ctx.JSON(http.StatusAccepted, gin.H{"result": "User Created"})
}
