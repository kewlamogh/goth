package goth

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSetURI(t *testing.T) {
	tests := []struct {
		name string
		uri  string
	}{
		{name: "test setting uri", uri: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldURI := URI
			SetURI(tt.uri)
			defer (func() { URI = oldURI })()

			if URI != tt.uri {
				t.Errorf("Set uri failed. Got %v, want %v", URI, tt.uri)
			}
		})
	}
}

/*
Tests the MongoDB connect utility function.

NOTE:
When running or debugging this test, you should have a test_config.env file in your development folder.
The test_config.env should have the uri environment variable defined with a working MongoDB URI. (place your real URI, that file is in the .gitignore)
*/
func Test_connectToMongoDB(t *testing.T) {
	tests := []struct {
		name   string
		insert bson.D
		filter bson.D
	}{
		{name: "test standard crd with c:d filter", insert: bson.D{primitive.E{Key: "A", Value: "B"}, primitive.E{Key: "C", Value: "D"}}, filter: bson.D{primitive.E{Key: "C", Value: "D"}}},
		{name: "test standard crd with a:b filter", insert: bson.D{primitive.E{Key: "A", Value: "B"}, primitive.E{Key: "C", Value: "D"}}, filter: bson.D{primitive.E{Key: "A", Value: "B"}}},
		{name: "test standard crd with entire document filter", insert: bson.D{primitive.E{Key: "A", Value: "B"}, primitive.E{Key: "C", Value: "D"}}, filter: bson.D{primitive.E{Key: "A", Value: "B"}, primitive.E{Key: "C", Value: "D"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, ctx, close := connectToMongoDB()
			defer close()

			_, err := client.Database("goth").Collection("sampledata").InsertOne(ctx, tt.insert)
			checkError(err)

			var obj struct {
				A string `json:"A"`
				C string `json:"C"`
			}

			client.Database("goth").Collection("sampledata").FindOne(ctx, tt.filter).Decode(&obj)

			if obj.A != "B" || obj.C != "D" {
				t.Logf("Either inserting or decoding failed.\nobj.A: %s\nobj.C:%s", obj.A, obj.C)
			}

			_, err = client.Database("goth").Collection("sampledata").DeleteOne(ctx, tt.filter)
			checkError(err)

			client.Database("goth").Collection("sampledata").FindOne(ctx, tt.filter).Decode(&obj)

			if obj.A != "" || obj.C != "" {
				t.Logf("Either delete or decoding failed. obj should be empty. \nobj.A: %s\nobj.C:%s", obj.A, obj.C)
			}
		})
	}
}
