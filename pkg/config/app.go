package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=fintech_app password=postgres")
	if err != nil {
		panic(err)
	}
	//database.DB().SetMaxIdleConns(20)
	//database.DB().SetMaxOpenConns(200)
	db = database
}

func GetDB() *gorm.DB {
	return db
}
