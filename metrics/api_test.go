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
		month MonthData
	}

	tests := []struct {
		name      string
		args      args
		want      int64
	}{
		{name: "example ar test case", args: args{month: MonthData{
			Month: "Aruary",
			Year: 1212,
		}}, want: 1},	
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, ctx, close := connectToMongoDB()
			defer close()

			client.Database("goth").Collection("metrics").InsertOne(ctx, bson.D{
				primitive.E{ Key: "month", Value: tt.args.month.Month },
				primitive.E{ Key: "year", Value: tt.args.month.Year },
				primitive.E{ Key: "hits", Value: 1 },
				primitive.E{ Key: "is_test", Value: true },
			})

			if got := GetNumberOfHitsToProtectedRoutesInMonth(tt.args.month); got != tt.want {
				t.Errorf("GetNumberOfHitsToProtectedRoutesInMonth() = %v, want %v", got, tt.want)
			}

			client.Database("goth").Collection("metrics").DeleteMany(ctx, bson.D{
				primitive.E{ Key: "month", Value: tt.args.month.Month },
				primitive.E{ Key: "year", Value: tt.args.month.Year },
				primitive.E{ Key: "is_test", Value: true },
			})
		})
	}
}
