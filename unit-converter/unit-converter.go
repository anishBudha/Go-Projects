package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// holds the data we'll send back to the user
type ConversionResult struct {
	Value float64
	FromUnit string
	ToUnit string
	Category string
	InputValue float64
}

func main () {
	// route handlers, these connect URLs to functions
	http.HandleFunc("/", homeHandler) // main page
	http.HandleFunc("/convert", convertHandler) // handles conversion requests

	// starting webserver in port 8080
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// define our html template as a string
	// the {{}} syntax is Go's template language for dynamic content
	tmpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=, initial-scale=1.0">
		<title>Unit Converter</title>
		<!-- Include HTMX Library from CDN -->
		<script src="http://unpkg.com/htmx.org@1.9.10"></script>
		<style>



		</style>
	</head>
	<body>
		<div class="container">
		<h1>Unit Converter</h1>
		<!-- 
			HTMX attributes explained:
			hx-post="/convert" - sends a POST request to /convert endpoint
			hx-target="#request" - puts the response into the element with id="result"
			hx-trigger="submit" - triggers when form is submitted
			hx-indicator=".htmx-indicator" - shows loading indicator
		-->
		<form hx-post="/convert" hx-target="#result" hx-trigger="submit" hx-indicator=".htmx-indicator">
			<div class="form-group">
			<label for="category">Category:</label>
			<select name="category" id="category" required>
				<option value="length">Length</option>
				<option value="weight">Weight</option>
				<option value="temperature">Temperature</option>
				<option value="volume">Volume</option>
			</select>
		</div>
	</body>
	`
}
