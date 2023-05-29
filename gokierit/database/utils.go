package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func newConnectMongo() *mongo.Client {
	// Set Client Options
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func newConn() *mongo.Client {
	if root == nil {
		lock.Lock()
		defer lock.Unlock()
		if root == nil {
			root = newConnectMongo()
			fmt.Println("mongodb connection successfully")
		} else {
			return root
		}
	}
	return root
}
