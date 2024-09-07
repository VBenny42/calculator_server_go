package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"reflect"
)

type genericNumberInput interface {
	validate() error
}

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

func genericHandler(w http.ResponseWriter, r *http.Request, input genericNumberInput, f func(genericNumberInput) int) {
	slog.Info("Received request", "method", r.Method, "url", r.URL.String())

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	inputType := reflect.TypeOf(input)
	newInput := reflect.New(inputType).Interface()

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Error("Failed to decode request", "error", err)
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	typedInput, ok := newInput.(genericNumberInput)
	if !ok {
		slog.Error("Failed to cast input", "error", errors.New("Failed to cast input"))
		http.Error(w, "Failed to cast input", http.StatusInternalServerError)
		return
	}

	err = typedInput.validate()
	if err != nil {
		slog.Error("Invalid request", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := f(typedInput)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]int{"result": result})
	if err != nil {
		slog.Error("Failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

	slog.Info("Request processed", "result", result)
}
