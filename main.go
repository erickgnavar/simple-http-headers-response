package main

import (
	"fmt"
	"net/http"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request with headers:")
	// Print each header key-value pair
	for key, value := range r.Header {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Set response headers
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "HTTP Headers\n\n")
	// Send response
	for key, value := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", key, value)
	}
}

func main() {
	// Register handler for the /headers route
	http.HandleFunc("/", headersHandler)
	// Start the server on port 8080
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
