package routes

import (
	"Book-Exchange/database"
	errorhandler "Book-Exchange/error-handler"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type UserSignin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SigninHandler(writer http.ResponseWriter, request *http.Request) {
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
	var user UserSignin
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	_, err = json.Marshal(user)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	writer.Header().Set("content-type", "application/json")

	// Get an auth client from the firebase.App
	client, err := database.App.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	u, err := client.GetUserByEmail(context.Background(), user.Email)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", user.Email, err)
	}
	log.Printf("Successfully fetched user data: %v\n", u.DisplayName)

	writer.Write([]byte("success"))
}
