package user

import (
	"database/sql"
	"fmt"
)

func CreateUser(dtb *sql.DB, user User) (*User, error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	var userID int
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id;`
	_ = dtb.QueryRow(query, user.Username, hashedPassword).Scan(&userID)

	query = `INSERT INTO receivers (id) VALUES ($1);`
	_, err = dtb.Exec(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error while creating a new user: %v", err)
	}

	thisUser := User{Username: user.Username, Password: hashedPassword}
	return &thisUser, nil
}
