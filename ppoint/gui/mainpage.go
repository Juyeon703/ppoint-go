package gui

import (
	"bytes"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/backup"
	"ppoint/db"
	"ppoint/logue"
	"ppoint/query"
	"sync"
)

type AppMainWindow struct {
	*MultiPageMainWindow
}

// MultiPageMainWindow 최상위 메인
var winMain *AppMainWindow
var dbconn *query.DbConfig
var log *logue.Logbook
var titleWidth, titleHeight = 950, 700
var subWidth, subHeight = 950, 700
var toolbarHeight = 60
var cfg *MultiPageMainWindowConfig
var multiPageMainWindow *MultiPageMainWindow

func MainPage(DbConf *query.DbConfig) {
	dbconn = DbConf
	log = DbConf.Logue
	walk.Resources.SetRootDirPath("img")
	var err error

	//main window
	winMain = new(AppMainWindow)

	//multiple page main
	multiPageMainWindow = new(MultiPageMainWindow)

	//Multi Page Main Window Config
	cfg = &MultiPageMainWindowConfig{
		Title: "PPOINT",
		Name:  "mainWindow",
		Size:  Size{titleWidth, titleHeight},
		//MENU ITEM
		MenuItems: []MenuItem{
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text:        "About",
						OnTriggered: func() { winMain.aboutAction_Triggered() },
					},
				},

				//페이지 변경시마다 업데이트
				//OnCurrentPageChanged: func() {
			},
		},
		//	winMain.updateTitle(winMain.CurrentPageTitle())
		//},
		ToolBar: ToolBar{
			Font:    Font{Bold: true},
			MinSize: Size{titleWidth, toolbarHeight},
			MaxSize: Size{titleWidth, toolbarHeight},
		},

		//페이지탭
		PageCfgs: []PageConfig{
			{"포인트 관리", "document-new.png", newPointPage},
			{"고객 관리", "document-new.png", newMemberPage},
			{"매출 관리", "document-new.png", newSalesPage},
			{"설정 페이지", "document-new.png", newSettingPage},
		},
	}

	multiPageMainWindow, err = NewMultiPageMainWindow(cfg)
	if err != nil {
		log.Error("Main Window Create 실패", err.Error())
		panic(err)
	}
	winMain.MultiPageMainWindow = multiPageMainWindow

	//winMain.updateTitle(winMain.CurrentPageTitle())
	winMain.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		if reason == walk.CloseReasonUser {
			*canceled = true
		}

		var wg sync.WaitGroup
		var backupResult chan bool
		backupResult = make(chan bool)

		wg.Add(1)

		go func() {
			backup.DbBackup(dbconn, &wg, backupResult)
			db.DisConn(DbConf.DbConnection)
		}()

		//마지막 db connection 종료
		flag := true
		for flag {
			select {
			case <-backupResult:
				flag = false
				break
			default:
				MsgBox("알림", "데이터 백업 중입니다. 잠시 기다려주세요")
			}
		}

		close(backupResult)
		log.Debug("포인트 프로그램 CLOSE")
	})

	winMain.Run()
}

// main title 변경 해주는 로직
func (mw *AppMainWindow) updateTitle(prefix string) {
	var buf bytes.Buffer

	if prefix != "" {
		buf.WriteString(prefix)
		buf.WriteString(" - ")
	}

	buf.WriteString("Walk Multiple Pages Example")

	mw.SetTitle(buf.String())
}

// about 클릭시 메세지 박스 표출
func (mw *AppMainWindow) aboutAction_Triggered() {
	walk.MsgBox(mw,
		"About Walk Multiple Pages Example",
		"An example that demonstrates a main window that supports multiple pages.",
		walk.MsgBoxOK|walk.MsgBoxIconInformation)
}
