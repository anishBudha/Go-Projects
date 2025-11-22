package main

import(
	"encoding/json"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	out := map[string]any{
		"from": "usd",
		"to":		"cad",
		"amount": 10,
		"converted": 13.7,
		"rate": 1.37,
		"date": "2025-10-09",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
