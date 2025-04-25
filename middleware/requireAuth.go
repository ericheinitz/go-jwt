package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ericheinitz/go-jwt/initializers"
	"github.com/ericheinitz/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		fmt.Println("Error getting cookie:", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "No authorization cookie found",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode the JWT string

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		fmt.Println("Error parsing token:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find the user with the token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Attach the user to the context
		c.Set("user", user)

		// Continue to the next middleware
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
