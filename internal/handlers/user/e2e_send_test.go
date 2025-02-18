package user_test

import (
	"bytes"
	"context"
	"encoding/json"
	uhd "merch-shop/internal/handlers/user"
	"merch-shop/internal/session"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func setValueMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := &session.Session{ID: "1", UserName: "user1"}
		ctx := context.WithValue(r.Context(), session.SessionKey, sess)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// E2E-тест на сценарий передачи монеток другим сотрудникам
func TestSendCoins(t *testing.T) {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/sendCoin", uhr.SendCoins).Methods("POST")

	rtr.Use(setValueMiddleware)
	ts := httptest.NewServer(rtr)
	defer ts.Close()

	requestPayload := uhd.SendCoinRequest{ToUser: "user2", Amount: 15}
	payload, err := json.Marshal(requestPayload)
	if err != nil {
		t.Fatalf("Failed to marshal request payload: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, ts.URL+"/api/sendCoin", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

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
