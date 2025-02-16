package utils

import (
	"database/sql"
)

func CheckUser(dtb *sql.DB, username string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1);`
	dtb.QueryRow(query, username).Scan(&exists)
	return exists
}

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

func CheckShop(dtb *sql.DB, shopname string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM shop WHERE shopname = $1);`
	dtb.QueryRow(query, shopname).Scan(&exists)
	return exists
}

func CheckItem(dtb *sql.DB, itemType string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM items WHERE item_type = $1);`
	dtb.QueryRow(query, itemType).Scan(&exists)
	return exists
}
