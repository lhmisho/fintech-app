package models

import (
	"fintech-app/pkg/config"
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Account{})
}
