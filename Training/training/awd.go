package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const usdt = 1.0
const eurt = 0.86
const uaht = 41.78

func main() {
	http.HandleFunc("/convert", convertHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func convertHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	amountStr := r.URL.Query().Get("amount")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amount, _ := strconv.ParseFloat(amountStr, 64)
	converted := 0.0
	a := 0.0
	if from == "UAH" {
		a = uaht
	} else if from == "EUR" {
		a = eurt
	} else {
		a = usdt
	}
	b := 0.0
	if to == "EUR" {
		b = eurt
	} else if to == "USD" {
		b = usdt
	} else {
		b = uaht
	}
	converted = amount * (b / a)
	if to == from {
		converted = 0.0
	}
	if converted == 0.0 {
		_, err := fmt.Fprint(w, "Error")
		if err != nil {
			return
		}
	}
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"from":      from,
		"to":        to,
		"amount":    amount,
		"converted": converted,
	})
	if err != nil {
		return
	}

	fmt.Println("Converted ", converted)
	fmt.Println("amount:", amount, "from:", from, "to:", to, "amountStr:", amountStr)

}
