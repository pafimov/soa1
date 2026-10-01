package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
	}

	s := &http.Server{
		Addr:              ":" + p,
		Handler:           handler(),
		ReadHeaderTimeout: 3 * time.Second,
	}

	log.Printf("catalog service started on :%s", p)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("GET /health", health)
	return m
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
