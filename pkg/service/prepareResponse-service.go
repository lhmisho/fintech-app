package service

import (
	"fintech-app/pkg/models"
	"net/http"
)

func PrepareResponse(user *models.User, accounts []models.ResponseAccount) map[string]interface{} {
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
