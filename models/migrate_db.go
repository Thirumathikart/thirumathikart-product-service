package models

import (
	"github.com/thirumathikart/thirumathikart-product-service/config"
)

func MigrateDB() {
	db := config.GetDB()

	for _, model := range []interface{}{
		// Add models here
	} {
		if err := db.AutoMigrate(&model); err != nil {
			panic(err)
		}
	}
}
