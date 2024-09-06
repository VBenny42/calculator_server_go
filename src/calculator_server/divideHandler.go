package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type dividePair struct {
	Dividend *int `json:"dividend"`
	Divisor  *int `json:"divisor"`
}

func (p dividePair) validate() error {
	if p.Dividend == nil {
		return errors.New("dividend field is required but missing")
	}
	if p.Divisor == nil {
		return errors.New("divisor field is required but missing")
	}
	if *p.Divisor == 0 {
		return errors.New("divisor cannot be zero")
	}
	return nil
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Received request", "method", r.Method, "url", r.URL.String())

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var pair dividePair

	err := json.NewDecoder(r.Body).Decode(&pair)
	if err != nil {
		errMsg := "Failed to decode request"
		http.Error(w, errMsg, http.StatusBadRequest)
		slog.Error(errMsg, "error", err)
		return
	}

	err = pair.validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.Error("Invalid request", "error", err)
		return
	}

	result := *pair.Dividend / *pair.Divisor

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		errMsg := "Failed to encode response"
		slog.Error(errMsg, "error", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	slog.Info("Request processed", "result", result)
}
