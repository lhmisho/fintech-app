package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	)

var (
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=fintech_app password=postgres")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB  {
	return db
}
