package initializers

import (
	"example.com/practice/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
