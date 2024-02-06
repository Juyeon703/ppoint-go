package query

import (
	"database/sql"
	"ppoint/logue"
)

type DbConfig struct {
	DbConnection *sql.DB
	Logue        *logue.Logbook
}
