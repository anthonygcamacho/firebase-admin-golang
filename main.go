package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func main() {
	// Setup --------------------------------------------------------------------------------------------
	ctx := context.Background()

	opt := option.WithCredentialsFile("service_account_sdk.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
		// return nil, fmt.Errorf("error initializing app: %v", err)
	}

	// Firestore --------------------------------------------------------------------------------------------

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting firebase client: %v\n", err)
	}

	// Get Doc
	result, err := client.Collection("sampleData").Doc("inspiration").Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Print(result.Data())

	// defer client.Close()

	// Users --------------------------------------------------------------------------------------------

	uid := "tpmussGri1MjvgSOcY6nSO3TAei1"

	clientAuth, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	// Get User By ID
	userByID, err := clientAuth.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %T\n", userByID)

	// Get User By Email
	email := "user2@example.com"
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

	// Create User
	createUserParams := (&auth.UserToCreate{}).
		Email("user2@example.com").
		EmailVerified(false).
		PhoneNumber("+15555550199").
		Password("secretPassword").
		DisplayName("John Doe").
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	log.Printf("createUserParams: %T\n", createUserParams)
	// u1, err := clientAuth.CreateUser(ctx, createUserParams)
	// if err != nil {
	// 	log.Fatalf("error creating user: %v\n", err)
	// }
	// log.Printf("Successfully created user: %v\n", u1)

	// Update User
	updateUserParams := (&auth.UserToUpdate{}).
		// Email("user2@example.com").
		EmailVerified(true)
	log.Printf("updateUserParams: %T\n", updateUserParams)
	// Password("newPassword").
	// PhoneNumber("+15555550144")
	// DisplayName("John Doe").
	// PhotoURL("http://www.example.com/12345678/photo.png").
	// Disabled(true)
	// u2, err := clientAuth.UpdateUser(ctx, uid, updateUserParams)
	// if err != nil {
	// 	log.Fatalf("error updating user: %v\n", err)
	// }
	// log.Printf("Successfully updated user: %v\n", u2)

	// Delete User
	// uidToDelete := "h4aoYZdSENgCZHSf92PbW5M5ot93"
	// err2 := clientAuth.DeleteUser(ctx, uidToDelete)
	// if err != nil {
	// 	log.Fatalf("error deleting user: %v\n", err2)
	// }
	// log.Printf("Successfully deleted user: %s\n", uid)

	// defer clientAuth.Close()

}
