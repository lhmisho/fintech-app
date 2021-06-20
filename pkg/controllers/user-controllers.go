package controllers

import (
	"fintech-app/pkg/models"
	"fintech-app/pkg/service"
	"fintech-app/pkg/utils"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	CreateAccount := &models.Account{}

	utils.ParseBody(r, CreateUser)
	// validate request
	valid := service.RegisterService(CreateUser.Username, CreateUser.Email, CreateUser.Password)
	if valid["Status"] == 200{
		CreateUser.Password = utils.HashAndSalt([]byte(CreateUser.Password))
		user := CreateUser.CreateUser()

		// create account
		CreateAccount.Name = user.Username + "'s" + " account"
		CreateAccount.Type = "S A/C"
		CreateAccount.Balance = 0
		CreateAccount.UserID = user.ID
		account := CreateAccount.CreateAccount()

		accounts := []models.ResponseAccount{}
		respAccount := models.ResponseAccount{ID: account.UserID, Name: account.Name, Balance: int(account.Balance)}
		accounts = append(accounts, respAccount)

		var response = service.PrepareResponse(user, accounts)
		service.PrepareCreateResponse(response, w)

	}else {
		service.PrepareErrResponse(valid, w)
	}

}
