package routes

import (
	"github.com/ericheinitz/go-jwt/controllers"
	"github.com/ericheinitz/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Singup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/logout", middleware.RequireAuth, controllers.Logout)
}
