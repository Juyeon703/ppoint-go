package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"ppoint/db"
	"ppoint/gui"
	"ppoint/logue"
	"ppoint/query"
	"ppoint/service"
	"ppoint/utils"
)

var DbConf *query.DbConfig
var log *logue.Logbook

func init() {
	var err error
	DbConf = new(query.DbConfig)
	DbConf.DbConnection = new(sql.DB)

	var dbConn *sql.DB
	dbConn = db.Conn("root", "1111", "ppoint")
	if dbConn == nil {
		panic("db conn nil")
	}
	DbConf.DbConnection = dbConn

	var logPath string
	if logPath, err = service.FindSettingStrValue(DbConf, "log_path"); err != nil {
		log.Error(err)
		panic(err)
	}

	if !utils.IsExistFile(logPath) {
		if err = utils.CreateFilePath(logPath); err != nil {
			log.Error(err)
			panic(err)
		}
	}

	if log, err = logue.Setup(logPath, "ppoint", 5); err != nil {
		log.Error(err)
		panic(err)
	}

	// test 2일 -- 수정 필요
	if err = service.ChangePointNoVisitFor3Month(DbConf); err != nil {
		log.Error(err)
		panic(err)
	}
}

func main() {
	log.Debug("MAIN START")
	gui.MainPage(DbConf)
} // main
