package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HealthResponse{
		Status: "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health/", healthHandler)

	log.Printf("Server starting on port 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
} 