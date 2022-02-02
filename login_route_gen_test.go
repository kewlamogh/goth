package goth

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Tests the login route generator's output.

NOTE:
When running or debugging this test, you should have a test_config.env file in your development folder.
The test_config.env should have the uri environment variable defined with a working MongoDB URI. (place your real URI, that file is in the .gitignore)
*/
func TestLoginRoute(t *testing.T) {
	message := "" 
	client, ctx, close := connectToMongoDB()
	user := bson.D{
		primitive.E{
			Key: "username",
			Value: "bob",
		}, primitive.E{
			Key: "token",
			Value: GenToken("bob", "john").Token,
		},
	}

	defer close()

	client.Database("goth").Collection("users").DeleteOne(ctx, user)
	client.Database("goth").Collection("users").InsertOne(ctx, user)

	loginHandler := GenLoginRoute(func (_ http.ResponseWriter) {
		message = "Client side UI loaded."
	}, func(request *http.Request) LoginData {
		data := LoginData{
			Username: "bob",
			Password: "john",
		}

		return data
	}, func (_ http.ResponseWriter, _ *http.Request) {
		message = "Authed successfully!"
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "https://example.com/login", nil)
	loginHandler(recorder, request)

	if message != "Authed successfully!" {
		t.Error("Not authed correctly")
	}
}