package db

import (
	"database/sql"
	"fmt"
	"log"
)

func Conn(user, passwd, dbname string) *sql.DB {
	var version string
	var db *sql.DB
	var err error

	db, err = sql.Open("mysql", fmt.Sprint(user, ":", passwd, "@tcp(localhost:3306)/", dbname))
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	db.QueryRow("SELECT VERSION()").Scan(&version)
	log.Println("Connected to:", version)

	return db
}

func DisConn(dbConn *sql.DB) {
	if dbConn != nil {
		dbConn.Close()
	}
}
