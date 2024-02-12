package main

import (
	"database/sql"
	"fmt"
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
		panic(err)
	}
	fmt.Printf("LOG_PATH : [%s]\n", logPath)
	if !utils.IsExistFile(logPath) {
		if err = utils.CreateFilePath(logPath); err != nil {
			fmt.Printf("LOG 파일 생성 실패\n")
			panic(err)
		}
	}
	if log, err = logue.Setup(logPath, "ppoint", 5); err != nil {
		panic(err)
	}
	DbConf.Logue = log
	log.Debug("(LOG SETUP) >>> SUCCESS")

	log.Debug("(3개월 미상 미접속 사용자 포인트 초기화) >>> START")
	if err = service.ChangePointNoVisitFor3Month(DbConf); err != nil {
		log.Error(err)
		panic(err)
	}
	log.Debug("(3개월 미상 미접속 사용자 포인트 초기화) >>> END")

}

func main() {
	log.Debug("포인트 프로그램 START")
	gui.MainPage(DbConf)
} // main
