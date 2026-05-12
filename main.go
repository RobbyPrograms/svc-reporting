package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// serviceName identifies this microservice in logs and health responses.
const serviceName = "svc-reporting"

func main() {
	// Register the health endpoint used by CI, Docker, and uptime checks.
	http.HandleFunc("/health", healthHandler)
	log.Printf("Starting %s on :8080", serviceName)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Return a tiny JSON payload so callers can confirm the service is alive.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": serviceName,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
