package service

import (
	"fintech-app/pkg/models"
	"fintech-app/pkg/utils"
)

func GetMyTransactions(id string, jwt string) map[string]interface{} {
	isValid := utils.ValidateToken(id, jwt)

	if !isValid{
		return map[string]interface{}{"Status": 400, "Message": "Invalid Token!"}
	}
	accounts := []models.ResponseAccount{}
	db.Table("accounts").Select("id, name, balance").Where("user_id=?", id).Scan(&accounts)

	transactions := []models.ResponseTransaction{}
	for i := 0; i < len(accounts); i++{
		accTransactions := models.GetTransactionByAccount(accounts[i].ID)
		transactions = append(transactions, accTransactions...)
	}
	var response = map[string]interface{}{"Status": 200, "Message": "Data successfully retrieved!"}
	response["data"] = transactions
	return response
}
