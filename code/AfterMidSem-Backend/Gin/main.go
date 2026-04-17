package main

import (
	// "fmt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AuthMiddleware (ctx *gin.Context) {
	// if berar token matched
	// return
	token:= ctx.GetHeader("Authorization")
	if token != "Bearer myToken" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": "Failed",
			"message": "Wrong token",
		})
		return
	} else {
		ctx.Next()
	}
	// else
}

func main(){
	r:= gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"result": "OK"})
	})

	// /user/:name
	r.GET("/users/:name", func(ctx *gin.Context) {
		name:= ctx.Param("name")
		ctx.JSON(http.StatusOK, gin.H{"result": "OK", "user": name})
		// ctx.JSON(http.StatusOK, gin.H{"result": "OK", "user": fmt.Sprintf(name)})
		// Sprintf = return formatted string
		// Fprintf = 
		// printf vs println
	})

	// endpoint /search query params q=, page
	r.GET("/search", func(ctx *gin.Context) {
		q:= ctx.Query("q")
		page:= ctx.Query("page")

		ctx.JSON(http.StatusAccepted, gin.H{
			"result":"OK",
			"message":fmt.Sprintf("Query is %s and Page is %s", q, page),
		})
	})



	// /login Bearer Token
	// Authentication vs Authorisation
	r.Use(AuthMiddleware)

	r.POST("/login", func(ctx *gin.Context) {

		// username
		// password

		var loginRequest LoginRequest

		// if statement; condition {}
		
		if err:= ctx.ShouldBindJSON(&loginRequest); err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"result": "Failed",
				"message": "JSON Binding failed",
			})
			return
		}

		if loginRequest.Username == "Polaris" && loginRequest.Password == "password" {
			ctx.JSON(http.StatusAccepted, gin.H{
				"result": "OK",
				"message": "You are authorised",
			})
		} else {
			ctx.JSON(http.StatusNotAcceptable, gin.H{
				"result": "Failed",
				"message": "You are not authorised",
			})
		}

	})
	
	

	r.Run(":9090")
}