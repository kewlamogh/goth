package metrics

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Internal list of months:
var months = []string{
	"January",
	"February",
	"March",
	"April", 
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

// Changes the granularity of old items.
// After a month, the day-level granularity
// will be compressed into month-level granularity.
// Month-level granularity is stored forever.
func ChangeGranularityOfOldHitMetrics(now MonthData) {
	month := now.Month
	year := now.Year
	
	oldIndex := -1
	client, ctx, close := connectToMongoDB()
	smushed := now

	defer close()

	var target string

	for i, v := range months {
		if v == month {
			target = months[oldIndex]
			break
		}

		oldIndex = i
	}

	if oldIndex == -1 {
		target = "December"
		year--
	}

		
	cur, _ := client.Database("goth").Collection("metrics").Find(ctx, bson.D{
		primitive.E{ Key: "month", Value: target },
		primitive.E{ Key: "year", Value: year },
		primitive.E{ Key: "granularity", Value: 1 },
	})

	smushed.Hits = 0
	for cur.Next(ctx) {
		var d DayData
		cur.Decode(&d)
		
		smushed.Hits += d.Hits
	}

	client.Database("goth").Collection("metrics").DeleteMany(ctx, bson.D{
		primitive.E{ Key: "month", Value: target },
		primitive.E{ Key: "year", Value: year },
		primitive.E{ Key: "granularity", Value: 1 },
	})

	client.Database("goth").Collection("metrics").InsertOne(ctx, bson.D{
		primitive.E{ Key: "month", Value: smushed.Month },
		primitive.E{ Key: "year", Value: smushed.Year },
		primitive.E{ Key: "hits", Value: smushed.Hits },
		primitive.E{ Key: "granularity", Value: 2 },
	})
}
