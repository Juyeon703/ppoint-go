package db

import (
	"database/sql"
	"fmt"
)

func Conn(user, passwd, dbname string) *sql.DB {
	var version string
	var db *sql.DB
	var err error

	db, err = sql.Open("mysql", fmt.Sprint(user, ":", passwd, "@tcp(localhost:3306)/", dbname))
	if err != nil {
		panic(err.Error())
	}

	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	return db
}
