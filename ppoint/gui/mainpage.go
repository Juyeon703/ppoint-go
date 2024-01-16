package gui

import (
	"bytes"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/query"
)

type AppMainWindow struct {
	*MultiPageMainWindow
}

// MultiPageMainWindow 최상위 메인
var winMain *AppMainWindow

var dbconn *query.DbConfig
var titleWidth, titleHeight = 1000, 800

func MainPage(DbConf *query.DbConfig) {
	dbconn = DbConf
	walk.Resources.SetRootDirPath("img")
	var err error
	//var titleWidth, titleHeight = 500, 500

	winMain = new(AppMainWindow)

	//multiple page main
	var multiPageMainWindow *MultiPageMainWindow
	multiPageMainWindow = new(MultiPageMainWindow)

	//Multi Page Main Window Config
	cfg := &MultiPageMainWindowConfig{
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
			},
		},

		//페이지 변경시마다 업데이트
		//OnCurrentPageChanged: func() {
		//	winMain.updateTitle(winMain.CurrentPageTitle())
		//},
		ToolBar: ToolBar{
			MinSize: Size{titleWidth, 100},
			MaxSize: Size{titleWidth, 100},
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
		panic(err)
	}
	winMain.MultiPageMainWindow = multiPageMainWindow

	//winMain.updateTitle(winMain.CurrentPageTitle())

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

func test() {
	/*	walk.AppendToWalkInit(func() {
			walk.FocusEffect, _ = walk.NewBorderGlowEffect(walk.RGB(0, 63, 255))
			walk.InteractionEffect, _ = walk.NewDropShadowEffect(walk.RGB(63, 63, 63))
			walk.ValidationErrorEffect, _ = walk.NewBorderGlowEffect(walk.RGB(255, 0, 0))
		})

		var inTE, outTE *walk.TextEdit
		var titleWidth, titleHeight = 1000, 700

		var winMain *walk.MainWindow
		winMain = new(walk.MainWindow)

		//win.SetWindowPos(winMain.Handle(),
		//win.HWND_TOPMOST, 0, 0, 0, 0,
		//win.SWP_NOACTIVATE|win.SWP_NOMOVE|win.SWP_NOSIZE)

		if _, err := (MainWindow{
			AssignTo: &winMain,
			Title:    "포인트 프로그램",
			MinSize:  Size{300, 200},
			MaxSize:  Size{2000, 1600},
			Size:     Size{titleWidth, titleHeight},
			Layout:   VBox{},
			Children: []Widget{
				HSplitter{
					MaxSize: Size{titleWidth, 50},
					Children: []Widget{
						PushButton{
							Text: "포인트 관리",
							OnClicked: func() {
								//outTE.SetText(strings.ToUpper(inTE.Text()))
							},
						},
						PushButton{
							Text: "고객 관리",
							OnClicked: func() {
							},
						},
						PushButton{
							Text: "매출 관리",
							OnClicked: func() {
							},
						},
						PushButton{
							Text: "설정",
							OnClicked: func() {
							},
						},
					},
				}, //end of HSplitter
				HSplitter{
					MaxSize: Size{titleWidth, 150},
					Children: []Widget{
						//Label{
						//	Text: "test",
						//},
						HSplitter{
							Children: []Widget{
								TextEdit{AssignTo: &inTE},
								TextEdit{AssignTo: &outTE, ReadOnly: true},
							},
						},
						TextEdit{AssignTo: &inTE},
						TextEdit{AssignTo: &outTE, ReadOnly: true},
					},
				}, //end of HSplitter
				HSplitter{
					MaxSize: Size{titleWidth, 400},
					Children: []Widget{
						TextEdit{AssignTo: &inTE},
						TextEdit{AssignTo: &outTE, ReadOnly: true},
					},
				}, //end of HSplitter
			}, //end of children_main
		}).Run(); err != nil {
			log.Fatal(err)
		}
	*/
	//sample
	/*	MainWindow{
		Title: "포인트 프로그램",
		//MinSize: Size{600, 400},
		Size:   Size{1000, 700},
		Layout: VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()*/
}
