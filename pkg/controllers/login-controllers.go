package controllers

import (
	"encoding/json"
	"fintech-app/pkg/service"
	"fintech-app/pkg/utils"
	"io/ioutil"
	"net/http"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Status  int
	Message string
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	//utils.ParseBody(r, &User{})
	// read body
	body, err := ioutil.ReadAll(r.Body)
	utils.HandleErr(err)

	// handle login
	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	utils.HandleErr(err)
	login := service.LoginService(formattedBody.Username, formattedBody.Password)

	if login["status"] == 200 {
		resp := login
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Status: http.StatusBadRequest, Message: "Failed to login!"}
		json.NewEncoder(w).Encode(resp)
	}

}
