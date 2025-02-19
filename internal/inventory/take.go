package inventory

import "fmt"

func (repo *InventoryDBRepostitory) TakeAnItem(userID string, itemType string) error {
	var itemID int
	err := repo.dtb.QueryRow("SELECT id FROM items WHERE item_type = $1;", itemType).Scan(&itemID)
	if err != nil {
		return fmt.Errorf("error while selecting the id of the item: %v", err)
	}

	exists, err := CheckInventory(repo.dtb, repo.mutex, itemID)
	if err != nil {
		return err
	}

	if !exists {
		query := `INSERT INTO inventory (user_id, item_id, quantity) VALUES ($1, $2, $3);`
		_, err := repo.dtb.Exec(query, userID, itemID, 1)
		if err != nil {
			return fmt.Errorf("error while adding a new item: %v", err)
		}
		return nil
	}

	query := `
		     UPDATE inventory
		     SET quantity = quantity + 1
		     WHERE user_id = $1 AND item_id = $2;`

	_, err = repo.dtb.Exec(query, userID, itemID)
	if err != nil {
		return fmt.Errorf("error while updating the quantity of the item: %v", err)
	}
	return nil
}
