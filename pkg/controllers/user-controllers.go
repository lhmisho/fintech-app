package controllers

import (
	"fintech-app/pkg/models"
	"fintech-app/pkg/service"
	"fintech-app/pkg/utils"
	"github.com/gorilla/mux"
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

		var response = service.PrepareResponse(user, accounts, true)
		service.PrepareCreateResponse(response, w)

	}else {
		service.PrepareErrResponse(valid, w)
	}

}

func GetUserController(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	userId := params["id"]
	auth := r.Header.Get("Authorization")
	user := service.GetUser(userId, auth)
	if user["Status"] == 200{
		service.PrepareSuccessResponse(user, w)
	}else {
		service.PrepareErrResponse(user, w)
	}
}