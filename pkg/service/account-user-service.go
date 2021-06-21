package service

import (
	"fintech-app/pkg/config"
	"fintech-app/pkg/models"
)

func init() {
	config.Connect()
	db = config.GetDB()
}

func updateAccount(id uint, amount int){
	db.Model(&models.Account{}).Where("id=?", id).Update("balance", amount)
}
