package main

import (
	"encoding/json"
	"html/template"
	"net/http"	
	"strconv"
)

var gravity = map[string]float64{
	"earth": 1.0,
	"mars": 0.38,
	"jupiter": 2.34,
	"moon": 0.16,
}

func main() {
	http.HandleFunc("/", uiHandler)
	http.HandleFunc("/weight", weightHandler)
	http.ListenAndServe(":8080", nil)
}

// Serve UI

func uiHandler(w http.ResponseWriter, r *http.Request) {
	page := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Planet Weight Calculator</title>
	<head>
	<body style="font-family: poppins; max-width: 400px; margin 40px auto">
		<h2>Planet Weight Calculator</h2>
		<label>Your weight on Earth(kg):</label><br>
		<input id="weight" type="number" placeholder="e.g. 70">
		<br></br>

		<label>Select a planet:</label>
		<select id="planet">
			<option value="earth">Earth</option>
			<option value="mars">Mars</option>
			<option value="jupiter">Jupiter</option>
			<option value="moon">Moon</option>
		</select>

		<br></br>

		<button onclick="calc()">Calculate</button>

		<h3 id="result"></h3>

		<script>
			async function calc() {
				const w =  document.getElementById("weight").value;
				const p = document.getElementById("planet").value;

				const res = await fetch("/weight?planet=" + p + "&weight=" + w);
				const data = await res.json();

				if (data.error) {
					document.getElementById("result").innerText = data.error;
					return;
				}

				document.getElementById("result").innerText = 
				"On " + data.planet + ", you would weigh " + data.planet_weight + " kg.";
			}
		</script>
	</body>
	</html>
	`
	t := template.Must(template.New("ui").Parse(page))
	t.Execute(w, nil)
}

// API endpoint

func weightHandler(w http.ResponseWriter, r *http.Request) {
	planet := r.URL.Query().Get("planet")
	weightStr := r.URL.Query().Get("weight")

	if planet == "" || weightStr == "" {
		writeJSON(w, map[string]any{"error": "missing planet or weight"})
		return
	}

	factor, ok := gravity[planet]
	if !ok {
		writeJSON(w, map[string]any{"error": "unkown planet"})
		return
	}

	weight, err := strconv.ParseFloat(weightStr, 64) 
	if err != nil {
		writeJSON(w, map[string]any{"error": "invalid output"})
		return
	}

	result := map[string]any {
		"planet": planet,
		"earth_weight": weight,
		"planet_weight": weight * factor,
		"gravity_factor": factor,
	}

	writeJSON(w, result)

}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
