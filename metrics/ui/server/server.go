package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kewlamogh/goth/metrics"
)

func StartServer() {
	http.HandleFunc("/api/total_users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprint(metrics.TotalUsers())))
	})

	http.HandleFunc("/api/hitsPerMonth", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		hits := metrics.GetNumberOfHitsToProtectedRoutesInMonth(metrics.MonthData{
			Month: now.Month().String(),
			Year: uint32(now.Year()),
		})

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

func main() {
	StartServer()
}
