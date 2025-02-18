package utils

import (
	"database/sql"
	"fmt"
)

func GetUsername(dtb *sql.DB, userID string) (*string, error) {
	var username string
	err := dtb.QueryRow("SELECT username FROM users WHERE id = $1;", userID).Scan(&username)
	if err != nil {
		return nil, fmt.Errorf("error while selecting the username of the user: %v", err)
	}
	return &username, nil
}

func GeReceiverID(dtb *sql.DB, username string) (*string, error) {
	var userID string
	err := dtb.QueryRow("SELECT id FROM users WHERE username = $1;", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error while selecting the id of the user: %v", err)
	}

	if userID != "" {
		return &userID, nil
	}

	err = dtb.QueryRow("SELECT id FROM shop WHERE shopname = $1;", username).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("error while selecting the id of the shop: %v", err)
	}
	return &userID, nil
}
