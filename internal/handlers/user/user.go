package user

import (
	"merch-shop/internal/coins"
	"merch-shop/internal/inventory"
	"merch-shop/internal/user"
)

type UserHandler struct {
	UserRepo      user.UserRepo
	CoinsRepo     coins.CoinsRepo
	InventoryRepo inventory.InventoryRepo
}
