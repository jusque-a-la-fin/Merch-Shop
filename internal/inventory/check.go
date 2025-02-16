package inventory

import "database/sql"

func CheckInventory(dtb *sql.DB, itemID int) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM inventory WHERE item_id = $1);`
	dtb.QueryRow(query, itemID).Scan(&exists)
	return exists
}
