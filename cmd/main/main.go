package main

import (
	"fintech-app/pkg/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.InitRouter()
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
