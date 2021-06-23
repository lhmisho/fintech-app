package service

import (
	"fintech-app/pkg/models"
	"fintech-app/pkg/utils"
	"net/http"
)

//func init() {
//	config.Connect()
//	db = config.GetDB()
//}

func GetUser(id string, jwt string) map[string]interface{} {
	isValid := utils.ValidateToken(id, jwt)

	if isValid {
		user := &models.User{}
		if db.Where("id=?", id).First(&user).RecordNotFound() {
			return map[string]interface{}{
				"Status":  http.StatusBadRequest,
				"Message": "User Not found",
			}
		}

		// find account for user
		accounts := []models.ResponseAccount{}
		db.Table("accounts").Select("id, name, balance").Where("user_id=?", user.ID).Scan(&accounts)

		var response = PrepareResponse(user, accounts, false)
		return response
	} else {
		return map[string]interface{}{"Status": http.StatusBadRequest, "Message": "Invalid token!"}
	}
}
