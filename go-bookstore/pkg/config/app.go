package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	databaseName := os.Getenv("DATABASE_NAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	gormDb, err := gorm.Open("mysql", "root:"+databasePassword+"@/"+databaseName+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	fmt.Println("connected")

	db = gormDb
}

func GetDB() *gorm.DB {
	return db
}
