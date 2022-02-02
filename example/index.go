package main

import (
	"goth"
	"net/http"
	"os"
)


func main() {
	goth.SetURI("mongodb+srv://jjj:L011yp0ppz@cluster0.olg0x.mongodb.net/myFirstDatabase?authSource=admin&replicaSet=atlas-12s1h0-shard-0&w=majority&readPreference=primary&retryWrites=true&ssl=true")

	var loginwall = goth.GenLoginWall(func (writer http.ResponseWriter, r *http.Request) {
		http.Redirect(writer, r, "/login", http.StatusFound)
	})

	var loginroute = goth.GenLoginRoute(func (writer http.ResponseWriter) {
		content, err := os.ReadFile("views/login.html")
		if err != nil {
			panic(err)
		}
		writer.Header().Set("Content-Type", "text/html")
		writer.Write(content)
	}, func (r *http.Request) goth.LoginData {
		return goth.LoginData{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
	}, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})

	var signuproute = goth.GenSignupRoute(func (w http.ResponseWriter) {
		content, err := os.ReadFile("views/signup.html")
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(content)
	}, func (r *http.Request) goth.SignupData {
		return goth.SignupData{	
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
	}, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	http.HandleFunc("/login", loginroute)
	http.HandleFunc("/signup", signuproute)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		loginwall(w, r)
		w.Write([]byte("Hi"))
	})

	http.ListenAndServe(":8080", nil)
}