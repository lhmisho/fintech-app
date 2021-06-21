package controllers

import (
	"fintech-app/pkg/service"
	"fintech-app/pkg/utils"
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