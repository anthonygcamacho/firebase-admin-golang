package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("service_account_sdk.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
		// return nil, fmt.Errorf("error initializing app: %v", err)
	}

	// Firestore --------------------------------------------------------------------------------------------

	client, err := app.Firestore(ctx)

	result, err := client.Collection("sampleData").Doc("inspiration").Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Print(result.Data())

	// defer client.Close()

	// Users --------------------------------------------------------------------------------------------

	clientAuth, err := app.Auth(ctx)

	uid := "SPXO9WzHVAU5kVxQtmO3npvSXHZ2"

	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	userByID, err := clientAuth.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %T\n", userByID)

	email := "anthony.g.camacho@mail.com"
	userByEmail, err := clientAuth.GetUserByEmail(ctx, email)
	if err != nil {
		log.Fatalf("error getting user by email %s: %v\n", email, err)
	}
	// %+v : Show field names
	log.Printf("USER: %+v\n", userByEmail)
	log.Printf("UID: %+v\n", userByEmail.UID)
	log.Printf("DisplayName: %+v\n", userByEmail.DisplayName)
	log.Printf("Email: %+v\n", userByEmail.Email)
	log.Printf("PhoneNumber: %+v\n", userByEmail.PhoneNumber)
	log.Printf("PhotoURL: %+v\n", userByEmail.PhotoURL)
	log.Printf("ProviderID: %+v\n", userByEmail.ProviderID)
	log.Printf("Verified: %+v\n", userByEmail.EmailVerified)
	log.Printf("Disabled: %+v\n", userByEmail.Disabled)
	log.Printf("CreationTimestamp: %+v\n", userByEmail.UserMetadata.CreationTimestamp)
	log.Printf("LastLogInTimestamp: %+v\n", userByEmail.UserMetadata.LastLogInTimestamp)
	log.Printf("LastRefreshTimestamp: %+v\n", userByEmail.UserMetadata.LastRefreshTimestamp)
	log.Printf("UserInfo: %+v\n", userByEmail.UserInfo)
	log.Printf("ProviderUserInfo: %+v\n", userByEmail.ProviderUserInfo[0])
	log.Printf("UserMetadata: %+v\n", userByEmail.UserMetadata)

	params := (&auth.UserToCreate{}).
		Email("user@example.com").
		EmailVerified(false).
		PhoneNumber("+15555550100").
		Password("secretPassword").
		DisplayName("John Doe").
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	u, err := clientAuth.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)

	// defer clientAuth.Close()

}
