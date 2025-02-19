package inventory

import (
	"database/sql"
	"sync"
)

type InventoryDBRepostitory struct {
	dtb   *sql.DB
	mutex *sync.Mutex
}

func NewDBRepo(sdb *sql.DB, mutex *sync.Mutex) *InventoryDBRepostitory {
	return &InventoryDBRepostitory{dtb: sdb, mutex: mutex}
}
