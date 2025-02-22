package user

import (
	"database/sql"
	"fmt"
)

type UserDBRepostitory struct {
	dtb *sql.DB
}

func NewDBRepo(sdb *sql.DB) *UserDBRepostitory {
	return &UserDBRepostitory{dtb: sdb}
}

func (repo *UserDBRepostitory) GetUserID(username string) (*string, error) {
	var userID string
	err := repo.dtb.QueryRow("SELECT id FROM users WHERE username = $1;", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error while selecting the user id: %v", err)
	}
	return &userID, nil
}
