package routes

import (
	"fintech-app/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRouters = func(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginController).Methods("POST")
}
