package main

import( 
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func userInput() string {
	fmt.Println("Enter the message you want to display")
	fmt.Print(": ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')


	userInput = strings.TrimSpace(userInput)
	return userInput
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	currentMessage := userInput()
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")


	response := Response{Message: currentMessage} 
	json.NewEncoder(w).Encode(response)
}

func main () {
	
	router := mux.NewRouter()

	router.HandleFunc("/api/message", getMessage).Methods("GET")

	log.Println("Server starting on http://localhost:8080")
	log.Println("Endpoint available:")
	log.Println("GET /api/message - Get message text")

	log.Fatal(http.ListenAndServe(":8080", router))
}
