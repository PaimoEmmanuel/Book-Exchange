package routes

import (
	errorhandler "Book-Exchange/error-handler"
	"net/http"
)

func WelcomeHandler(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	if method != "GET" {
		errorhandler.NotFoundError(writer)
		return
	}
	writer.Write([]byte("Welcome to book exchange!"))
}
