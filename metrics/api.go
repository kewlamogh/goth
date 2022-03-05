package metrics

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Data model for which is stored every day in the database
type DayData struct {
	Year uint32 `json:"year"`
	Month string `json:"month"`
	Day uint32 `json:"day"`
	Hits int `json:"hits"`
}

// Month data is stored every month in the database
// when the granular DayData is scrapped.
type MonthData struct {
	Year uint32 `json:"year"`
	Month string `json:"month"`
	Hits int `json:"hits"`
}

// Returns the total number of users in the goth/users MongoDB collection.
func TotalUsers() int64 {
	client, ctx, close := connectToMongoDB()
	defer close()
	amnt, _ := client.Database("goth").Collection("users").CountDocuments(ctx, bson.D{})
	return amnt
}

// Gets the number of hits to a protected route in a given month.
func GetNumberOfHitsToProtectedRoutesInMonth(data MonthData) int64 {
	client, ctx, close := connectToMongoDB()
	hits := 0

	defer close()

	cur, _ := client.Database("goth").Collection("metrics").Find(ctx, bson.D{
		primitive.E{ Key: "month", Value: data.Month },
		primitive.E{ Key: "year", Value: data.Year },
	})

	for cur.Next(ctx) {
		var d DayData
		cur.Decode(&d)

		hits += d.Hits
	}
	
	return int64(hits)
}
