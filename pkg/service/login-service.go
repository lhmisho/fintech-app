package service

import (
	"fintech-app/pkg/config"
	"fintech-app/pkg/interfaces"
	"fintech-app/pkg/models"
	"fintech-app/pkg/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}



func LoginService(username string, pass string) map[string]interface{} {
	valid := utils.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})

	if valid{
		user := &models.User{}
		if db.Where("username=?", username).First(&user).RecordNotFound() {
			return map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Record Not found",
			}
		}

		// Verify password
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "User or Password mismatched!",
			}
		}

		// find account for user
		accounts := []models.ResponseAccount{}
		db.Table("accounts").Select("id, name, balance").Where("user_id=?", user.ID).Scan(&accounts)

		var response = PrepareResponse(user, accounts)
		return response
	}else{
		return map[string]interface{}{
			"status": http.StatusBadRequest,
			"message": "Please provide valid value",
		}
	}
}
