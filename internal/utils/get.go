package utils

import (
	"database/sql"
)

func GetUsername(dtb *sql.DB, userID string) string {
	var username string
	dtb.QueryRow("SELECT username FROM users WHERE id = $1;", userID).Scan(&username)
	return username
}

func GeReceiverID(dtb *sql.DB, username string) string {
	var userID string
	dtb.QueryRow("SELECT id FROM users WHERE username = $1;", username).Scan(&userID)
	if userID != "" {
		return userID
	}

	dtb.QueryRow("SELECT id FROM shop WHERE shopname = $1;", username).Scan(&userID)
	return userID
}
