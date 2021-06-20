package controllers

import (
	"fintech-app/pkg/models"
	"fintech-app/pkg/service"
	"fintech-app/pkg/utils"
	"net/http"
)

type ErrResponse struct {
	Status  int
	Message string
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	// read body
	user := &models.User{}
	utils.ParseBody(r, user)
	login := service.LoginService(user.Username, user.Password)

	if login["Status"] == 200 {
		service.PrepareSuccessResponse(login, w)
	} else {
		resp := map[string]interface{}{"Status": http.StatusBadRequest, "Message": "User or password mismatch"}
		service.PrepareErrResponse(resp, w)
	}

}
