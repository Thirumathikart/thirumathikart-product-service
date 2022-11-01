package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `gorm:"default:not null;"`
	CategoryID  int    `gorm:"default:0;"`
	SellerID    int    `gorm:"default:0;"`
	Price       int    `gorm:"default:0;"`
	Description string `gorm:"default:null;"`
	Stock       int    `gorm:"default:0;"`
	ImageURL    string `gorm:"default:null;"`
	IsAvailable bool   `gorm:"default:true;"`
}
