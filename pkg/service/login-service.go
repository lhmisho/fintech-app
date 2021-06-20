package service

import (
	"fintech-app/pkg/config"
	"fintech-app/pkg/models"
	"fintech-app/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
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

	// setup response
	responseUser := models.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	// sign token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	utils.HandleErr(err)

	// prepare response
	var response = map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Successfully logged in",
	}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}
