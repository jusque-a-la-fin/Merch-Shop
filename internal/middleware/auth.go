package middleware

import (
	"database/sql"
	"log"
	"merch-shop/internal/token"
	"net/http"
)

func RequireAuth(next http.HandlerFunc, dtb *sql.DB) http.HandlerFunc {
	return func(wrt http.ResponseWriter, rqt *http.Request) {
		Authenticate(next, dtb).ServeHTTP(wrt, rqt)
	}
}

func Authenticate(next http.Handler, dtb *sql.DB) http.Handler {
	return http.HandlerFunc(func(wrt http.ResponseWriter, rqt *http.Request) {
		check, err := token.Check(rqt, dtb)
		if !check {
			log.Println("the token check has failed")
			return
		}
		if err != nil {
			log.Printf("error while checking the token: %v\n", err)
			return
		}
		next.ServeHTTP(wrt, rqt)
	})
}
