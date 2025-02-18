package inventory

import (
	"fmt"
)

func (repo *InventoryDBRepostitory) Get(userID string) ([]Item, error) {
	query := `SELECT itm.item_type, inv.quantity 
	           FROM inventory inv 
			   JOIN items itm ON inv.item_id = itm.id 
			   WHERE inv.user_id = $1;`

	rows, err := repo.dtb.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error while selecting the inventory info: %v", err)
	}
	defer rows.Close()

	items := make([]Item, 0)
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Item_type, &item.Quantity)
		if err != nil {
			return nil, fmt.Errorf("error from method `Scan`, package sql: %v", err)
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating over rows returned by query: %v", err)
	}

	return items, nil
}
