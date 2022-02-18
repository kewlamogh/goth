package goth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Tests the login wall generator's output.

NOTE:
When running or debugging this test, you should have a test_config.env file in your development folder.
The test_config.env should have the uri environment variable defined with a working MongoDB URI. (place your real URI, that file is in the .gitignore)
*/
func TestLoginWall(t *testing.T) {
	message := ""
	client, ctx, close := connectToMongoDB()
	token := GenToken("bob", "john").Token
	user := bson.D{
		primitive.E{
			Key: "username",
			Value: "bob",
		}, primitive.E{
			Key: "token",
			Value: token,
		},
	}

	client.Database("goth").Collection("users").DeleteMany(ctx, user)
	client.Database("goth").Collection("users").InsertOne(ctx, user)
	defer close()
	loginWall := GenLoginWall(func (_ http.ResponseWriter, _ *http.Request) {
		message = "not authed"
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)

	loginWall(recorder, request)
	if message != "not authed" {
		t.Error("didn't mark as not authed")
		return
	}

	recorder = httptest.NewRecorder()
	request = httptest.NewRequest("GET", "/", nil)

	request.AddCookie(&http.Cookie{
		Name: "username",
		Value: "bob",
	})

	request.AddCookie(&http.Cookie{
		Name: "token",
		Value: token,
	})

	message = ""
	loginWall(recorder, request)
	if message == "not authed" {
		t.Error("didn't let it pass when it should have")
	}
	
	client.Database("goth").Collection("users").DeleteMany(ctx, user)
}