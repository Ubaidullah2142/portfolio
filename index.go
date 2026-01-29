package handler

import (
	"fmt"
	"net/http"
)

// Handler is the main entrypoint for all API requests on Vercel.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Create a new router (multiplexer)
	mux := http.NewServeMux()

	// Define your API routes here
	mux.HandleFunc("/api/hello", helloHandler)
	mux.HandleFunc("/api/user", userHandler)

	// The router will find the correct handler for the request path
	mux.ServeHTTP(w, r)
}

// helloHandler is a sample handler for the "/api/hello" route.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the Go API!")
}

// userHandler is a sample handler for the "/api/user" route.
func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the user endpoint.")
}