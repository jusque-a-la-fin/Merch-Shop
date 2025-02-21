package user_test

import (
	"bytes"
	"encoding/json"
	uhd "merch-shop/internal/handlers/user"
	"merch-shop/internal/middleware"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var buyUrl = "/api/buy/book"

// TestBuyBadRequest тестирует некорректный запрос
func TestBuyBadRequest(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, buyUrl, nil)
	if err != nil {
		t.Fatal("Ошибка создания объекта *http.Request:", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uhr.BuyAnItem)
	handler.ServeHTTP(rr, req)
	expected := "Неверный запрос: wrong http method"
	HandleBadReq(t, rr, expected)

	incorrectInputs := []struct {
		Url    string
		ErrStr string
	}{
		{Url: "/api/buy/item", ErrStr: "Неверный запрос: the item has not been found"},
	}

	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/buy/{item}", uhr.BuyAnItem).Methods("GET")
	for _, inp := range incorrectInputs {
		req, err = http.NewRequest(http.MethodGet, inp.Url, nil)
		if err != nil {
			t.Fatal("Ошибка создания объекта *http.Request:", err)
		}

		rr := httptest.NewRecorder()
		rtr.ServeHTTP(rr, req)
		HandleBadReq(t, rr, inp.ErrStr)
	}
}

// TestBuyUnauthorized тестирует случай, когда не удалось пройти аутентификацию
func TestBuyUnauthorized(t *testing.T) {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/buy/{item}", uhr.BuyAnItem).Methods("GET")

	req, err := http.NewRequest(http.MethodGet, buyUrl, nil)
	if err != nil {
		t.Fatal("Ошибка создания объекта *http.Request:", err)
	}

	rr := httptest.NewRecorder()
	rtr.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Ожидался код состояния ответа: %d, но получен: %d", http.StatusUnauthorized, rr.Code)
	}
}

// TestBuyOK тестирует успешный ответ
func TestBuyOK(t *testing.T) {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/auth", uhr.GetAuthenticated).Methods("POST")
	rtr.HandleFunc("/api/buy/{item}", middleware.RequireAuth(uhr.BuyAnItem, dtb)).Methods("GET")

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
	rtr.ServeHTTP(rr, req)
	CheckCodeAndMime(t, rr)

	var authResp uhd.AuthResponse
	if err := json.NewDecoder(rr.Body).Decode(&authResp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	parts := strings.Split(authResp.Token, ".")
	if len(parts) != 3 {
		t.Fatalf("Ошибка: строка не является JWT-токеном: %v", err)
	}

	req = httptest.NewRequest(http.MethodGet, buyUrl, nil)
	req.Header.Set("Authorization", authResp.Token)
	rr = httptest.NewRecorder()
	rtr.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Ожидался код состояния ответа: %d, но получен: %d", http.StatusOK, rr.Code)
	}
}
