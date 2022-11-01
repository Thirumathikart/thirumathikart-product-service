package models

import (
	"github.com/thirumathikart/thirumathikart-product-service/config"
)

func MigrateDB() {
	db := config.GetDB()

	for _, model := range []interface{}{
		&Product{},
	} {
		if err := db.AutoMigrate(&model); err != nil {
			panic(err)
		}
	}
}
