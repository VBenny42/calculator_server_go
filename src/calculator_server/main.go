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
		genericPairHandler(w, r, func(a, b int) int { return a + b })
	})

	http.HandleFunc("/subtract", func(w http.ResponseWriter, r *http.Request) {
		genericPairHandler(w, r, func(a, b int) int { return a - b })
	})

	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		genericPairHandler(w, r, func(a, b int) int { return a * b })
	})

	http.HandleFunc("/divide", divideHandler)

	http.HandleFunc("/sum", sumHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
