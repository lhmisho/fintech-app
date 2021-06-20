package controllers

import (
	"encoding/json"
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
	if valid["status"] == 200{
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
		res, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(res)

	}else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&valid)
	}

}
