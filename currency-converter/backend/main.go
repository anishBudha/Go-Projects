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

func main() {
	




}
















