package goth

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenLoginWall(ifNotAuthed func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) bool {
	return func(w http.ResponseWriter, r *http.Request) bool {		
		token, _ := r.Cookie("token")
		username, _ := r.Cookie("username")

		if token == nil || username == nil {
			ifNotAuthed(w, r)
			return false
		}

		ok, user := GetUser(bson.D{
			primitive.E{ Key: "username", Value: username.Value },
		})

		if !ok {
			ifNotAuthed(w, r)
			return false
		}

		if !(user.Token == token.Value) {
			ifNotAuthed(w, r)
			return false
		}
		
		return true
	}
}