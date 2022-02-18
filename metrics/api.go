package metrics

import (
	"go.mongodb.org/mongo-driver/bson"
)

func TotalUsers() int64 {
	client, ctx, close := connectToMongoDB()
	defer close()
	amnt, _ := client.Database("goth").Collection("users").CountDocuments(ctx, bson.D{})
	return amnt
}
