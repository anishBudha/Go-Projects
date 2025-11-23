package main

import (
	"fmt"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const ( 

	jsdelivrTemplate = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/%s.json"
	pagesDevTemplate = "https://latest.currency-api.pages.dev/v1/currencies/%s.json"
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

	// API requires lowercase currency codes
	baseLower := strings.ToLower(base)

	// try primary CDN jsdelivr 
	primaryURL := fmt.Sprintf(jsdelivrTemplate, url.PathEscape(baseLower)) // Sprintf takes the template url and replaces the %s with the base
	resp, err := client.Get(primaryURL)
	var primaryErr error

	if err != nil {
		primaryErr = fmt.Errorf("primary API request failed: %w", err)
	} else if resp.StatusCode != http.StatusOK {
		primaryErr = fmt.Errorf("primary API returned status %d", resp.StatusCode)
		if resp != nil {
			resp.Body.Close()
		}
	}

	if primaryErr != nil {
		// attempt fallback to pages.dev 
		fallbackURL := fmt.Sprintf(pagesDevTemplate, url.PathEscape(baseLower))
		resp, err = client.Get(fallbackURL)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch rates for currency '%s': primary error (%v), fallback error: %w", base, primaryErr, err)
		}
		
		// check fallback response status code
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("failed to fetch rates for currency '%s': primary error (%v), fallback returned status %d (URL: %s)", base, primaryErr, resp.StatusCode, fallbackURL)
		}
	}
	// waith until fetchRates is completely finished, then run 
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body for currency '%s': %w", base, err)
	}
	
	// check if body is empty
	if len(body) == 0 {
		return nil, fmt.Errorf("received empty response from API for currency '%s'", base)
	}

	// New API format: {"date": "...", "{currency}": {...rates...}}
	var rawData map[string]interface{}
	if err := json.Unmarshal(body, &rawData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response for currency '%s': %w", base, err)
	}

	// Extract date
	date, ok := rawData["date"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid date field in response for currency '%s'", base)
	}

	// Extract rates from the currency key (e.g., "usd", "eur")
	ratesMap, ok := rawData[baseLower].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("missing or invalid rates data for currency '%s'", base)
	}

	// Convert rates to map[string]float64
	rates := make(map[string]float64)
	for key, value := range ratesMap {
		if rate, ok := value.(float64); ok {
			rates[key] = rate
		}
	}

	return &RatesResponse{
		Date:  date,
		Rates: rates,
		Base:  base,
	}, nil
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

	// API returns currency codes in lowercase, so convert to lowercase for lookup
	toLower := strings.ToLower(to)
	toRate, ok := ratesResp.Rates[toLower]
	if !ok {
		http.Error(w, "target currency not found", http.StatusBadRequest)
		return
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
