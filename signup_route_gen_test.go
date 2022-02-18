package goth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Tests the output of the GenSignupRoute function.

NOTE:
When running or debugging this test, you should have a test_config.env file in your development folder.
The test_config.env should have the uri environment variable defined with a working MongoDB URI. (place your real URI, that file is in the .gitignore)*/
func TestSignupRoute(t *testing.T) {
	message := ""
	client, ctx, close := connectToMongoDB()
	user := bson.D{
		primitive.E{
			Key: "username",
			Value: "bob",
		},
	}

	client.Database("goth").Collection("users").DeleteMany(ctx, user)
	close()

	signupRoute := GenSignupRoute(func (_ http.ResponseWriter) {
		message = "Client side UI loaded."
	}, func (r *http.Request) SignupData {
		data := SignupData{
			Username: "bob",
			Password: "john",
		}

		return data
	}, func (writer http.ResponseWriter, r *http.Request) {
		message = "Successfully signed up"
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "https://example.com/signup", nil)
	signupRoute(recorder, request)

	if message != "Successfully signed up" {
		t.Errorf("Did not successfully sign up. message: %s", message)
	}
}