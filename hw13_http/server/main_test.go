package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	// Создаем мок-сервер
	handler := http.NewServeMux()
	handler.HandleFunc("/", handleRequest)
	server := httptest.NewServer(handler)
	defer server.Close()

	tests := []struct {
		method       string
		expectedBody string
		expectedCode int
	}{
		{http.MethodGet, "This is the GET Response\n", http.StatusOK},
		{http.MethodPost, "This is the POST Response\n", http.StatusOK},
		{http.MethodPut, "Method not allowed\n", http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			// Создаем запрос с контекстом
			req, err := http.NewRequestWithContext(context.Background(), tt.method, server.URL, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Отправляем запрос
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()

			// Проверка статуса
			if resp.StatusCode != tt.expectedCode {
				t.Errorf("Expected status %d, got %d", tt.expectedCode, resp.StatusCode)
			}

			// Проверка тела ответа
			body := make([]byte, len(tt.expectedBody))
			resp.Body.Read(body)
			if string(body) != tt.expectedBody {
				t.Errorf("Expected body %s, got %s", tt.expectedBody, string(body))
			}
		})
	}
}
