package utils

import (
	"database/sql"
	"fmt"
)

func CheckUser(dtb *sql.DB, username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1);`
	err := dtb.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error while checking if the user exists: %v", err)
	}
	return exists, nil
}

func CheckShop(dtb *sql.DB, shopname string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM shop WHERE shopname = $1);`
	err := dtb.QueryRow(query, shopname).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error while checking if the shop exists: %v", err)
	}
	return exists, nil
}

func CheckItem(dtb *sql.DB, itemType string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM items WHERE item_type = $1);`
	err := dtb.QueryRow(query, itemType).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error while checking if the item exists: %v", err)
	}
	return exists, nil
}
