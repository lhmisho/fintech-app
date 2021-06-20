package controllers

import (
	"encoding/json"
	"fintech-app/pkg/models"
	"fintech-app/pkg/utils"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	CreateUser.Password = utils.HashAndSalt([]byte(CreateUser.Password))
	user := CreateUser.CreateUser()
	res, _ := json.Marshal(user)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
