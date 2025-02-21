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

// E2E-тест на сценарий получения информации
func TestGetInfo(t *testing.T) {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/auth", uhr.GetAuthenticated).Methods("POST")
	rtr.HandleFunc("/api/info", middleware.RequireAuth(uhr.GetInfo, dtb)).Methods("GET")

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

	ts := httptest.NewServer(rtr)
	defer ts.Close()

	req, err = http.NewRequest(http.MethodGet, ts.URL+"/api/info", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Authorization", authResp.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался код состояния ответа: %d, но получен: %d", http.StatusOK, resp.StatusCode)
	}
}
