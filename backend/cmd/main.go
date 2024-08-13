package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/assaidy/unit-converter/backend/converter"
	"github.com/gorilla/mux"
)

type Result struct {
	Result float64 `json:"result"`
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

func handleConversion(w http.ResponseWriter, r *http.Request) error {
	quereyParams := r.URL.Query()

	section := quereyParams.Get("section")
	fromUnit := quereyParams.Get("from")
	toUnit := quereyParams.Get("to")
	amount := quereyParams.Get("amount")

	amount = strings.TrimSpace(amount)
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("'%s' is not a valid number.", amount)
	}

	log.Printf("converting from: %s to: %s amount: %f section: %s", fromUnit, toUnit, amountFloat, section)
	result := converter.Convert(section, fromUnit, toUnit, amountFloat)
	log.Println("result: ", result)

	writeJSON(w, Result{Result: result})

	return nil
}

func makeHttpHandlerFunc(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println(err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
	}
}

func writeJSON(w http.ResponseWriter, v any) error {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return fmt.Errorf("Failed to encode json response: %v", err)
	}
	return nil
}

func addCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// If it's a preflight request (OPTIONS), respond with 200 OK
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/convert", makeHttpHandlerFunc(handleConversion)).Methods("GET")

	// Apply the CORS middleware
	corsHandler := addCORS(r)

	log.Println("Starting server on :6868...")
	log.Fatal(http.ListenAndServe(":6868", corsHandler))
}
