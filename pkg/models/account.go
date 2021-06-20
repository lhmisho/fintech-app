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

func (account *Account) CreateAccount() *Account {
	db.NewRecord(account)
	db.Create(&account)
	return account
}

type ResponseAccount struct {
	ID      uint
	Name    string
	Balance int
}

