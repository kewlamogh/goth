package metrics

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/api/total_users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprint(TotalUsers())))
	})

	http.ListenAndServe(":3000", nil)
}
