package goth

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Tests the GetUser method.

NOTE: 
When running or debugging this test, you should have a test_config.env file in your development folder.
The test_config.env should have the uri environment variable defined with a working MongoDB URI. (place your real URI, that file is in the .gitignore)
*/
func TestGetUser(t *testing.T) {
	ok, user := GetUser(bson.D{
		primitive.E{ Key: "username", Value: "joe" }, 
	})	

	if (!ok && user != (User{})) || (ok && user == User{}) {
		t.Errorf("innacurate ok. is ok: %t. user: %v", ok, user)
	}
}