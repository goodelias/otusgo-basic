package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ip := flag.String("ip", "localhost", "IP address")
	port := flag.String("port", "8080", "Port number")
	flag.Parse()
	address := *ip + ":" + *port

	log.Printf("server is running on %s\n", address)

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Server failed to start on %s: %v", address, err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request from %s", r.Method, r.RemoteAddr)

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprintln(w, "This is the GET Response"); err != nil {
			log.Printf("Error writing GET response: %v", err)
		}
	case http.MethodPost:
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprintln(w, "This is the POST Response"); err != nil {
			log.Printf("Error writing POST response: %v", err)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
