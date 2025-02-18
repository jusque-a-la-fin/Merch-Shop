package user_test

import (
	"context"
	"merch-shop/internal/session"
	"net/http"
	"net/http/httptest"
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
	sess := &session.Session{ID: "1", UserName: "user1"}
	ctx := context.WithValue(context.Background(), session.SessionKey, sess)

	req := httptest.NewRequest(http.MethodGet, getUrl, nil).WithContext(ctx)
	rr := httptest.NewRecorder()

	uhr.GetInfo(rr, req)
	CheckCodeAndMime(t, rr)
}
