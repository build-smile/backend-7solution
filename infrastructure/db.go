package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

// ConnectDB connects to MongoDB and assigns the global client and database
func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(CFG.MongoDB.ConnectTimeoutMilli)*time.Second)
	defer cancel()

	// Set client options
	clientOpts := options.Client().ApplyURI(CFG.MongoDB.Uri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	// Ping to test connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ MongoDB ping failed: %v", err)
	}

	fmt.Println("✅ Connected to MongoDB!")

	// Set global client and database
	MongoClient = client
	MongoDB = client.Database(CFG.MongoDB.Database)
}
