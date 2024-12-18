package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendRequest(t *testing.T) {
	// Мок-сервер для GET-запроса
	getHandler := http.NewServeMux()
	getHandler.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("GET response"))
	})

	// Мок-сервер для POST-запроса
	postHandler := http.NewServeMux()
	postHandler.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("POST response"))
	})

	// Создаем тестовые серверы
	getServer := httptest.NewServer(getHandler)
	defer getServer.Close()

	postServer := httptest.NewServer(postHandler)
	defer postServer.Close()

	// Тест для GET-запроса
	t.Run("GET Request", func(t *testing.T) {
		body, err := sendRequest(http.MethodGet, getServer.URL)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		expected := "GET response"
		if string(body) != expected {
			t.Errorf("Expected %s, got %s", expected, string(body))
		}
	})

	// Тест для POST-запроса
	t.Run("POST Request", func(t *testing.T) {
		body, err := sendRequest(http.MethodPost, postServer.URL)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		expected := "POST response"
		if string(body) != expected {
			t.Errorf("Expected %s, got %s", expected, string(body))
		}
	})
}
