package main

import (
  "fmt"
  "os"
  "flag"
  "net/http"
  "log/slog"
)

var (
  addr = flag.String("addr", ":8080", "http service address")
)

func main() {
  flag.Parse()
  logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome")
    logger.Info("Response: ", slog.String("request", "/"), slog.Int("status", http.StatusOK))
  })

  http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Healthy")
    logger.Info("Response: ", slog.String("request", "/health"), slog.Int("status", http.StatusOK))
  })

  logger.Error("Server stopped", slog.Any("error", http.ListenAndServe(*addr, nil)))
}
