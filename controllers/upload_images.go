package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-product-service/models"
	"gorm.io/gorm"
)

func UploadProductImage(c echo.Context, productID uint, db *gorm.DB) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	hash := sha256.New()
	path := "product_images/" + hex.EncodeToString(hash.Sum(nil)) + ".jpg"
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	productImage := models.ProductImage{
		ProductID: productID,
		ImageURL:  path,
	}
	db.Create(&productImage)
	return nil
}

func UploadProductImages(files []*multipart.FileHeader, productID uint, db *gorm.DB) error {

	// Get Product ID and validate with seller ID

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		hash := sha256.New()
		if _, err := io.Copy(hash, src); err != nil {
			log.Println("1")
			log.Println(err)
			return err
		}
		if _, err := io.Copy(hash, src); err != nil {
			log.Println("2")
			log.Println(err)
			return err
		}
		srcCopy, err := file.Open()
		if err != nil {
			log.Println("3")
			log.Println(err)
			return err
		}
		defer srcCopy.Close()
		filePath := filepath.Join("product_images", hex.EncodeToString(hash.Sum(nil))+".jpg")
		log.Println(filePath)
		dst, err := os.Create(filePath)
		if err != nil {
			log.Println("4")
			log.Println(err)
			return err
		}
		defer dst.Close()
		if _, err = io.Copy(dst, srcCopy); err != nil {
			log.Println("5")
			log.Println(err)
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
