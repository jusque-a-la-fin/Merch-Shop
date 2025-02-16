package middleware

import (
	"merch-shop/internal/session"
	"net/http"
)

func Authenticate(smg *session.SessionsManager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(wrt http.ResponseWriter, rqt *http.Request) {
		sess, err := smg.Check(wrt, rqt)
		if sess != nil && err == nil {
			ctx := session.ContextWithSession(rqt.Context(), sess)
			next.ServeHTTP(wrt, rqt.WithContext(ctx))
			return
		}
		next.ServeHTTP(wrt, rqt)
	})
}
