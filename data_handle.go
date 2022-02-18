package goth

import (
	"go.mongodb.org/mongo-driver/bson"
)

type UserDataHandle struct {
	data map[string]interface{}
	filter bson.D
}

func NewDataHandle(filter bson.D) (UserDataHandle) {
	type data = struct{
		data map[string]interface{} `json:"data"`
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

	if r.data != nil {
		for k, v := range r.data {
			d.data[k] = v
		}
	}

	return d
}

func (u *UserDataHandle) GetDataHandle() *map[string]interface{} {
	return &u.data
}

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