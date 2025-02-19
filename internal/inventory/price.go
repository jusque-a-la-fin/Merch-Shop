package inventory

import (
	"fmt"
	"merch-shop/internal/utils"
)

func (repo *InventoryDBRepostitory) GetPrice(itemType string) (*int, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	exists, err := utils.CheckItem(repo.dtb, itemType)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}

	var price int
	err = repo.dtb.QueryRow("SELECT price FROM items WHERE item_type = $1;", itemType).Scan(&price)
	if err != nil {
		return nil, fmt.Errorf("error while selecting the price of the item: %v", err)
	}
	return &price, nil
}
