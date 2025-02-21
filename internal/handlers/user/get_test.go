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

var getUrl = "/api/info"

// TestGetBadRequest тестирует некорректный запрос
func TestGetBadRequest(t *testing.T) {
	// некорректный метод запроса
	req, err := http.NewRequest(http.MethodPost, getUrl, nil)
	if err != nil {
		t.Fatal("Ошибка создания объекта *http.Request:", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uhr.GetInfo)
	handler.ServeHTTP(rr, req)
	expected := "Неверный запрос: wrong http method"
	HandleBadReq(t, rr, expected)
}

// TestGetUnauthorized тестирует случай, когда не удалось пройти аутентификацию
func TestGetUnauthorized(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, getUrl, nil)
	if err != nil {
		t.Fatal("Ошибка создания объекта *http.Request:", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uhr.GetInfo)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Ожидался код состояния ответа: %d, но получен: %d", http.StatusUnauthorized, rr.Code)
	}
}

// TestGetCoinsOK тестирует успешный ответ
func TestGetCoinsOK(t *testing.T) {
	arq := uhd.AuthRequest{Username: "user4", Password: "password4"}
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

	req = httptest.NewRequest(http.MethodGet, getUrl, nil)
	req.Header.Set("Authorization", authResp.Token)
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(uhr.GetInfo)
	handler.ServeHTTP(rr, req)
	CheckCodeAndMime(t, rr)
}
