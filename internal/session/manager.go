package session

import (
	"fmt"
	"merch-shop/internal/user"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SessionsManager struct {
	data map[string]*Session
	mu   *sync.RWMutex
}

func NewSessionsManager() *SessionsManager {
	return &SessionsManager{
		data: make(map[string]*Session, 10),
		mu:   &sync.RWMutex{},
	}
}

func (sm *SessionsManager) CreateSession(newUser user.User) *Session {
	sess := NewSession(newUser)
	return sess
}

func (sm *SessionsManager) CreateJWTtoken(sess *Session) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": map[string]interface{}{
			"username": sess.UserName,
		},
		"iat": time.Now().Unix(),
		"exp": time.Now().Unix() + 1300,
	})
	tokenString, err := token.SignedString(ExampleTokenSecret)

	return tokenString, err
}

func (sm *SessionsManager) Check(w http.ResponseWriter, r *http.Request) (*Session, error) {
	inToken := r.Header.Get("Authorization")
	if inToken == "" {
		return nil, ErrNoAuth
	}

	hashSecretGetter := func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method.Alg() != "HS256" {
			return nil, fmt.Errorf("bad sign method")
		}
		return ExampleTokenSecret, nil
	}

	inToken = strings.Split(inToken, " ")[1]
	token, errJwt := jwt.Parse(inToken, hashSecretGetter)
	if errJwt != nil {
		return nil, errJwt
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if ok {
		sessClaims := payload["user"].(map[string]interface{})

		sess := &Session{}
		sess.UserName = sessClaims["username"].(string)

		return sess, nil
	}
	return nil, ErrNoAuth
}
