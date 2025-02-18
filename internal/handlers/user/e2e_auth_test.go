package user_test

import (
	"bytes"
	"encoding/json"
	uhd "merch-shop/internal/handlers/user"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// E2E-тест на сценарий аутентификации
func TestAuth(t *testing.T) {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/auth", uhr.GetAuthenticated).Methods("POST")

	ts := httptest.NewServer(rtr)
	defer ts.Close()

	requestPayload := uhd.AuthRequest{Username: "user4", Password: "password4"}
	payload, err := json.Marshal(requestPayload)
	if err != nil {
		t.Fatalf("Failed to marshal request payload: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/auth", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался код состояния ответа: %d, но получен: %d", http.StatusOK, resp.StatusCode)
	}

	if mime := resp.Header.Get("Content-Type"); mime != "application/json" {
		t.Errorf("Заголовок Content-Type должен иметь MIME-тип application/json, но имеет %s", mime)
	}

	var authResp uhd.AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	parts := strings.Split(authResp.Token, ".")
	if len(parts) != 3 {
		t.Fatalf("Ошибка: строка не является JWT-токеном: %v", err)
	}
}
