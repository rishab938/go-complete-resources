go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm 
go get -tool github.com/air-verse/air@latest
go get -u gorm.io/driver/sqlite


```package config

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
```


```
package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HealthCheck (ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{"result": "Server is up and running"})
}

func getUsers(){
	//

}


func CreateUser(){
	//

}

```

```
package models

type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
}
```

```
package routes

import (
	"gorm/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine){

	r.GET("/health", controller.HealthCheck)
}
```

```
package main

import (
	"gorm/config"
	"gorm/routes"

	"github.com/gin-gonic/gin"
)


func main(){
	config.ConnectDB()
	r:= gin.Default()
	// setup
	routes.SetupRoutes(r)
	r.Run(":9092")
}
```


```
```

```
```
