package goth

import (
	"fmt"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The encapsulator for the login data returned by the getLoginData parameter function for GenLoginRoute. This data is typically from a form.
type LoginData struct {
	// The username provided to the login route.
	Username string
	// The password provided to the login route. This password is only used to generate the token for authentication.
	Password string
}

// Generates a login route to be given as the http handler for the /login route.
func GenLoginRoute(serve func(writer http.ResponseWriter), getLoginData func(*http.Request) LoginData, afterAuthed func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			data := getLoginData(r)
			ok, user := GetUser(bson.D{ primitive.E{ Key: "username", Value: data.Username }, })
			token := Token{ Token: user.Token }
			isEqual, err := token.IsEqualToTokenOf(data.Username, data.Password)
			
			if ok && isEqual && err == nil {
				for _, v := range []http.Cookie{
					{ Name: "token", Value: token.Token },
					{ Name: "username", Value: user.Username },
				} {
					http.SetCookie(writer, &v)
				}

				afterAuthed(writer, r)
			} else {
				withoutQuery := strings.Split(r.URL.String(), "?")[0]
				fmt.Println(r.URL.String())
				fmt.Println(withoutQuery+"?err="+err.Error())
				http.Redirect(writer, r, withoutQuery+"?err="+err.Error(), http.StatusFound)
			}
		} else {
			serve(writer)
		}
	}
}