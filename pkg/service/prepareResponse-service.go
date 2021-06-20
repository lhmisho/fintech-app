package service

import (
	"encoding/json"
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
		"Status":  http.StatusOK,
		"Message": "Successfully logged in",
	}
	response["jwt"] = token
	response["data"] = responseUser
	return response
}

func PrepareSuccessResponse(respData map[string]interface{}, w http.ResponseWriter){
	resp := respData
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func PrepareCreateResponse(respData map[string]interface{}, w http.ResponseWriter){
	resp := respData
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func PrepareErrResponse(errResp map[string]interface{}, w http.ResponseWriter)  {
	resp := errResp
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(resp)
}
