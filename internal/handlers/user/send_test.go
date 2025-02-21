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

var sendUrl = "/api/sendCoin"

var testsSendCoinsBadRequest = []uhd.SendCoinRequest{
	{},
	{ToUser: "username"},
	{Amount: 5},
}

// TestSendCoinsBadRequest тестирует некорректный запрос
func TestSendCoinsBadRequest(t *testing.T) {
	// некорректный метод запроса
	req, err := http.NewRequest(http.MethodGet, sendUrl, nil)
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
	for _, testBR := range testsSendCoinsBadRequest {
		data, err := json.Marshal(testBR)
		if err != nil {
			t.Fatalf("Ошибка сериализации тела запроса клиента: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, sendUrl, bytes.NewBuffer(data))
		if err != nil {
			t.Fatal("Ошибка создания объекта *http.Request:", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(uhr.GetAuthenticated)
		handler.ServeHTTP(rr, req)
		HandleBadReq(t, rr, expected)
	}

	req, err = http.NewRequest(http.MethodPost, sendUrl, bytes.NewBuffer([]byte("Hello, Wotld!")))
	if err != nil {
		t.Fatal("Ошибка создания объекта *http.Request:", err)
	}

	rr = httptest.NewRecorder()
	expected = "Неверный запрос: wrong request body"
	handler.ServeHTTP(rr, req)
	HandleBadReq(t, rr, expected)
}

// TestSendCoinsUnauthorized тестирует случай, когда не удалось пройти аутентификацию
func TestSendCoinsUnauthorized(t *testing.T) {
	srq := uhd.SendCoinRequest{ToUser: "user1", Amount: 5}
	data, err := json.Marshal(srq)
	if err != nil {
		t.Fatalf("Ошибка сериализации тела запроса клиента: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, sendUrl, bytes.NewBuffer(data))
	if err != nil {
		t.Fatal("Ошибка создания объекта *http.Request:", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uhr.SendCoins)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Ожидался код состояния ответа: %d, но получен: %d", http.StatusUnauthorized, rr.Code)
	}
}

// TestSendCoinsOK тестирует успешный ответ
func TestSendCoinsOK(t *testing.T) {
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

	srq := uhd.SendCoinRequest{ToUser: "user2", Amount: 15}
	data, err = json.Marshal(srq)
	if err != nil {
		t.Fatalf("Ошибка сериализации тела запроса клиента: %v", err)
	}

	req = httptest.NewRequest(http.MethodPost, sendUrl, bytes.NewBuffer(data))
	req.Header.Set("Authorization", authResp.Token)
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(uhr.SendCoins)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Ожидался код состояния ответа: %d, но получен: %d", http.StatusOK, rr.Code)
	}
}
