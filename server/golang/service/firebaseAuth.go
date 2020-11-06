package service

import (
	"context"
	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"
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

//getAccessToken get access Token from context
func getHeader(ctx context.Context, key string) (token string, ok bool) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	token = headers.Get(key)[0]
	return
}

//VerifyTokenAndGetUUID check token and get token's user id
func VerifyTokenAndGetUUID(ctx context.Context) (uuid string, ok bool) {
	token, ok := getHeader(ctx, "X-Chat-Access")
	if !ok {
		return "", false
	}
	jwt, err := auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		return "", false
	}
	uuid, ok = jwt.Claims["chat-uuid"].(string)
	return
}
