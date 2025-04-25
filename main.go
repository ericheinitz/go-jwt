package main

import (
	"github.com/ericheinitz/go-jwt/controllers"
	"github.com/ericheinitz/go-jwt/initializers"
	"github.com/ericheinitz/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.POST("/signup", controllers.Singup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run()
}
