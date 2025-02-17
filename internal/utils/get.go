package utils

import "database/sql"

func GetUsername(dtb *sql.DB, userID int) string {
	var username string
	dtb.QueryRow("SELECT username FROM users WHERE id = $1;", userID).Scan(&username)
	return username
}

func GetUserID(dtb *sql.DB, username string) int {
	var userID int
	dtb.QueryRow("SELECT id FROM users WHERE username = $1;", username).Scan(&userID)
	return userID
}
