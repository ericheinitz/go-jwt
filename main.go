package main

import (
	"os"

	"github.com/ericheinitz/go-jwt/initializers"
	"github.com/ericheinitz/go-jwt/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	initializers.SetupCors(router)
	routes.AuthRoutes(router)

	router.Run(":" + os.Getenv("PORT"))
}
