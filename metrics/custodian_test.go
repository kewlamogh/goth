package metrics

import (
    "testing"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestChangeGranularityOfOldHitMetrics(t *testing.T) {
    type args struct {
        now MonthData
    }

    tests := []struct {
        name            string
        args            args
        data            []bson.D
        hitsWantSmushed MonthData
    }{
        {name: "test February", args: args{
            now: MonthData{
                Month: "February",
                Year:  1000,
            },
        }, data: []bson.D{
            {
                primitive.E{Key: "month", Value: "January"},
                primitive.E{Key: "year", Value: 1000},
                primitive.E{Key: "hits", Value: 10},
                primitive.E{Key: "granularity", Value: 1},
                primitive.E{Key: "is_test", Value: true},
            },
            {
                primitive.E{Key: "month", Value: "January"},
                primitive.E{Key: "year", Value: 1000},
                primitive.E{Key: "hits", Value: 11},
                primitive.E{Key: "granularity", Value: 1},
                primitive.E{Key: "is_test", Value: true},
            },
            {
                primitive.E{Key: "month", Value: "January"},
                primitive.E{Key: "year", Value: 1000},
                primitive.E{Key: "hits", Value: 25},
                primitive.E{Key: "granularity", Value: 1},
                primitive.E{Key: "is_test", Value: true},
            },
        }, hitsWantSmushed: MonthData{
            Month: "February",
            Year:  1000,
            Hits:  11 + 10 + 25,
        }},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            client, ctx, close := connectToMongoDB()
            defer close()

            for _, b := range tt.data {
                client.Database("goth").Collection("metrics").InsertOne(ctx, b)
            }

            ChangeGranularityOfOldHitMetrics(tt.args.now)
            var x MonthData

            client.Database("goth").Collection("metrics").FindOneAndDelete(ctx, bson.D{
                primitive.E{Key: "month", Value: "February"},
                // primitive.E{Key: "is_test", Value: true},
            }).Decode(&x)
                    
            if x != tt.hitsWantSmushed {
                t.Errorf("Want %v, got %v", tt.hitsWantSmushed, x)
            }
        })
    }
}