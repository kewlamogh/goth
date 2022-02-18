package goth

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Tests the MongoDB connect utility function.

NOTE:
When running or debugging this test, you should have a test_config.env file in your development folder.
The test_config.env should have the uri environment variable defined with a working MongoDB URI. (place your real URI, that file is in the .gitignore)
*/
func TestMongoDBConnect(t *testing.T) {
	client, ctx, close := connectToMongoDB()
	defer close()

	if recover() != nil {
		t.Error("failed to connect to mongodb")
	}

	obj := bson.D{ 
		primitive.E{ Key: "a", Value: "b" },
	}

	retv := struct{
		A string `json:"a"`
	}{}

	client.Database("goth").Collection("sample").DeleteOne(ctx, obj)
	client.Database("goth").Collection("sample").InsertOne(ctx, obj)
	client.Database("goth").Collection("sample").FindOne(ctx, obj).Decode(&retv)

	if retv.A == "" {
		t.Errorf("crd failed: retv.A was %s", retv.A)
	}
	
	client.Database("goth").Collection("sample").DeleteOne(ctx, obj)
}