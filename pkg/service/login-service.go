package service

import (
	"fintech-app/pkg/config"
	"fintech-app/pkg/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func prepareResponse(user *models.User, accounts []models.ResponseAccount) map[string]interface{} {
	// setup response
	responseUser := models.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}
	// prepare response
	var token = prepareToken(user)
	var response = map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Successfully logged in",
	}
	response["jwt"] = token
	response["data"] = responseUser
	return response
}

func LoginService(username string, pass string) map[string]interface{} {
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

	var response = prepareResponse(user, accounts)
	return response
}
