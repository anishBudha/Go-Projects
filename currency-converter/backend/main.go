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

const {
	// endpoint %s
	jsdelivrTemplate =  "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/%s.min.json"
	// fallback url on cloudflare
	pagesDevTemplate = "https://latest.currency-api.pages.dev/v1/currencies/%s.json"
}

// structure expected fromt the remote API 
type RatesResponse struct {
	Date string								`json:"date"`
	Rates map[string]float64 	`json:"rates"`
	Base string 							`json:"base"`
}

func fetchRates(base string) (*RateRespone, error) {
	client := &http.Client{Timeout: 8 * time.Second}

	// try primary CDN (jsdelivr)
	u := fmt.Sprintf(jsdelivrTemplate, url.PathEscape(base))
	resp, err := client.Get(u)
	if err != nil || resp.StatusCode != http.StatusOK {
		// attempt fallback to pages.dev
		if resp != nil {
			resp.Body.Close() // closing the connetion 
		}
		u = fmt.Sprintf(pagesDevTemplate, url.PathEscape(base)) // PathEscape sanitizes for url like adding this%20is
		resp, err = client.Get(u)
		if err != nil {
			return nil, err
		}
	}
	defer resp.Body.Close() //dev

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r RatesResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	// some variants may omit Base; ensure its set

	if r.Base == "" {
		r.Base = base
	}

	return &r, nil
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	from := q.Get("from")
	to := q.Get("to")
	amountStr := q.Get("amount")

	// if any of these are empty, respond with 400 Bad Request and a message
	if from == "" || to == "" || amountStr == "" {
		http.Error(w, "missing query parameters", http.StatusBadRequest)
		return
	}

	// parse amountStr into numeric amount
	var amount float64
	_, err := fmt.Sscan(amountStr, &amount)
	if err != nil {
		http.Error(w, "invalid amount", http.StatusBadRequest); 
		return
	}

	// fetch rates with base = from
	ratesResp, err := fectRates(from)
	if err != nil {
		log.Printf("fetchRates error: %v\n", err)
		http.Error(w, "failed to fetch rates", http.StatusInternalServerError)
		return
	}

	toRate, ok := ratesResp.Rates[to]
	if !ok {
		http.Error(w, "target currency not found", http.StatusBadRequest)
		return
	}

	// If API returns rates relative to base, then rate[to] is direct multiplier.

	converted := amount * toRate 
	
	out := map[string]any {
		"from": from,
		"to": to,
		"amount": amount,
		"converted": converted,
		"date": ratesResp.Date,
		"rate": toRate,
	}
	w.Header().Set("Contect-Type", "application/json") // this is json
	json.NewEncoder(w).Encode(out) // here is json!
}


func main() {
	




}



















