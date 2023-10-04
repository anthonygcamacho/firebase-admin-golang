package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("project-sandbox-10e45-firebase-adminsdk-1er26-8a5de9c68e.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
		// return nil, fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Firestore(context.Background())

	result, err := client.Collection("sampleData").Doc("inspiration").Get(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	log.Print(result.Data())

	defer client.Close()

}
