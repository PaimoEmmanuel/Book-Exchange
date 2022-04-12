package main

import (
	"Book-Exchange/database"
	"Book-Exchange/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if database.Err != nil {
		fmt.Errorf("error initializing app: %v", database.Err)
	}
	http.HandleFunc("/", routes.WelcomeHandler)
	http.HandleFunc("/signup", routes.SignupHandler)
	http.HandleFunc("/signin", routes.SigninHandler)
	serverError := http.ListenAndServe(":8080", nil)
	if serverError != nil {
		log.Fatal("serverError", serverError)
	}
}
