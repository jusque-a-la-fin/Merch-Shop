package user

import (
	"merch-shop/internal/coins"
	"merch-shop/internal/inventory"
	"merch-shop/internal/session"
	"merch-shop/internal/user"
	"net/http"
)

type UserHandler struct {
	UserRepo      user.UserRepo
	Sessions      *session.SessionsManager
	CoinsRepo     coins.CoinsRepo
	InventoryRepo inventory.InventoryRepo
}

func (hnd *UserHandler) GetUser(wrt http.ResponseWriter, rqt *http.Request) (*user.User, error) {
	thisSession, errSession := session.SessionFromContext(rqt.Context())
	if errSession != nil {
		return nil, errSession
	}

	thisUser := &user.User{}
	thisUser.Username = thisSession.UserName
	return thisUser, nil
}
