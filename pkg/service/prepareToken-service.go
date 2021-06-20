package service

import (
	"fintech-app/pkg/models"
	"fintech-app/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func prepareToken(user *models.User) string {
	// sign token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	utils.HandleErr(err)

	return token
}
