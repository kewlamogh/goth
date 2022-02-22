package main

import (
	"fmt"
	"net/http"
	"os"
)

func StartServer() {
	http.HandleFunc("/api/total_users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprint(TotalUsers())))
	})

	http.HandleFunc("/api/hitsPerMonth", func(w http.ResponseWriter, r *http.Request) {
		hits := GetNumberOfHitsToProtectedRoutesInMonth(r.URL.Query().Get("month"))
		
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprint(hits)))
	})

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("ui/index.html")
		checkError(err)

		w.Header().Set("Content-Type", "text/html")
		w.Write(content)
	})

	http.ListenAndServe(":3000", nil)
}

func main() {
	StartServer()
}