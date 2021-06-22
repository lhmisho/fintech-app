package controllers

import (
	"fintech-app/pkg/service"
	"fintech-app/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionBody struct {
	UserId uint
	From uint
	To uint
	Amount int
}

func TransactionController(w http.ResponseWriter, r *http.Request)  {
	transactionBody := &TransactionBody{}
	utils.ParseBody(r, transactionBody)
	auth := r.Header.Get("Authorization")
	transaction := service.TransactionService(transactionBody.UserId, transactionBody.From, transactionBody.To, transactionBody.Amount, auth)
	service.PrepareSuccessResponse(transaction, w)
}

func GetMyTransactionController(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	userId := params["userID"]
	auth := r.Header.Get("Authorization")
	user := service.GetMyTransactions(userId, auth)
	if user["Status"] == 200{
		service.PrepareSuccessResponse(user, w)
	}else {
		service.PrepareErrResponse(user, w)
	}
}