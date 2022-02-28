package goth

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "golang.org/x/crypto/bcrypt"
)

// Generates a login wall that blocks entrance to the route.
// It also saves some visit-related metric data and saves it with
// a granularity level of 1 (meaning day-level granularity). The metric custodian 
// gradually reduces the granularity - for example, every month, the previous
// month's day-level data will be scrapped for the month level data.
func GenLoginWall(ifNotAuthed func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) bool {
	return func(w http.ResponseWriter, r *http.Request) bool {		
		token, _ := r.Cookie("token")
		username, _ := r.Cookie("username")
		client, ctx, close := connectToMongoDB()
		defer close()

		if token == nil || username == nil {
			ifNotAuthed(w, r)
			return false
		}

		ok, user := GetUser(bson.D{
			primitive.E{ Key: "username", Value: username.Value },
		})

		if !ok {
			ifNotAuthed(w, r)
			return false
		}

		if user.Token != token.Value {
			ifNotAuthed(w, r)
			return false
		}

		currentTime := time.Now()
		var d struct{
			Year uint32 `json:"year"`
			Month string `json:"month"`
			Day uint32 `json:"day"`
			Hits int `json:"hits"`
		}

		dateBson := bson.D{
			primitive.E{ Key: "year", Value: currentTime.Year() },
			primitive.E{ Key: "month", Value: currentTime.Month().String() },
			primitive.E{ Key: "day", Value: currentTime.Day() },
		}

		client.Database("goth").Collection("metrics").FindOneAndDelete(ctx, dateBson).Decode(&d)

		if d.Year == 0 {
			d.Year = uint32(currentTime.Year())
			d.Month = currentTime.Month().String()
			d.Day = uint32(currentTime.Day())
			
			d.Hits++
		}

		obj := dateBson
		obj = append(obj, primitive.E{ Key: "hits", Value: d.Hits })
		obj = append(obj, primitive.E{ Key: "granularity", Value: 1 })

		client.Database("goth").Collection("metrics").InsertOne(ctx, obj)
		return true
	}
}