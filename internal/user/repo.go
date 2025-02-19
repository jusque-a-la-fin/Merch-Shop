package user

import (
	"database/sql"
	"fmt"
	"sync"
)

type UserDBRepostitory struct {
	dtb   *sql.DB
	mutex *sync.Mutex
}

func NewDBRepo(sdb *sql.DB, mutex *sync.Mutex) *UserDBRepostitory {
	return &UserDBRepostitory{dtb: sdb, mutex: mutex}
}

func (repo *UserDBRepostitory) GetUserID(usr User) (*string, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	var userID string
	err := repo.dtb.QueryRow("SELECT id FROM users WHERE username = $1;", usr.Username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error while selecting the user id: %v", err)
	}
	return &userID, nil
}
