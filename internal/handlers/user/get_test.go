package user_test

import (
	"context"
	"merch-shop/internal/session"
	"merch-shop/test"
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
	test.HandleBadReq(t, rr, expected)
}

// TestGetCoinsOK тестирует успешный ответ
func TestGetCoinsOK(t *testing.T) {
	sess := &session.Session{ID: "1", UserName: "user1"}
	ctx := context.WithValue(context.Background(), session.SessionKey, sess)

	req := httptest.NewRequest(http.MethodGet, getUrl, nil).WithContext(ctx)
	rr := httptest.NewRecorder()

	uhr.GetInfo(rr, req)
	test.CheckCodeAndMime(t, rr)
}
