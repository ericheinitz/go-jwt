package initializers

import (
	"github.com/ericheinitz/go-jwt/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
