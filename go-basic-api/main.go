package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// VersionResponse represents the JSON structure for our response.
type VersionResponse struct {
	Version string `json:"version"`
}

func main() {
	// Set the log output to standard output (default is standard error)
	log.SetOutput(os.Stdout)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s", r.RemoteAddr)
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
			log.Printf("Method not allowed from %s", r.RemoteAddr)
			return
		}

		version := os.Getenv("API_VERSION")
		if version == "" {
			version = "v0"  // Default value if not set
		}

		resp := VersionResponse{
			Version: version,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

		log.Printf("Responded to %s with version: %s", r.RemoteAddr, version)
	})

	log.Printf("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
