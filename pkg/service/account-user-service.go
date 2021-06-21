package service

import (
	"fintech-app/pkg/config"
	"fintech-app/pkg/models"
	"fintech-app/pkg/utils"
	"strconv"
)

func init() {
	config.Connect()
	db = config.GetDB()
}

// UpdateAccount Create function update account
func updateAccount(id uint, amount int) *models.ResponseAccount{
	account := &models.Account{}
	responseAcc := &models.ResponseAccount{}

	db.Where("id=?", id).First(&account)
	account.Balance = uint(amount)
	db.Save(&account)

	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = int(account.Balance)
	return responseAcc
}

func getAccount(id uint) *models.Account {
	account := &models.Account{}
	 if db.Where("id=?", id).First(&account).RecordNotFound(){
	 	return nil
	 }
	 return account
}

func TransactionService(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{}{
	userIdString := strconv.FormatUint(uint64(userId), 10)
	isValid := utils.ValidateToken(userIdString, jwt)

	// if invalid token than return
	if !isValid{
		return map[string]interface{}{"Status":400, "Message": "Invalid token"}
	}

	fromAcc := getAccount(from)
	toAcc := getAccount(to)

	if fromAcc == nil || toAcc == nil {
		return map[string]interface{}{"Status": 400, "Message": "Account not found!"}
	}else if fromAcc.UserID != userId{
		return map[string]interface{}{"Status": 400, "Message": "You are not owner of the account!"}
	}else if int(fromAcc.Balance) < amount {
		return map[string]interface{}{"Status": 400, "Message": "Insufficient balance!"}
	}

	updatedAccount := updateAccount(from, int(fromAcc.Balance) - amount)
	updateAccount(to, int(fromAcc.Balance) + amount)

	models.CreateTransaction(from, to, amount)
	var response = map[string]interface{}{"Status": 200, "Message": "Transaction success"}
	response["data"] = updatedAccount
	return response

}