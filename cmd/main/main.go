package main

import (
	"fintech-app/pkg/routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRouters(r)
	http.Handle("/", r)
	fmt.Println("Server is running at: http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
