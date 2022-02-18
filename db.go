package goth

import (
	"context"
	"os"
	"time"
	goenv "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var URI = (func() string {
	goenv.Load("test_config.env")
	return os.Getenv("uri")
})()

// Configures the URI
func SetURI(uri string) {
	URI = uri
}

// Utility to return a Mongo client.
func connectToMongoDB() (*mongo.Client, context.Context, func()) {
	if URI == "" {
		panic("Invalid URI")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	checkError(err)
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	checkError(err)
	return client, ctx, close
}

// Model of a user.
type User struct {
	// Username of the user.
	Username string `json:"username"`
	// User token (generated by the GenToken function).
	Token string `json:"token"`
}