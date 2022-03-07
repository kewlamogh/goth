package goth

import (
	"go.mongodb.org/mongo-driver/bson"
)

// A data handle which you can use to edit and get user data. 
type UserDataHandle struct {
	data map[string]interface{}
	filter bson.D
}

// Generates a data handle based on the users after filtering by filter.
func NewDataHandle(filter bson.D) (UserDataHandle) {
	type data = struct{
		Data map[string]interface{} `json:"data"`
	}

	client, ctx, close := connectToMongoDB()
	defer close()

	d := UserDataHandle{
		filter: filter,
	}
	r := data{}	

	client.Database("goth").Collection("userdata").FindOne(ctx, filter).Decode(&r)
	
	d.filter = filter
	d.data = map[string]interface{}{}

	if r.Data != nil {
		for k, v := range r.Data {
			d.data[k] = v
		}
	}

	return d
}

// Returns the data handle (a pointer to a Go dictionary)
func (u *UserDataHandle) GetDataHandle() *map[string]interface{} {
	return &u.data
}

// Pushes modifications to the database.
func (u *UserDataHandle) Push() {
	client, ctx, close := connectToMongoDB()
	l := bson.D{
		bson.E{
			Key: "data",
			Value: u.data,
		},
	}

	l = append(l, u.filter...)
	defer close()

	client.Database("goth").Collection("userdate").DeleteMany(ctx, bson.D{
		bson.E{
			Key: "filter",
			Value: u.filter,
		},
	})
	client.Database("goth").Collection("userdata").InsertOne(ctx, l)
}