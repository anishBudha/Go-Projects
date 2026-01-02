package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// ConversionResult holds the data we'll send back to the user
type ConversionResult struct {
	Value      float64 // The converted value
	FromUnit   string  // Original unit
	ToUnit     string  // Target unit
	Category   string  // Type of conversion (length, weight, etc.)
	InputValue float64 // The original value entered by user
}

// main is the entry point of our program
func main() {
	// Set up route handlers - these connect URLs to functions
	http.HandleFunc("/", homeHandler)           // Main page
	http.HandleFunc("/convert", convertHandler) // Handles conversion requests

	// Start the web server on port 8080
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// homeHandler serves the main HTML page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Define our HTML template as a string
	// The {{}} syntax is Go's template language for dynamic content
	tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Unit Converter</title>
    <!-- Include HTMX library from CDN -->
    <script src="https://unpkg.com/htmx.org@1.9.10"></script>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: monospace;
            background: white;
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 20px;
        }
        .container {
            background: white;
            padding: 30px;
            border: 2px solid black;
            max-width: 400px;
            width: 100%;
        }
        h1 {
            color: black;
            margin-bottom: 20px;
            font-size: 1.5em;
        }
        .form-group {
            margin-bottom: 15px;
        }
        label {
            display: block;
            margin-bottom: 5px;
            color: black;
        }
        input, select {
            width: 100%;
            padding: 8px;
            border: 1px solid black;
            font-size: 14px;
            font-family: monospace;
        }
        input:focus, select:focus {
            outline: 2px solid black;
        }
        button {
            width: 100%;
            padding: 10px;
            background: black;
            color: white;
            border: none;
            font-size: 16px;
            cursor: pointer;
            font-family: monospace;
        }
        button:hover {
            background: #333;
        }
        #result {
            margin-top: 20px;
            padding: 15px;
            background: white;
            border: 1px solid black;
        }
        .result-text {
            font-size: 14px;
            color: black;
            line-height: 1.6;
        }
        .result-number {
            font-size: 18px;
            font-weight: bold;
            color: black;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Unit Converter</h1>
        
        <!-- 
            HTMX attributes explained:
            hx-post="/convert" - sends a POST request to /convert endpoint
            hx-target="#result" - puts the response into the element with id="result"
            hx-swap="innerHTML" - replaces the inner HTML of the target
        -->
        <form id="converterForm">
            
            <div class="form-group">
                <label for="category">Category:</label>
                <select name="category" id="category" required onchange="updateUnits()">
                    <option value="">Select category</option>
                    <option value="length">Length</option>
                    <option value="weight">Weight</option>
                    <option value="temperature">Temperature</option>
                    <option value="volume">Volume</option>
                </select>
            </div>

            <div class="form-group">
                <label for="value">Value:</label>
                <input type="number" id="value" name="value" step="any" required>
            </div>

            <div class="form-group">
                <label for="from">From:</label>
                <select name="from" id="from" required>
                    <option value="">Select category first</option>
                </select>
            </div>

            <div class="form-group">
                <label for="to">To:</label>
                <select name="to" id="to" required>
                    <option value="">Select category first</option>
                </select>
            </div>

            <button type="submit">Convert</button>
        </form>

        <!-- This div will be replaced with the conversion result -->
        <div id="result"></div>
    </div>

    <script>
        // Define units for each category
        const units = {
            length: ['meters', 'kilometers', 'feet', 'miles', 'inches'],
            weight: ['kilograms', 'pounds', 'grams', 'ounces'],
            temperature: ['celsius', 'fahrenheit', 'kelvin'],
            volume: ['liters', 'gallons', 'milliliters', 'cups']
        };

        // Function to update the From and To dropdowns based on selected category
        function updateUnits() {
            const category = document.getElementById('category').value;
            const fromSelect = document.getElementById('from');
            const toSelect = document.getElementById('to');
            
            // Clear existing options
            fromSelect.innerHTML = '';
            toSelect.innerHTML = '';
            
            if (category && units[category]) {
                // Add new options based on category
                units[category].forEach(unit => {
                    // Create option for "from" dropdown
                    const fromOption = document.createElement('option');
                    fromOption.value = unit;
                    fromOption.textContent = unit.charAt(0).toUpperCase() + unit.slice(1);
                    fromSelect.appendChild(fromOption);
                    
                    // Create option for "to" dropdown
                    const toOption = document.createElement('option');
                    toOption.value = unit;
                    toOption.textContent = unit.charAt(0).toUpperCase() + unit.slice(1);
                    toSelect.appendChild(toOption);
                });
            } else {
                // If no category selected, show placeholder
                fromSelect.innerHTML = '<option value="">Select category first</option>';
                toSelect.innerHTML = '<option value="">Select category first</option>';
            }
        }

        // Handle form submission
        document.getElementById('converterForm').addEventListener('submit', function(e) {
            e.preventDefault(); // Prevent default form submission
            
            // Get form values directly from inputs
            const category = document.getElementById('category').value;
            const value = document.getElementById('value').value;
            const fromUnit = document.getElementById('from').value;
            const toUnit = document.getElementById('to').value;
            
            // Create URL-encoded form data
            const formData = new URLSearchParams();
            formData.append('category', category);
            formData.append('value', value);
            formData.append('from', fromUnit);
            formData.append('to', toUnit);
            
            // Send POST request using fetch
            fetch('/convert', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: formData.toString()
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.text();
            })
            .then(html => {
                // Update the result div with the response
                document.getElementById('result').innerHTML = html;
            })
            .catch(error => {
                console.error('Error:', error);
                document.getElementById('result').innerHTML = '<p>Error converting units</p>';
            });
        });
    </script>
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

// convertHandler processes the conversion request
func convertHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data from the request
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get form values
	valueStr := r.FormValue("value")
	fromUnit := r.FormValue("from")
	toUnit := r.FormValue("to")
	category := r.FormValue("category")

	// Debug logging to see what we received
	log.Printf("Received - Value: %s, From: %s, To: %s, Category: %s", valueStr, fromUnit, toUnit, category)

	// Validate that we got all required fields
	if valueStr == "" || fromUnit == "" || toUnit == "" || category == "" {
		log.Printf("Missing fields - Value: '%s', From: '%s', To: '%s', Category: '%s'", valueStr, fromUnit, toUnit, category)
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Convert the string value to a float64 number
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Printf("Error parsing value: %v", err)
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
		Value:      result,
		FromUnit:   fromUnit,
		ToUnit:     toUnit,
		Category:   category,
		InputValue: value,
	}

	// Template for the result HTML that HTMX will inject
	resultTmpl := `
<div class="result-text">
    <p><strong>{{.InputValue}} {{.FromUnit}}</strong> equals</p>
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

// convertLength converts between length units
// First converts to meters (base unit), then to target unit
func convertLength(value float64, from, to string) float64 {
	// Conversion factors to meters
	toMeters := map[string]float64{
		"meters":     1.0,
		"kilometers": 1000.0,
		"feet":       0.3048,
		"miles":      1609.34,
		"inches":     0.0254,
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
		"pounds":    0.453592,
		"grams":     0.001,
		"ounces":    0.0283495,
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
		"liters":      1.0,
		"gallons":     3.78541,
		"milliliters": 0.001,
		"cups":        0.236588,
	}

	// Convert to liters first
	liters := value * toLiters[from]
	
	// Convert from liters to target unit
	return liters / toLiters[to]
}
