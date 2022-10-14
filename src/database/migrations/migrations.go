package migrations

import (
	"github.com/lucca-rodrigues/myStock-API-GO/src/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Product{})
}