package routes

import (
	"Book-Exchange/database"
	errorhandler "Book-Exchange/error-handler"
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
	"io/ioutil"
	"log"
	"net/http"
)

type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignupHandler(writer http.ResponseWriter, request *http.Request) {
	request.Header.Set("Accept", "application/json")
	method := request.Method
	if method != "POST" {
		errorhandler.NotFoundError(writer)
		return
	}

	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// Unmarshal
	var user Person
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	writer.Header().Set("content-type", "application/json")

	params := (&auth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password).
		DisplayName(user.Name)
	client, err := database.App.Auth(context.Background())
	u, err := client.CreateUser(context.Background(), params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", &u)
	writer.Write(output)
}
