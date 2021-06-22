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

type ResponseTransaction struct {
	ID uint
	From uint
	To uint
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

func GetTransactionByAccount(id uint) []ResponseTransaction {
	transactions := []ResponseTransaction{}
	db.Table("transactions").Select("id, transactions.from, transactions.to, amount").Where(Transaction{From: id}).Or(Transaction{To: id}).Scan(&transactions)
	return transactions
}