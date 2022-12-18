package config

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/joho/godotenv"
)

var DbHost string
var DbUser string
var DbPassword string
var DbName string
var DbPort string
var ServerPort string
var RPCPort string
var AuthService string
var CategoriesID map[string]int

func Environment() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(color.RedString("Error loading .env"))
	}

	DbHost = os.Getenv("DB_HOST")
	DbUser = os.Getenv("POSTGRES_USER")
	DbPassword = os.Getenv("POSTGRES_PASSWORD")
	DbName = os.Getenv("POSTGRES_DB")
	DbPort = os.Getenv("DB_PORT")
	ServerPort = os.Getenv("SERVER_PORT")
	RPCPort = os.Getenv("RPC_PORT")
	AuthService = os.Getenv("AUTH_SERVICE")

	CategoriesID = map[string]int{
		"Food":                        1,
		"Clothing":                    2,
		"Handricrafts":                3,
		"Grocery":                     4,
		"Fahion and Jewellery":        5,
		"Beauty and Healthcare":       6,
		"Office Code":                 7,
		"Organic Fruits & Vegetables": 8,
		"Others":                      9,
	}
}
