package query

import "database/sql"

type DbConfig struct {
	DbConnection *sql.DB
}
