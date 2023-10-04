package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"

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
	log.Printf("Successfully fetched user data: %+v\n", userByEmail)

	defer client.Close()

}
