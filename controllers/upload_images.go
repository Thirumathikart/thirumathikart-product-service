package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

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
		hash := sha256.New()
		if _, err := io.Copy(hash, src); err != nil {
			return err
		}
		if _, err := io.Copy(hash, src); err != nil {
			return err
		}
		srcCopy, err := file.Open()
		if err != nil {
			return err
		}
		defer srcCopy.Close()
		filePath := filepath.Join("product_images", hex.EncodeToString(hash.Sum(nil))+".jpg")
		dst, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer dst.Close()
		if _, err = io.Copy(dst, srcCopy); err != nil {
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
