package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/kewlamogh/goth/metrics"
)

func StartServer() {
	http.HandleFunc("/api/total_users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprint(metrics.TotalUsers())))
	})

	http.HandleFunc("/api/hitsPerMonth", func(w http.ResponseWriter, r *http.Request) {
		hits := metrics.GetNumberOfHitsToProtectedRoutesInMonth(r.URL.Query().Get("month"))

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprint(hits)))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("ui/index.html")
		if err != nil {
			panic(err)
		}
		

		w.Header().Set("Content-Type", "text/html")
		w.Write(content)
	})

	http.ListenAndServe(":3000", nil)
}
