package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func sendRequest(method, url string) ([]byte, error) {
	var response *http.Response
	var err error

	// Создаем новый запрос с контекстом
	req, err := http.NewRequestWithContext(context.Background(), method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s request: %w", method, err)
	}

	// Отправляем запрос через DefaultClient
	response, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send %s request: %w", method, err)
	}
	defer response.Body.Close()

	// Чтение тела ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s response body: %w", method, err)
	}

	return body, nil
}

func main() {
	serverURL := flag.String("url", "http://localhost:8080", "URL of the server")
	resourcePath := flag.String("path", "/", "Path of the resource")
	flag.Parse()

	url := *serverURL + *resourcePath

	// Send GET request
	fmt.Println("Sending GET request...")
	getBody, err := sendRequest(http.MethodGet, url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GET Response: %s\n", string(getBody))

	// Send POST request
	fmt.Println("Sending POST request...")
	postBody, err := sendRequest(http.MethodPost, url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("POST Response: %s\n", string(postBody))
}
