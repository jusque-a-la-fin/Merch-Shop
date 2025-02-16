package user

import (
	"encoding/json"
	"log"
	"merch-shop/internal/coins"
	"merch-shop/internal/handlers"
	"net/http"
)

type SendCoinRequest struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

func (hnd *UserHandler) SendCoins(wrt http.ResponseWriter, rqt *http.Request) {
	if rqt.Method != http.MethodPost {
		errSend := handlers.SendBadReq(wrt, "wrong http method")
		if errSend != nil {
			log.Printf("error while sending the bad request message: %v\n", errSend)
		}
		return
	}

	var srq SendCoinRequest
	err := json.NewDecoder(rqt.Body).Decode(&srq)
	if err != nil {
		errSend := handlers.SendBadReq(wrt, "wrong request body")
		if errSend != nil {
			log.Printf("error while sending the bad request message: %v\n", errSend)
		}
		return
	}

	if srq.ToUser == "" || srq.Amount == 0 {
		errSend := handlers.SendBadReq(wrt, "empty fields of request body")
		if errSend != nil {
			log.Printf("error while sending the bad request message: %v\n", errSend)
		}
		return
	}

	balance, userID := hnd.GetBalance(wrt, rqt)
	if balance == nil || userID == nil {
		return
	}

	transaction := coins.TransactionInDetail{
		SenderID:     *userID,
		ReceiverName: srq.ToUser,
		Balance:      *balance,
		Amount:       srq.Amount,
	}

	code, err := hnd.CoinsRepo.SendCoins(transaction)
	if err != nil {
		log.Println(err)
	}

	switch code {
	case 400:
		errSend := handlers.SendBadReq(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the bad request message: %v\n", errSend)
		}
		return

	case 500:
		errSend := handlers.SendInternalServerError(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the internal server error message: %v\n", errSend)
		}
		return
	}

	wrt.Header().Set("Content-Type", "application/json")
	wrt.WriteHeader(http.StatusOK)
}
