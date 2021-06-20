package service

import (
	"fintech-app/pkg/interfaces"
	"fintech-app/pkg/utils"
	"net/http"
)

func RegisterService(username string, email string, pass string) map[string]interface{}{
	valid := utils.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})
	if valid{
		return map[string]interface{}{
			"Status": http.StatusOK,
			"Message": "All is ok",
		}
	}else {
		return map[string]interface{}{
			"Status": http.StatusBadRequest,
			"Message": "Please provide valid value",
		}
	}
}
