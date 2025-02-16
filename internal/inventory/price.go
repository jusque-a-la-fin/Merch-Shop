package inventory

import "merch-shop/internal/utils"

func (repo *InventoryDBRepostitory) GetPrice(itemType string) *int {
	exists := utils.CheckItem(repo.dtb, itemType)
	if !exists {
		return nil
	}

	var price int
	repo.dtb.QueryRow("SELECT price FROM items WHERE item_type = $1;", itemType).Scan(&price)
	return &price
}
