package goth

import (
	"net/http"
	"strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The encapsulator for the signup data returned by the getSignupData parameter function for GenSignupRoute. This data is typically from a form.
type SignupData struct {
	// The username provided to the signup route.
	Username string
	// The password provided to the signup route. This password is only used to generate the token for authentication.
	Password string
}

// Generates a signup route to be used as the http handler for the /signup route
func GenSignupRoute(serve func(http.ResponseWriter), getSignupData func (*http.Request) SignupData, successfullyCreated func(http.ResponseWriter, *http.Request)) func(writer http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			data := getSignupData(r)
			ok, _ := GetUser(bson.D{ primitive.E{ Key: "username", Value: data.Username }, })

			if !ok {
				client, ctx, close := connectToMongoDB()
				defer close()

				client.Database("goth").Collection("users").InsertOne(ctx, bson.D{
					primitive.E{ Key: "username", Value: data.Username },
					primitive.E{ Key: "token", Value: GenToken(data.Username, data.Password).Token },
				})

				successfullyCreated(writer, r)
			} else {
				withoutQuery := strings.Split(r.URL.String(), "?")[0]
				http.Redirect(writer, r, withoutQuery+"?err=username taken", http.StatusFound)
			}
		} else {
			serve(writer)
		}
	}
}