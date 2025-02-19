package coins

import (
	"database/sql"
	"sync"
)

type CoinsDBRepostitory struct {
	dtb   *sql.DB
	mutex *sync.Mutex
}

func NewDBRepo(sdb *sql.DB, mutex *sync.Mutex) *CoinsDBRepostitory {
	return &CoinsDBRepostitory{dtb: sdb, mutex: mutex}
}
