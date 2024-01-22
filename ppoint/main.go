package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"ppoint/db"
	"ppoint/gui"
	"ppoint/query"
	"ppoint/service"
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

	// test 2일 -- 수정 필요
	if err := service.ChangePointNoVisitFor3Month(DbConf); err != nil {
		panic(err)
	}
}

func main() {
	gui.MainPage(DbConf)
	//test.CmdTest(DbConf)
} // main
