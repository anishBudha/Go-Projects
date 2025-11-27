package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync" // for synchronization primitives (like Mutex)
)

// When multiple requests come at the same time, the Mutex ensures only one modifies the counter at a time
type Counter struct {
	Value int
	mu sync.Mutex
}

var counter = Counter{Value: 0}

type Response struct {
	Count int `json:"count"`
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	counter.mu.Lock()
	currentValue := counter.Value
	counter.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")

	// create a response and send it as JSON
	respone := Response{Count: currentValue}
	json.NewEncoder(w).Encode(respone)
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	counter.mu.Lock()
	counter.Value++
	currentValue := counter.Value
	counter.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	response := Response {Count: currentValue}
	json.NewEncoder(w).Encode(response)
}
func decrementCounter(w http.ResponseWriter, r *http.Request) {
	counter.mu.Lock()
	counter.Value--
	currentValue := counter.Value
	counter.mu.Unlock()
	
	w.Header().Set("Content-Type", "application/json")
	response := Response{Count: currentValue}
	json.NewEncoder(w).Encode(response)
}
// enableCORS is middleware that adds CORS headers to response
// Cross-Origin Resource Sharing allows our frontend to talk to our backend
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin (in production we replace it with domain)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")	
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests (browsers send these before actual requests)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// call the actual handler function
		next(w, r)
	}
}

func main() {
	// serve static files from the static folder 
	// when someone visits the root URL, they'll get our frontend
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/api/counter", enableCORS(getCounter))
	http.HandleFunc("/api/counter/increment", enableCORS(incrementCounter))
	http.HandleFunc("/api/counter/decrement", enableCORS(decrementCounter))

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Serving frontend from ./static folder")
	fmt.Println("Counter API ready at /api/counter")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
