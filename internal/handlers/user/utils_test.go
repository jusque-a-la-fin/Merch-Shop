package user_test

import (
	"encoding/json"
	"log"
	"merch-shop/internal/coins"
	"merch-shop/internal/datastore"
	"merch-shop/internal/handlers"
	uhd "merch-shop/internal/handlers/user"
	"merch-shop/internal/inventory"
	"merch-shop/internal/session"
	"merch-shop/internal/user"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func GetUserHandler() *uhd.UserHandler {
	dtb, err := datastore.CreateNewDB()
	if err != nil {
		log.Fatalf("ошибка подключения к базе данных: %v", err)
	}

	var mutex sync.Mutex
	usr := user.NewDBRepo(dtb, &mutex)
	smg := session.NewSessionsManager()
	coins := coins.NewDBRepo(dtb, &mutex)
	inv := inventory.NewDBRepo(dtb, &mutex)
	userHandler := &uhd.UserHandler{
		UserRepo:      usr,
		Sessions:      smg,
		CoinsRepo:     coins,
		InventoryRepo: inv,
	}

	return userHandler
}

func HandleBadReq(t *testing.T, rr *httptest.ResponseRecorder, expected string) {
	code := rr.Code
	if code != http.StatusBadRequest {
		t.Errorf("Ожидался код состояния ответа: %d, но получен: %d", http.StatusBadRequest, code)
	}

	HandleError(t, rr, expected)
}

func HandleError(t *testing.T, rr *httptest.ResponseRecorder, expected string) {
	if mime := rr.Header().Get("Content-Type"); mime != "application/json" {
		t.Errorf("Заголовок Content-Type должен иметь MIME-тип application/json, но имеет %s", mime)
	}

	var errResp handlers.ErrorResponse
	err := json.Unmarshal(rr.Body.Bytes(), &errResp)
	if err != nil {
		t.Fatalf("Ошибка десериализации тела ответа сервера: %v", err)
	}

	result := errResp.Reason
	if result != expected {
		t.Errorf("Ожидалось %s, но получено %s", expected, result)
	}
}

func CheckCodeAndMime(t *testing.T, rr *httptest.ResponseRecorder) {
	if rr.Code != http.StatusOK {
		t.Errorf("Ожидался код состояния ответа: %d, но получен: %d", http.StatusOK, rr.Code)
	}

	if mime := rr.Header().Get("Content-Type"); mime != "application/json" {
		t.Errorf("Заголовок Content-Type должен иметь MIME-тип application/json, но имеет %s", mime)
	}
}
