package user

import (
	"log"
	"merch-shop/internal/coins"
	"merch-shop/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func (hnd *UserHandler) BuyAnItem(wrt http.ResponseWriter, rqt *http.Request) {
	if rqt.Method != http.MethodGet {
		errSend := handlers.SendBadReq(wrt, "wrong http method")
		if errSend != nil {
			log.Printf("error while sending the bad request message: %v\n", errSend)
		}
		return
	}

	itemType := mux.Vars(rqt)["item"]
	price, err := hnd.InventoryRepo.GetPrice(itemType)
	if err != nil {
		log.Println(err)
	}

	if price == nil {
		log.Println("the item has not been found")
		errSend := handlers.SendBadReq(wrt, "the item has not been found")
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
		ReceiverName: "avito-shop",
		Balance:      *balance,
		Amount:       *price,
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

	err = hnd.InventoryRepo.TakeAnItem(*userID, itemType)
	if err != nil {
		errSend := handlers.SendInternalServerError(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the internal server error message: %v\n", errSend)
		}
		return
	}

	wrt.WriteHeader(http.StatusOK)
}
