package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type numberPair struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

func (p numberPair) validate() error {
	if p.Number1 == nil {
		return errors.New("number1 field is required but missing")
	}
	if p.Number2 == nil {
		return errors.New("number2 field is required but missing")
	}
	return nil
}

func genericPairHandler(w http.ResponseWriter, r *http.Request, f func(int, int) int) {
	slog.Info("Received request", "method", r.Method, "url", r.URL.String())

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var pair numberPair

	err := json.NewDecoder(r.Body).Decode(&pair)
	if err != nil {
		slog.Error("Failed to decode request", "error", err)
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	err = pair.validate()
	if err != nil {
		slog.Error("Invalid request", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := f(*pair.Number1, *pair.Number2)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]int{"result": result})
	if err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

	slog.Info("Request processed", "result", result)
}
