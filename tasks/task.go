package tasks

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartUserCountLogger(ctx context.Context, db *mongo.Database) {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				count, err := db.Collection("user").CountDocuments(ctx, bson.M{})
				if err != nil {
					log.Printf("❌ Failed to count users: %v", err)
					continue
				}
				log.Printf("📊 Total users: %d", count)

			case <-ctx.Done():
				log.Println("🛑 Stopping user count logger")
				return
			}
		}
	}()
}
