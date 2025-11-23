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
	 IsRunning bool `json: "isRunning"` // true = timer is counting, false = timer is stopped
	 Milliseconds int64 `json:"milliseconds` 
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
	 defer.mutex.Unlock()

	 // create a response object with current timer state
	 response := TimerResponse{
			IsRunning: isRunning
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
		 isRunning = true``
	 }

	 // send the success message
	 w.Header().Set("Content-Type", "application/json")
	 w.Header().Set("Access-Control-Allow-Origin", "*")

	 // send a simple JSON object back to confirm the action
	 json.NewEncoder(w).Encode(map[sting]string{"status":"started"})
 }

 // handleStop pauses the timer
 func handleStop (w http.ResponseWriter, r *http.Request) {
	 mutex.Lock()
	 defer mutex.Unlock()

	 if IsRunning {
		 accumulated = getCurrentMilliseconds()

		 isRunning = false

		 startTime = nil
	 }
	 w.Header.Set("Content-Type", "application/json")
	 w.Header.Set("Access-Control-Allow-Origin")

	 json.NewEncoder(w).Encode(map[string]string{"status", "stopped"})
 }


