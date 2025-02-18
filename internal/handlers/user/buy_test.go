package user_test

import (
	"context"
	"merch-shop/internal/session"
	"merch-shop/test"
	"net/http"
	"net/http/httptest"
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
	test.HandleBadReq(t, rr, expected)

	incorrectInputs := []struct {
		Url    string
		ErrStr string
	}{
		{Url: "/api/buy/item", ErrStr: "Неверный запрос: an item has not been found"},
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
		test.HandleBadReq(t, rr, inp.ErrStr)
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
	rtr.HandleFunc("/api/buy/{item}", uhr.BuyAnItem).Methods("GET")

	sess := &session.Session{ID: "1", UserName: "user1"}
	ctx := context.WithValue(context.Background(), session.SessionKey, sess)

	req := httptest.NewRequest(http.MethodGet, buyUrl, nil).WithContext(ctx)
	rr := httptest.NewRecorder()
	rtr.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Ожидался код состояния ответа: %d, но получен: %d", http.StatusOK, rr.Code)
	}
}
