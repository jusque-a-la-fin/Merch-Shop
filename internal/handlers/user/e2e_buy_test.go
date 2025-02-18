package user_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// E2E-тест на сценарий покупки мерча
func TestBuyAnItem(t *testing.T) {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/buy/{item}", uhr.BuyAnItem).Methods("GET")

	rtr.Use(setValueMiddleware)
	ts := httptest.NewServer(rtr)
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL+"/api/buy/book", nil)
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
}
