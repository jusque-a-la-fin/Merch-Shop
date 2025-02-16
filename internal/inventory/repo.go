package inventory

import "database/sql"

type InventoryDBRepostitory struct {
	dtb *sql.DB
}

func NewDBRepo(sdb *sql.DB) *InventoryDBRepostitory {
	return &InventoryDBRepostitory{dtb: sdb}
}
