package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var version = "dev"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/version", handleVersion)

	fmt.Printf("Server starting on port %s (version: %s)\n", port, version)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from go-release-action! Version: %s\n", version)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"version": "%s"}`, version)
}