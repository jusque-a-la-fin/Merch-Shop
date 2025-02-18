package inventory

type InventoryRepo interface {
	Get(userID string) ([]Item, error)
	GetPrice(itemType string) (*int, error)
	TakeAnItem(userID string, itemType string) error
}

type Item struct {
	Item_type string `json:"type"`
	Quantity  int    `json:"quantity"`
}
