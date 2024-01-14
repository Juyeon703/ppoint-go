package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"ppoint/db"
	"ppoint/gui"
	"ppoint/query"
)

var DbConf *query.DbConfig

func init() {
	DbConf = new(query.DbConfig)
	DbConf.DbConnection = new(sql.DB)

	var dbConn *sql.DB
	dbConn = db.Conn("root", "1111", "ppoint")
	if dbConn == nil {
		panic("db conn nil")
	}
	DbConf.DbConnection = dbConn

}

func main() {
	gui.MainPage()
	//test.CmdTest(DbConf)
} // main
