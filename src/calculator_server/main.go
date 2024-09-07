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

	var pair numberPair

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		genericHandler[numberPair](w, r, pair, func(*input genericNumberInput) int {
			return *numberPair.Number1 + *numberPair.Number2
		})
	})

	// http.HandleFunc("/subtract", func(w http.ResponseWriter, r *http.Request) {
	// 	genericHandler(w, r, pair, func(input genericNumberInput) int {
	// 		numberPair := input.(numberPair)
	// 		return *numberPair.Number1 - *numberPair.Number2
	// 	})
	// })
	//
	// http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
	// 	genericHandler(w, r, pair, func(input genericNumberInput) int {
	// 		numberPair := input.(numberPair)
	// 		return *numberPair.Number1 * *numberPair.Number2
	// 	})
	// })

	http.HandleFunc("/divide", divideHandler)

	http.HandleFunc("/sum", sumHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
