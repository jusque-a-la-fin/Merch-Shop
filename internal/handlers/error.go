package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Reason string `json:"reason"`
}

func RespondWithError(wrt http.ResponseWriter, err string, statusCode int) error {
	wrt.Header().Set("Content-Type", "application/json")
	wrt.WriteHeader(statusCode)
	errorResponse := ErrorResponse{Reason: err}
	errJSON := json.NewEncoder(wrt).Encode(errorResponse)
	return errJSON
}

func SendBadReq(wrt http.ResponseWriter, errStr string) error {
	err := fmt.Sprintf("Неверный запрос: %s", errStr)
	errResp := RespondWithError(wrt, err, http.StatusBadRequest)
	return errResp
}

func SendInternalServerError(wrt http.ResponseWriter, errStr string) error {
	err := fmt.Sprintf("Внутренняя ошибка сервера: %s", errStr)
	errResp := RespondWithError(wrt, err, http.StatusInternalServerError)
	return errResp
}

func SendUnauthorized(wrt http.ResponseWriter, errStr string) error {
	err := fmt.Sprintf("Неавторизован: %s", errStr)
	errResp := RespondWithError(wrt, err, http.StatusUnauthorized)
	return errResp
}
