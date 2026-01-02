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
			* {
				margin: 0;
				padding: 0;
				box-sizing: border-box;
			}
			body {
				font-family: 'Segoe UI', Tahoma, Genevam Verdana, sans-serif;
				background: linear-gradient(135deg, white 0%, gray 100%);
				min-height: 100vh;
				display: flex;
				justify-content: center;
				align-items: center;
				padding: 20px;
			}
			.container {
				background: white;
				padding: 40px;
				border-radius: 20px;
				box-shadow: 0 20px 60px rgba(0,0,0,0,3);
				max-width: 500px;
				width: 100%;
			}
			h1 {
				color: black;
				text-align: center;
				margin-bottom: 30px;
				font-size: 2em;
			}
			.form-group {
				margin-bottom: 20px;
			}
			label {
				display: block;
				margin-bottom: 8px;
				color: #333;
				font-weight: 600;
			}
			input, select {
				width: 100%;
				padding: 12px;
				border: 2px solid #e0e0e0;
				border-radius: 8px;
				font-size: 16px;
				transition: border-color 0.3s;
			}
			input:focus, select:focus {
				outline: none;
				border-color: black;
			}
			button {
				width: 100%;
				padding: 14px;
				background: linear-gradient(135degm #667eea 0%, #764ba2 100%);
				color: white;
				border: none;
				border-radius: 8px;
				font-size: 18px;
				font-weight: 600;
				cursor: pointer;
				transition: transform 0.2s;
			}
			button:hover {
				transform: translate(-2px);
			}
			button:active {
				transform: translateY(0);
			}
			#result {
				margin-top: 25px;
				padding: 20px;
				background: #f8f9fa;
				border-radius: 10px;
				border-left: 4px solid #667eea;
			}
			.result-text {
				font-size: 18px;
				color: #333;
				line-height: 1.6;
			}
			.htmx-indicator {
				display: inline-block;
				width:20px;
				height: 20px;
				border: 3px solid #f3f3f3;
				border-top: 3px solid #667eea;
				border-radius: 50%;
				animation: spin 1s linear infinite;
				margin-left: 10px;
			}
			@keyframes spin {
				0% { transform: rotate(0deg) }
				100% { transform: rotate(360deg); }
			}


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
				<div class="form-group">
					<label for="value">Value:</label>
					<input type="number" id="value" name="value" step="any" required placeholder="Enter a number">
				</div>
				<div class="form-group">
					<label for="from">From:</label>
					<select name="from" id="from" required>
						<!-- Length units -->
						<optgroup label="Length">
							<option value="meters">Meters</option>
							<option value="kilometers">Kilometers</option>
							<option value="feet">Feet</option>
							<option value="miles">Miles</option>
							<option value="inches">Inches</option>
						</optgroup>
						<!-- Weight units -->
						<optgroup label="Weight">
							<option value="kilograms">Kilograms</option>
							<option value="pounds">Pounds</option>
							<option value="grams">Grams</option>
							<option value="ounces">Ounces</option>
						</optgroup>
						<!-- Temperature units -->
						<optgroup label="Temperature">
							<option value="celsius">Celsius</option>
							<option value="fahrenheit">Fahrenheit</option>
							<option value="kelvin">kelvin</option>
						</optgroup>
						<!-- Volume units -->
						<optgroup label="Volume">
							<option value="liters">Liters</option>
							<option value="gallons">Gallons</option>
							<option value="cups">Cups</option>
						</optgroup>
					</select>
				</div>
				<div class="form-group">
					<label for="to">To:</label>
					<select name="to" id="to" required>
						<!-- Length units -->
						<optgroup label="Length">
							<option value="meters">Meters</option>
							<option value="kilometers">Kilometers</option>
							<option value="feet">Feet</option>
							<option value="miles">Miles</option>
							<option value="inches">Inches</option>
						</optgroup>
						<!-- Weight units -->
						<optgroup label="Weight">
							<option value="kilograms">Kilograms</option>
							<option value="pounds">Pounds</option>
							<option value="grams">Grams</option>
							<option value="ounces">Ounces</option>
						</optgroup>
						<!-- Temperature units -->
						<optgroup label="Temperature">
							<option value="celsius">Celsius</option>
							<option value="fahrenheit">Fahrenheit</option>
							<option value="kelvin">kelvin</option>
						</optgroup>
						<!-- Volume units -->
						<optgroup label="Volume">
							<option value="liters">Liters</option>
							<option value="gallons">Gallons</option>
							<option value="cups">Cups</option>
						</optgroup>
					</select>
				</div>
				<button type="submit">
					Convert
					<span class="htmx-indicator"></span>
				</button>
			</form>
			<!-- This div will be replaced with the conversion result -->
			<div id="result"></div>
		</div>
	</body>
	</html>
	`
	// Parse and execute the template
	t, err := template.New("home").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

// convertHandler processes the conersion request
func convertHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Parse the form data from the request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get form values
	valueStr := r.FormValue("value")
	fromUnit := r.FormValue("from")
	toUnit := r.FormValue("to")
	category := r.FormValue("category")

	// Convert the string value to a float64 number
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	// Perform the conversion based on category
	var result float64
	switch category {
	case "length":
		result = convertLength(value, fromUnit, toUnit)
	case "weight":
		result = convertWeight(value, fromUnit, toUnit)
	case "temperature":
		result = convertTemperature(value, fromUnit, toUnit)
	case "volume":
		result = convertVolume(value, fromUnit, toUnit)
	default:
		http.Error(w, "Invalid category", http.StatusBadRequest)
		return
	}

	// Create result object
	convResult := ConversionResult{
		Value: result,
		FromUnit: fromUnit,
		ToUnit: toUnit,
		Category: category,
		InputValue: value,
	}

	// Template for the result HTML that HTMX will inject
	resultTmpl := `
	<div class="result-text">
		<p><strong>{{InputValue}} {{.FromUnit}}</strong> equals</p>
		<p class="result-number">{{printf "%.4f" .Value}} {{.ToUnit}}</p>
	</div>
	`
	// Parse and execute the result template
	t, err := template.New("result").Parse(resultTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, convResult)
}

// convertLength converts between lenght units
// First converst to meters (base units), then to target unit
func convertLength(value float64, from, to string) float64 {
	// Conversion factors to meters
	toMeters := map[string]float64{
		"meters": 1.0,
		"kilometers": 1000.0,
		"feet": 0.3048,
		"miles": 1609.34,
		"inches": 0.0254,
	}
	// Convert input to meters first
	meters := value * toMeters[from]
	// Convert from meters to target unit
	return meters / toMeters[to]
}

// convertWeight converts between weight units
func convertWeight(value float64, from, to string) float64 {
	// Conversion factors to kilograms
	toKilograms := map[string]float64{
		"kilograms": 1.0,
		"pounds": 0.453592,
		"grams": 0.001,
		"ounces": 0.0283495,
	}
	// Convert to kilograms first
	kilograms := value * toKilograms[from]

	// Convert from kilograms to target unit
	return kilograms / toKilograms[to]
}

// convertTemperature converts between temperature units
// Temperature requires special formulas, not just multiplication
func convertTemperature(value float64, from, to string) float64 {
	var celsius float64

	// First convert everything to Celsius
	switch from {
	case "celsius":
		celsius = value
	case "fahrenheit":
		celsius = (value - 32) * 5 / 9
	case "kelvin":
		celsius = value - 273.15
	}

	// Then convert from Celsius to target unit
	switch to {
	case "celsius":
		return celsius
	case "fahrenheit":
		return celsius*9/5 + 32
	case "kelvin":
		return celsius + 273.15
	}
	return celsius
}

// convertVolume converts between volume units
func convertVolume(value float64, from, to string) float64 {
	// Conversion factors to liters
	toLiters := map[string]float64{
		"liters": 1.0,
		"gallons": 3.78541,
		"milliliters": 0.001,
		"cups": 0.236588,
	}
	// Convert to liters first
	liters := value * toLiters[from]

	// Convert from liters to target unit
	return liters / toLiters[to]
}
