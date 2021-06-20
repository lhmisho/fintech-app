package utils

import (
	"encoding/json"
	"fintech-app/pkg/interfaces"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)
	return string(hashed)
}

func Validation(values []interfaces.Validation) bool{
	username := regexp.MustCompile(`^([A-Za-z0-9]{5,})+$`)
	email := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z0-9]+$`)
	for i := 0; i<len(values); i++{
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value){
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value){
				return false
			}
		case "password":
			if len(values[i].Value) < 5{
				return false
			}
		}
	}
	return true
}