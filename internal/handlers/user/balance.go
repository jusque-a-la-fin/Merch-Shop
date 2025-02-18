package user

import (
	"log"
	"merch-shop/internal/handlers"
	"net/http"
)

func (hnd *UserHandler) GetBalance(wrt http.ResponseWriter, rqt *http.Request) (*int, *string) {
	thisUser, err := hnd.GetUser(wrt, rqt)
	if err != nil {
		errSend := handlers.SendUnauthorized(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the unauthorized error message: %v\n", errSend)
		}
		return nil, nil
	}

	userID, err := hnd.UserRepo.GetUserID(*thisUser)
	if err != nil {
		errSend := handlers.SendInternalServerError(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the internal server error message: %v\n", errSend)
		}
		return nil, nil
	}

	balance, err := hnd.CoinsRepo.GetBalance(*userID)
	if err != nil {
		errSend := handlers.SendInternalServerError(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the internal server error message: %v\n", errSend)
		}
		return nil, nil
	}
	return balance, userID
}
