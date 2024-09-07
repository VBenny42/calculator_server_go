package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func genericHandler[T genericNumberInput](w http.ResponseWriter, r *http.Request, input T, f func(T) int) {
	slog.Info("Received request", "method", r.Method, "url", r.URL.String())

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Error("Failed to decode request", "error", err)
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	err = input.validate()
	if err != nil {
		slog.Error("Invalid request", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := f(input)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]int{"result": result})
	if err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

	slog.Info("Request processed", "result", result)
}
