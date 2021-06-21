package models

import (
	"fintech-app/pkg/config"
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	To uint
	From uint
	Amount int
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Transaction{})
}

func CreateTransaction(From uint, To uint, Amount int)  {
	transaction := &Transaction{To: To, From: From, Amount: Amount}
	db.Create(&transaction)
}