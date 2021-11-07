package utils

import (
	"context"
	"crawl_data/conf"
	"fmt"
	"log"

	// "gitlab.com/goxp/cloud0/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout = 5
)

func GetConnection() *mongo.Client {
	mongoConnectStr := conf.LoadEnv().MongoConnectStr
	fmt.Println(mongoConnectStr)
	// connectionURI := fmt.Sprintf(mongoConnectStr)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoConnectStr))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
