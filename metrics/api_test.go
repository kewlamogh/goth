package metrics

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Tests the TotalUsers api.

NOTE:
When running or debugging this test, you should have a test_config.env file in your development folder.
The test_config.env should have the uri environment variable defined with a working MongoDB URI. (place your real URI, that file is in the .gitignore)
*/
func TestTotalUsers(t *testing.T) {
	client, ctx, close := connectToMongoDB()
	defer close()
	amntUsers, _ := client.Database("goth").Collection("users").CountDocuments(ctx, bson.D{})
	tests := []struct {
		name string
		want int64
	}{
		{name: "test total users function", want: amntUsers},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalUsers(); got != tt.want {
				t.Errorf("TotalUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
Tests the number of route hits api.

NOTE:
When running or debugging this test, you should have a test_config.env file in your development folder.
The test_config.env should have the uri environment variable defined with a working MongoDB URI. (place your real URI, that file is in the .gitignore)
*/
func TestGetNumberOfHitsToProtectedRoutesInMonth(t *testing.T) {
	type args struct {
		month string
	}

	tests := []struct {
		name      string
		args      args
		want      int64
		dummyData bson.D
	}{
		{name: "example feb test case", args: args{month: "February"}, dummyData: bson.D{
			primitive.E{Key: "time", Value: "Aruary"},
			primitive.E{Key: "type", Value: "hitProtectedRoute"},
			primitive.E{Key: "is_test", Value: true},
		}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, ctx, close := connectToMongoDB()
			defer close()

			client.Database("goth").Collection("metrics").InsertOne(ctx, tt.dummyData)

			if got := GetNumberOfHitsToProtectedRoutesInMonth(tt.args.month); got != tt.want {
				t.Errorf("GetNumberOfHitsToProtectedRoutesInMonth() = %v, want %v", got, tt.want)
			}

			client.Database("goth").Collection("metrics").DeleteMany(ctx, tt.dummyData)
		})
	}
}
