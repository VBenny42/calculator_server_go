package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func main() {
	port := ":3000"

	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger := slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}))
	slog.SetDefault(logger)

	slog.Info("Starting server", "port", port)

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		var pair numberPair
		genericHandler(w, r, pair, func(input numberPair) int {
			return *input.Number1 + *input.Number2
		})
	})

	http.HandleFunc("/subtract", func(w http.ResponseWriter, r *http.Request) {
		var pair numberPair
		genericHandler(w, r, pair, func(input numberPair) int {
			return *input.Number1 - *input.Number2
		})
	})

	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		var pair numberPair
		genericHandler(w, r, pair, func(input numberPair) int {
			return *input.Number1 * *input.Number2
		})
	})

	http.HandleFunc("/divide", func(w http.ResponseWriter, r *http.Request) {
		var pair dividePair
		genericHandler(w, r, pair, func(input dividePair) int {
			return *input.Dividend / *input.Divisor
		})
	})

	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		var numbers sumArray
		genericHandler(w, r, numbers, func(input sumArray) int {
			sum := 0
			for _, n := range input.Items {
				sum += n
			}
			return sum
		})
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
