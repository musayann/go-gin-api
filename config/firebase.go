package config

import (
	// "context"
	"context"
	"google.golang.org/api/iterator"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type FirestoreConnection struct {
	client *firestore.Client
	ctx    context.Context
}

type FirestoreData map[string]interface{}

func ConnectFirestore() FirestoreConnection {
	// Use the application default credentials
	ctx := context.Background()

	conf := &firebase.Config{ProjectID: "go-learning-e5412"}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return FirestoreConnection{
		client, ctx,
	}
}

func (connection *FirestoreConnection) FirestoreAdd(collection string, data map[string]interface{}) {

	_, _, err := connection.client.Collection(collection).Add(connection.ctx, data)

	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}
}

func (connection *FirestoreConnection) FirestoreGetAll(collection string) []FirestoreData {
	iter := connection.client.Collection(collection).Documents(connection.ctx)
	var result []FirestoreData
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			continue
		}
		result = append(result, doc.Data())
	}
	return result
}
