package session

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"merch-shop/internal/user"
)

type Session struct {
	ID       string
	UserName string
}

func NewSession(thisUser user.User) *Session {
	randID := make([]byte, 16)
	_, err := rand.Read(randID)
	if err != nil {
		log.Printf("error while creating session, %#v", err)
	}
	return &Session{
		ID:       fmt.Sprintf("%x", randID),
		UserName: thisUser.Username,
	}
}

type sessKey string

var (
	SessionKey         sessKey = sessKey("SessionKey")
	ExampleTokenSecret         = []byte("ExampleTokenSecret")
	ErrNoAuth                  = errors.New("no session found")
)

func SessionFromContext(ctx context.Context) (*Session, error) {
	sess, ok := ctx.Value(SessionKey).(*Session)
	if !ok || sess == nil {
		return nil, ErrNoAuth
	}
	return sess, nil
}

func ContextWithSession(ctx context.Context, sess *Session) context.Context {
	return context.WithValue(ctx, SessionKey, sess)
}
