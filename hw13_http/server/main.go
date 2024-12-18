package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	ip := flag.String("ip", "localhost", "IP address")
	port := flag.String("port", "8080", "Port number")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *ip, *port)
	log.Printf("Server is running on %s\n", address)

	http.HandleFunc("/", handleRequest)

	server := &http.Server{
		Addr:         address,
		Handler:      nil,
		ReadTimeout:  10 * time.Second, // время ожидания на чтение запроса
		WriteTimeout: 10 * time.Second, // время ожидания на запись ответа
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start on %s: %v", address, err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request from %s", r.Method, r.RemoteAddr)

	switch r.Method {
	case http.MethodGet:
		respondWithText(w, "This is the GET Response")
	case http.MethodPost:
		respondWithText(w, "This is the POST Response")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func respondWithText(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintln(w, message); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
