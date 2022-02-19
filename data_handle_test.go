package goth

import (
	"testing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gotest.tools/assert"
)


/*
Tests the User Data feature.


NOTE:
When running or debugging this test, you should have a test_config.env file in your development folder.
The test_config.env should have the uri environment variable defined with a working MongoDB URI. (place your real URI, that file is in the .gitignore)
*/
func TestUserData(t *testing.T) {
	type data = struct {
		Data map[string]interface{} `json:"data"`
	}

	datahandle := NewDataHandle(bson.D{
		primitive.E{Key: "username", Value: "hinon"},
	})

	datahandle.data["x"] = "y"
	datahandle.Push()

	client, ctx, close := connectToMongoDB()
	defer close()

	var x data
	client.Database("goth").Collection("userdata").FindOne(ctx, datahandle.filter).Decode(&x)

	v := x.Data["x"]
	assert.Equal(t, v, "y")

	client.Database("goth").Collection("userdata").DeleteMany(ctx, datahandle.filter)
}
