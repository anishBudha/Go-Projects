 package main

 import (
	 "encoding/json"
	 "log"
	 "net/http"
	 "sync" // for synchroniztion - helps prevent race conditions
	 "time"
	 
	 "github.com/gorilla/mux" // external router package for handling different URL paths
 )

 type TimerResponse struct {
	 IsRunning bool `json:"isRunning"` // true = timer is counting, false = timer is stopped
	 Milliseconds int64 `json:"milliseconds"` 
 }

 // Global variables

 var (
	 isRunning bool = false
	 // using a pointer so that it can be nil when timer is not running
	 startTime *time.Time = nil
	 accumulated int64 = 0
	 // mutex prevents race conditions, where two operations try to change data simultaneously
	 mutex sync.Mutex
 )

// Helper functions

// getCurrentMilliseconds() calculates the total elapsed time in milliseconds
func getCurrentMilliseconds() int64 {
	if isRunning && startTime != nil {
		return accumulated + time.Since(*startTime).Milliseconds()
	}
	return accumulated
}
 // HTTP handler functions
 // handleGetTimer sends the current timer state to the frontend
 // called when frontend requests: GET http://localhost:8080/api/timer
 // w = http.ResponseWriter - where we write our response (output)
 // r = http.Request - contains information about the incoming request

 func handleGetTimer (w http.ResponseWriter, r *http.Request) {
	 // Lock the mutex to prevent other requests from changing their timer
	 mutex.Lock()

	 // when function ends
	 defer mutex.Unlock()

	 // create a response object with current timer state
	 response := TimerResponse{
			IsRunning: isRunning,
			Milliseconds: getCurrentMilliseconds(),
	 }

	 // Set response header
	 w.Header().Set("Content-Type", "application/json")
	 w.Header().Set("Access-Control-Allow-Origin", "*") // allow frontend on different port to access

	 // converting Go struct response in JSON format and writing it to reponse(w)
	 json.NewEncoder(w).Encode(response)
 }

 // handleStart starts the timer
 // POST http://localhost:8080/api/timer/start
 func handleStart(w http.ResponseWriter, r *http.Request) {
	 mutex.Lock()
	 
	 defer mutex.Unlock()

	 // only start is the timer is not running
	 if !isRunning {
		 // record the current time as the start time
		 now := time.Now()
		 startTime = &now

		 // mark the timer
		 isRunning = true
	 }

	 // send the success message
	 w.Header().Set("Content-Type", "application/json")
	 w.Header().Set("Access-Control-Allow-Origin", "*")

	 // send a simple JSON object back to confirm the action
	 json.NewEncoder(w).Encode(map[string]string{"status":"started"})
 }

 // handleStop pauses the timer
 func handleStop (w http.ResponseWriter, r *http.Request) {
	 mutex.Lock()
	 defer mutex.Unlock()

	 if isRunning {
		 accumulated = getCurrentMilliseconds()

		 isRunning = false

		 startTime = nil
	 }
	 w.Header().Set("Content-Type", "application/json")
	 w.Header().Set("Access-Control-Allow-Origin", "*")

	 json.NewEncoder(w).Encode(map[string]string{"status": "stopped"})
 }

// handleReset resets the timer back to 0
func handleReset (w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	// Reset to initial state
	isRunning = false
	startTime = nil
	accumulated = 0

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(map[string]string{"status": "reset"})
}


// handleOptions handles CORS = Cross Origin Resource Sharing
// Browsers sends an OPTIONS request before POST requests to check if it's allowed
// browser checks if frontend at port 5273 is allowed to talk to backend at port 8080.
// This functions says "yes, it is allowed"

func handleOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
}

// main function

func main () {
	// router controls and directs incoming requests to the correct handler function based on the URL path
	router := mux.NewRouter()

	// register URL paths with their handler functions

	// GET /api/timer - Get current timer state
	router.HandleFunc("/api/timer", handleGetTimer).Methods("GET")

	// POST /api/timer/start - Start the timer
	router.HandleFunc("/api/timer/start", handleStart).Methods("POST")

	// POST /api/timer/stop - Stop the timer
	router.HandleFunc("/api/timer/stop", handleStop).Methods("POST")

	// POST /api/timer/reset - Reset the timer
	router.HandleFunc("/api/timer/reset", handleReset).Methods("POST")

	// OPTIONS for all routes - Handle CORS preflight
	router.Methods("OPTIONS").HandlerFunc(handleOptions)

	// Print a message so we know the server started successfully
	log.Println("Server starting on http://localhost:8080")
	log.Println("Endpoints available:")
	log.Println("	GET /api/timer - Get timer state")
	log.Println("	POST /api/timer/start - Start timer")
	log.Println("	POST /api/timer/stop - Stop timer")
	log.Println("	POST /api/timer/reset - Reset timer")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
