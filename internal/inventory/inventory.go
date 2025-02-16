package inventory

type InventoryRepo interface {
	Get(userID int) ([]Item, error)
	GetPrice(itemType string) *int
	TakeAnItem(userID int, itemType string) error
}

type Item struct {
	Item_type string `json:"type"`
	Quantity  int    `json:"quantity"`
}
