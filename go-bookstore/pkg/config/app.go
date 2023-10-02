package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	gormDb, err := gorm.Open("mysql", "root:MySQL#25@/dineshdb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	fmt.Println("connected")

	db = gormDb
}

func GetDB() *gorm.DB {
	return db
}
