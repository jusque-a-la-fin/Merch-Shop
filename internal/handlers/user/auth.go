package user

import (
	"encoding/json"
	"log"
	"merch-shop/internal/handlers"
	"merch-shop/internal/user"
	"net/http"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func (hnd *UserHandler) GetAuthenticated(wrt http.ResponseWriter, rqt *http.Request) {
	if rqt.Method != http.MethodPost {
		errSend := handlers.SendBadReq(wrt, "wrong http method")
		if errSend != nil {
			log.Printf("error while sending the bad request message: %v\n", errSend)
		}
		return
	}

	var arq AuthRequest
	err := json.NewDecoder(rqt.Body).Decode(&arq)
	if err != nil {
		errSend := handlers.SendBadReq(wrt, "wrong request body")
		if errSend != nil {
			log.Printf("error while sending the bad request message: %v\n", errSend)
		}
		return
	}

	if arq.Username == "" || arq.Password == "" {
		errSend := handlers.SendBadReq(wrt, "empty fields of request body")
		if errSend != nil {
			log.Printf("error while sending the bad request message: %v\n", errSend)
		}
		return
	}

	usr := user.User{
		Username: arq.Username,
		Password: arq.Password,
	}

	user, code, err := hnd.UserRepo.GetAuthenticated(usr)
	if err != nil {
		log.Println(err)
	}

	switch code {
	case 401:
		errSend := handlers.SendUnauthorized(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the unauthorized error message: %v\n", errSend)
		}
		return

	case 500:
		errSend := handlers.SendInternalServerError(wrt, err.Error())
		if errSend != nil {
			log.Printf("error while sending the internal server error message: %v\n", errSend)
		}
		return
	}

	tokenString := hnd.CreateSessionAndToken(wrt, user)
	resp := AuthResponse{
		Token: tokenString,
	}

	wrt.Header().Set("Content-Type", "application/json")
	wrt.WriteHeader(http.StatusOK)
	errJSON := json.NewEncoder(wrt).Encode(resp)
	if errJSON != nil {
		log.Printf("error while sending response body: %v\n", errJSON)
	}
}

func (hnd *UserHandler) CreateSessionAndToken(wrt http.ResponseWriter, thisUser *user.User) string {
	sess := hnd.Sessions.CreateSession(*thisUser)
	tokenString, errToken := hnd.Sessions.CreateJWTtoken(sess)
	if errToken != nil {
		errSend := handlers.SendInternalServerError(wrt, errToken.Error())
		if errSend != nil {
			log.Printf("error while sending the internal server error message: %v\n", errSend)
		}

		log.Println("error while creating the JWT token: ", errToken)
		return ""
	}
	return tokenString
}
