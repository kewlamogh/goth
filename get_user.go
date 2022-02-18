package goth

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Gets a user filtered by the given bson.D.
func GetUser(filter bson.D) (bool, User) {
	user := User{}
	client, ctx, close := connectToMongoDB()
	defer close()
	client.Database("goth").Collection("users").FindOne(ctx, filter).Decode(&user)

	return user != User{}, user 
}