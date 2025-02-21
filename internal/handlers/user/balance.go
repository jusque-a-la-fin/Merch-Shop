package user

import (
	"log"
	"merch-shop/internal/handlers"
	"merch-shop/internal/token"
	"net/http"
)

func (hnd *UserHandler) GetBalance(wrt http.ResponseWriter, rqt *http.Request) (*int, *string) {
	username, err := token.GetPayload(rqt)
	if err != nil {
		errSend := handlers.SendUnauthorized(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the unauthorized error message: %v\n", errSend)
		}
	}

	userID, err := hnd.UserRepo.GetUserID(username)
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

func (hnd *UserHandler) UpdateBalance(wrt http.ResponseWriter, rqt *http.Request, username string) {
	if username == "" {
		var err error
		username, err = token.GetPayload(rqt)
		if err != nil {
			errSend := handlers.SendUnauthorized(wrt, err.Error())
			if errSend != nil {
				log.Printf("error while sending the unauthorized error message: %v\n", errSend)
			}
		}
	}

	userID, err := hnd.UserRepo.GetUserID(username)
	if err != nil {
		errSend := handlers.SendInternalServerError(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the internal server error message: %v\n", errSend)
		}
	}

	err = hnd.CoinsRepo.UpdateBalance(*userID)
	if err != nil {
		errSend := handlers.SendInternalServerError(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the internal server error message: %v\n", errSend)
		}
	}
}
