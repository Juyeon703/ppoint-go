package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"ppoint/dto"
)

func RunMemberAddDialog(owner walk.Form, member *dto.MemberAddDto) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

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
						Text:          Bind("PhoneNumber", Regexp{Pattern: "010([0-9]{7,8}$)"}, SelRequired{}),
						TextAlignment: AlignCenter,
					},
					Label{
						Text:          "생일 : ",
						TextAlignment: AlignFar,
					},
					// DateEdit
					LineEdit{
						Text: Bind("Birth", Regexp{Pattern: "(19[0-9][0-9]|20[0-9][0-9])-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$"}),
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
