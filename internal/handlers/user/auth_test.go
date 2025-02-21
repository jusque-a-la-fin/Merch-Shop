package user_test

import (
	"bytes"
	"encoding/json"
	uhd "merch-shop/internal/handlers/user"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var authUrl = "/api/auth"
var uhr, dtb = GetUserHandler()
var testsBadRequest = []uhd.AuthRequest{
	{},
	{Username: "username"},
	{Password: "password"},
}

// TestBadRequest тестирует некорректный запрос
func TestBadRequest(t *testing.T) {
	// некорректный метод запроса
	req, err := http.NewRequest(http.MethodGet, authUrl, nil)
	if err != nil {
		t.Fatal("Ошибка создания объекта *http.Request:", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uhr.GetAuthenticated)
	handler.ServeHTTP(rr, req)
	expected := "Неверный запрос: wrong http method"
	HandleBadReq(t, rr, expected)

	expected = "Неверный запрос: empty fields of request body"
	// некорректные параметры тела запроса
	for _, testBR := range testsBadRequest {
		data, err := json.Marshal(testBR)
		if err != nil {
			t.Fatalf("Ошибка сериализации тела запроса клиента: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, authUrl, bytes.NewBuffer(data))
		if err != nil {
			t.Fatal("Ошибка создания объекта *http.Request:", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(uhr.GetAuthenticated)
		handler.ServeHTTP(rr, req)
		HandleBadReq(t, rr, expected)
	}

	req, err = http.NewRequest(http.MethodPost, authUrl, bytes.NewBuffer([]byte("Hello, Wotld!")))
	if err != nil {
		t.Fatal("Ошибка создания объекта *http.Request:", err)
	}

	rr = httptest.NewRecorder()
	expected = "Неверный запрос: wrong request body"
	handler.ServeHTTP(rr, req)
	HandleBadReq(t, rr, expected)
}

// TestUnauthorized тестирует случай, когда не удалось пройти аутентификацию
func TestUnauthorized(t *testing.T) {
	arq := uhd.AuthRequest{Username: "user1", Password: "password2"}
	data, err := json.Marshal(arq)
	if err != nil {
		t.Fatalf("Ошибка сериализации тела запроса клиента: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, authUrl, bytes.NewBuffer(data))
	if err != nil {
		t.Fatal("Ошибка создания объекта *http.Request:", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uhr.GetAuthenticated)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Ожидался код состояния ответа: %d, но получен: %d", http.StatusUnauthorized, rr.Code)
	}
}

// TestOK тестирует успешный ответ
func TestOK(t *testing.T) {
	arq := uhd.AuthRequest{Username: "user3", Password: "password3"}
	data, err := json.Marshal(arq)
	if err != nil {
		t.Fatalf("Ошибка сериализации тела запроса клиента: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, authUrl, bytes.NewBuffer(data))
	if err != nil {
		t.Fatal("Ошибка создания объекта *http.Request:", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uhr.GetAuthenticated)
	handler.ServeHTTP(rr, req)
	CheckCodeAndMime(t, rr)

	var authResp uhd.AuthResponse
	if err := json.NewDecoder(rr.Body).Decode(&authResp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	parts := strings.Split(authResp.Token, ".")
	if len(parts) != 3 {
		t.Fatalf("Ошибка: строка не является JWT-токеном: %v", err)
	}
}
