package metrics

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TotalUsers() int64 {
	client, ctx, close := connectToMongoDB()
	defer close()
	amnt, _ := client.Database("goth").Collection("users").CountDocuments(ctx, bson.D{})
	return amnt
}

func GetNumberOfHitsToProtectedRoutesInMonth(month string) int64 {
	client, ctx, close := connectToMongoDB()
	defer close()

	res, err := client.Database("goth").Collection("metrics").CountDocuments(ctx, bson.D{
		primitive.E{Key: "time", Value: month},
		primitive.E{Key: "type", Value: "hitProtectedRoute"},
	})

	checkError(err)
	return res
}
