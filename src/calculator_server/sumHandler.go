package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type sumArray struct {
	Items []int `json:"items"`
}

func (s sumArray) validate() error {
	if s.Items == nil {
		return errors.New("items field is required but missing")
	}
	return nil
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Received request", "method", r.Method, "url", r.URL.String())

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var sum sumArray

	err := json.NewDecoder(r.Body).Decode(&sum)
	if err != nil {
		errMsg := "Failed to decode request"
		http.Error(w, errMsg, http.StatusBadRequest)
		slog.Error(errMsg, "error", err)
		return
	}

	err = sum.validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.Error("Invalid request", "error", err)
		return
	}

	result := 0
	for _, item := range sum.Items {
		result += item
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]int{"result": result})
	if err != nil {
		errMsg := "Failed to encode response"
		slog.Error(errMsg, "error", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	slog.Info("Request processed", "result", result)
}
