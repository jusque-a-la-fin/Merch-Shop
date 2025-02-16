package main

import (
	"log"
	"merch-shop/internal/coins"
	"merch-shop/internal/datastore"
	uhd "merch-shop/internal/handlers/user"
	"merch-shop/internal/inventory"
	"merch-shop/internal/middleware"
	"merch-shop/internal/session"
	"merch-shop/internal/user"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	dtb, err := datastore.CreateNewDB()
	if err != nil {
		log.Fatalf("error while connecting to the database: %v", err)
	}

	usr := user.NewDBRepo(dtb)
	smg := session.NewSessionsManager()
	coins := coins.NewDBRepo(dtb)
	inv := inventory.NewDBRepo(dtb)
	userHandler := &uhd.UserHandler{
		UserRepo:      usr,
		Sessions:      smg,
		CoinsRepo:     coins,
		InventoryRepo: inv,
	}

	rtr := mux.NewRouter()
	apiRouter := rtr.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/info", requireAuth(smg, userHandler.GetInfo)).Methods("GET")
	apiRouter.HandleFunc("/sendCoin", requireAuth(smg, userHandler.SendCoins)).Methods("POST")
	apiRouter.HandleFunc("/buy/{item}", requireAuth(smg, userHandler.BuyAnItem)).Methods("GET")
	apiRouter.HandleFunc("/auth", userHandler.GetAuthenticated).Methods("POST")

	port := os.Getenv("SERVER_PORT")
	err = http.ListenAndServe(":"+port, rtr)
	if err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}

func requireAuth(smg *session.SessionsManager, next http.HandlerFunc) http.HandlerFunc {
	return func(wrt http.ResponseWriter, rqt *http.Request) {
		middleware.Authenticate(smg, next).ServeHTTP(wrt, rqt)
	}
}
