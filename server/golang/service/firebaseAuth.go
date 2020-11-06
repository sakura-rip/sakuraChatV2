package service

import (
	"context"
	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
	"fmt"
	"google.golang.org/api/option"
	"os"
)

var auth *firebaseAuth.Client

// init firebase authentication
func init() {
	ctx := context.Background()
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SECRET"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(fmt.Errorf("error on init firebase auth: %v", err))
	}
	auth, err = app.Auth(ctx)
	if err != nil {
		panic(fmt.Errorf("error on get auth %v", err))
	}
}
