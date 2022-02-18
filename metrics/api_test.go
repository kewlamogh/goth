package metrics

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
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
