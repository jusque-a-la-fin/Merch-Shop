package user

import (
	"encoding/json"
	"log"
	"merch-shop/internal/coins"
	"merch-shop/internal/handlers"
	"merch-shop/internal/inventory"
	"net/http"
)

func (hnd *UserHandler) GetInfo(wrt http.ResponseWriter, rqt *http.Request) {
	if rqt.Method != http.MethodGet {
		errSend := handlers.SendBadReq(wrt, "wrong http method")
		if errSend != nil {
			log.Printf("error while sending the bad request message: %v\n", errSend)
		}
		return
	}

	balance, userID := hnd.GetBalance(wrt, rqt)
	if balance == nil || userID == nil {
		return
	}

	items, err := hnd.InventoryRepo.Get(*userID)
	if err != nil {
		errSend := handlers.SendInternalServerError(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the internal server error message: %v\n", errSend)
		}
		return
	}

	history, err := hnd.CoinsRepo.GetHistory(*userID)
	if err != nil {
		errSend := handlers.SendInternalServerError(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the internal server error message: %v\n", errSend)
		}
		return
	}

	InfoResponse := struct {
		Coins       int              `json:"coins"`
		Inventory   []inventory.Item `json:"inventory"`
		CoinHistory coins.History    `json:"coinHistory"`
	}{
		Coins:       *balance,
		Inventory:   items,
		CoinHistory: *history,
	}

	wrt.Header().Set("Content-Type", "application/json")
	wrt.WriteHeader(http.StatusOK)
	errJSON := json.NewEncoder(wrt).Encode(InfoResponse)
	if errJSON != nil {
		log.Printf("error while sending response body: %v\n", errJSON)
	}
}
