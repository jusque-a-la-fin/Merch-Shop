package coins

import (
	"database/sql"
)

type CoinsDBRepostitory struct {
	dtb *sql.DB
}

func NewDBRepo(sdb *sql.DB) *CoinsDBRepostitory {
	return &CoinsDBRepostitory{dtb: sdb}
}
