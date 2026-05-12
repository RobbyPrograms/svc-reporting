package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const serviceName = "svc-reporting"

func main() {
	http.HandleFunc("/health", healthHandler)
	log.Printf("Starting %s on :8080", serviceName)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": serviceName,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
