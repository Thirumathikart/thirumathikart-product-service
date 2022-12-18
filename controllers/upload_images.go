package controllers

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/thirumathikart/thirumathikart-product-service/models"
	"gorm.io/gorm"
)

func UploadProductImage(files []*multipart.FileHeader, productID uint, db *gorm.DB) error {

	// Get Product ID and validate with seller ID

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		filePath := filepath.Join("product_images", strconv.FormatUint(uint64(productID), 10)+".jpg")
		dst, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		productImage := models.ProductImage{
			ProductID: productID,
			ImageURL:  filePath,
		}
		db.Create(&productImage)
	}

	// Create Product with Product ID and product images
	return nil
}
