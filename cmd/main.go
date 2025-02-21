package main

import (
	"log"
	"merch-shop/internal/coins"
	"merch-shop/internal/datastore"
	uhd "merch-shop/internal/handlers/user"
	"merch-shop/internal/inventory"
	"merch-shop/internal/middleware"
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
	coins := coins.NewDBRepo(dtb)
	inv := inventory.NewDBRepo(dtb)
	userHandler := &uhd.UserHandler{
		UserRepo:      usr,
		CoinsRepo:     coins,
		InventoryRepo: inv,
	}

	rtr := mux.NewRouter()
	apiRouter := rtr.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/info", middleware.RequireAuth(userHandler.GetInfo, dtb)).Methods("GET")
	apiRouter.HandleFunc("/sendCoin", middleware.RequireAuth(userHandler.SendCoins, dtb)).Methods("POST")
	apiRouter.HandleFunc("/buy/{item}", middleware.RequireAuth(userHandler.BuyAnItem, dtb)).Methods("GET")
	apiRouter.HandleFunc("/auth", userHandler.GetAuthenticated).Methods("POST")

	port := os.Getenv("SERVER_PORT")
	err = http.ListenAndServe(":"+port, rtr)
	if err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
