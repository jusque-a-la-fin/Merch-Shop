package user_test

import (
	"bytes"
	"encoding/json"
	uhd "merch-shop/internal/handlers/user"
	"merch-shop/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

var authUrl = "/api/auth"
var uhr = test.GetUserHandler()
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
	test.HandleBadReq(t, rr, expected)

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
		test.HandleBadReq(t, rr, expected)
	}

	req, err = http.NewRequest(http.MethodPost, authUrl, bytes.NewBuffer([]byte("Hello, Wotld!")))
	if err != nil {
		t.Fatal("Ошибка создания объекта *http.Request:", err)
	}

	rr = httptest.NewRecorder()
	expected = "Неверный запрос: wrong request body"
	handler.ServeHTTP(rr, req)
	test.HandleBadReq(t, rr, expected)
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
	test.CheckCodeAndMime(t, rr)
}
