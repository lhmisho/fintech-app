package routes

import (
	"fintech-app/pkg/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

var RegisterRouters = func(router *mux.Router) {
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginController).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.GetUserController).Methods("GET")
	router.HandleFunc("/transaction", controllers.TransactionController).Methods("POST")
	router.HandleFunc("/transactions/{userID}", controllers.GetMyTransactionController).Methods("GET")
}

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	RegisterRouters(r)
	http.Handle("/", r)

	return r
}
