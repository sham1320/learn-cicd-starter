package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// read PORT from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// only the /ready endpoint
	mux.HandleFunc("GET /v1/ready", handlerReadiness)

	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(":"+port, mux)
}

// handlerReadiness responds with a simple JSON status
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}


