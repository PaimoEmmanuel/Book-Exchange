package database

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitializeFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("bookexchange-23c7b-firebase-adminsdk-bzuhp-8cfea87e00.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	return app, err
}

var App, Err = InitializeFirebase()
