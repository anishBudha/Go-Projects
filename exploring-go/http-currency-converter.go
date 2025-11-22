package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", uiHandler)
	http.HandleFunc("/convert", convertHandler)
	http.ListenAndServe(":8080", nil)
}

func uiHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<body>
		<h3>USD to CAD converter</h3>
		<input id="amt" type="number" placeholder="Enter USD"/>
		<button onclick="convert()">Convert</button>

		<p id="result"></p>

		<script>
			async function convert() {
				const amount = document.getElementById("amt").value;
				const res = await fetch("/convert?amount=" + amount);
				const data = await res.json();
				document.getElementById("result").innerText = 
					"CAD: " + data.cad + "(rate " + data.rate + ")";
			}
		</script>
	</body>
	</html>`
	w.Write([]byte(html))
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	amountStr := r.URL.Query().Get("amount")
	amount, _ := strconv.ParseFloat(amountStr, 64)

	rate := 1.37

	out := map[string]any {
		"usd": amount,
		"cad": amount * rate,
		"rate": rate,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}
