package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"ppoint/dto"
	"time"
)

func RunMemberAddDialog(owner walk.Form, member *dto.MemberAddDto) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	var dateEditor = new(walk.DateEdit)

	return Dialog{
		AssignTo:      &dlg,
		Title:         "회원 등록",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "member",
			DataSource:     member,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text:          "이름 : ",
						TextAlignment: AlignFar,
					},
					LineEdit{
						Text:          Bind("MemberName", SelRequired{}),
						TextAlignment: AlignCenter,
					},
					Label{
						Text:          "핸드폰 번호 : ",
						TextAlignment: AlignFar,
					},
					LineEdit{
						Text:          Bind("PhoneNumber", Regexp{Pattern: "^01([0|1|6|7|8|9])-([0-9]{3,4})-([0-9]{4})$"}, SelRequired{}),
						TextAlignment: AlignCenter,
					},
					Label{
						Text:          "생일 : ",
						TextAlignment: AlignFar,
					},
					// DateEdit
					DateEdit{
						AssignTo: &dateEditor,
						Date:     Bind("Birth"),
						OnBoundsChanged: func() {
							dateEditor.SetDate(time.Now().Add(-time.Duration(72000) * time.Hour))
						},
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}
							if memberId, err := dbconn.CreateMember(member); err != nil {
								log.Print(err)
								return
							} else {
								fmt.Println("======> CreateMember() 호출")
								fmt.Println("등록된 회원 번호 : ", memberId)
							}
							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}