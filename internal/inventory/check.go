package inventory

import (
	"database/sql"
	"fmt"
	"sync"
)

func CheckInventory(dtb *sql.DB, mutex *sync.Mutex, itemID int) (bool, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM inventory WHERE item_id = $1);`
	err := dtb.QueryRow(query, itemID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error while checking if the item exists: %v", err)
	}
	return exists, nil
}
