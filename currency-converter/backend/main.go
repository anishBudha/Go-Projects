package main

import (
	"fmt"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const ( 

	jsdelivrTemplate = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/latest/%s.min.json"
	pagesDevTemplate = "https://latest.currency-api.pages.dev/v1/latest/%s.json"
	// Note: README recommends using a fallback mechanism. See repo.  [oai_citation:1‡GitHub](https://github.com/fawazahmed0/exchange-api)

)


// Go is a strongly typed language. It needs to know exactly what data looks like before it receives it.

type RatesResponse struct {
	Date string								`json:"date"`
	Rates map[string]float64	`json:"rates"`
	Base string								`json:"base"`
}

// worker functions

// fetchRates function 
// create a HTTP Client
// try the primary url : jsdelivr 
// if not try the fallback pages.dev
// unmarshal the response 
// defer resp.Body.Close() before function finishes, make sure the connection is closed
func fetchRates(base string) (*RatesResponse, error) {
	
	client := &http.Client{Timeout: 8 * time.Second}

	// try primary CDN jsdelivr 
	u := fmt.Sprintf(jsdelivrTemplate, url.PathEscape(base)) // Sprintf takes the template url and replaces the %s with the base
	resp, err := client.Get(u)

	if err != nil || resp.StatusCode != http.StatusOK {
		// attempt fallback to pages.dev 
		if resp != nil {
		resp.Body.Close() // close the connection
		}

		u = fmt.Sprintf(pagesDevTemplate, url.PathEscape(base))
		resp, err = client.Get(u)
		if err != nil {
			return nil, err 
		}
	}
	// waith until fetchRates is completely finished, then run 
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err	
	}

	var r RatesResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	// if api omit the Base currency 
	if r.Base == "" {
		r.Base = base
	}
	return &r, nil 
}


// function convertHandler is a HTTP Handler 
// parsing the request
// verify the information if something is missing
// calls the worker function fetchRates
// calculates the final amount
// respond in JSON format

func convertHandler(w http.ResponseWriter, r *http.Request) {
	// w used to send data back to the user
	// r used the package the containing details about what user sent

	q := r.URL.Query()
	from := q.Get("from")
	to := q.Get("to")
	amountStr := q.Get("amount")
	if from == "" || to == "" || amountStr == "" {
		http.Error(w, "missing query parameters", http.StatusBadRequest)
		return
	}
	
	// parse amount 
	var amount float64 
	_, err := fmt.Sscan(amountStr, &amount)
	if err != nil {
		http.Error(w, "invalid amount", http.StatusBadRequest)
		return
	}

	// fetch rates with base = from
	ratesResp, err := fetchRates(from)
	if err != nil {
		log.Printf("fetchRates error: %v\n", err)
		http.Error(w, "failed to fetch rates", http.StatusInternalServerError)
		return
	}

	toRate, ok := ratesResp.Rates[to]
	if !ok {
		http.Error(w, "target currency not found", http.StatusBadRequest)
	}

	// if API returns rates relative to base, then rate[to] is direct multiplier
	converted := amount *toRate

	out := map[string]any{
		"from": 				from,
		"to":						to,
		"amount":				amount,
		"converted":		converted,
		"date":					ratesResp.Date,
		"rate":					toRate,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/api/convert", convertHandler)

	// optionally serve frontend static in prodcution from ../frontend/dist

	fs := http.FileServer(http.Dir("../frontend/dist"))
	http.Handle("/", fs) 

	srv := &http.Server {
		Addr: ":" + port,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Printf("backend listening on: %s", port)
	log.Fatal(srv.ListenAndServe())
}
