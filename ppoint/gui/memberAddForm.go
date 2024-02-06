package gui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/dto"
	"ppoint/service"
	"ppoint/utils"
	"time"
)

func RunMemberAddDialog(owner walk.Form, member *dto.MemberAddDto) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton
	var dateEditor *walk.DateEdit
	var phoneNumLE *walk.LineEdit

	return Dialog{
		AssignTo:      &dlg,
		Title:         "회원 등록",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:   &db,
			Name:       "member",
			DataSource: member,
			//ErrorPresenter: ToolTipErrorPresenter{},
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
						AssignTo:      &phoneNumLE,
						Text:          Bind("PhoneNumber", Regexp{Pattern: "^01([0|1|6|7|8|9])-([0-9]{3,4})-([0-9]{4})|01([0|1|6|7|8|9])([0-9]{3,4})([0-9]{4})$"}, SelRequired{}),
						TextAlignment: AlignCenter,
						OnEditingFinished: func() {
							phoneNumLE.SetText(utils.PhoneNumAddHyphen(phoneNumLE.Text()))
						},
					},
					Label{
						Text:          "생일 : ",
						TextAlignment: AlignFar,
					},
					DateEdit{
						AssignTo: &dateEditor,
						Date:     Bind("Birth"),
						OnBoundsChanged: func() {
							dateEditor.SetDate(time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local))
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
								log.Error(err.Error())
								panic(err)
								return
							}
							if existMember, err := service.FindMemberPhoneNumber(dbconn, member.PhoneNumber); existMember != nil {
								MsgBox("알림", "이미 존재하는 핸드폰 번호 입니다.")
							} else {
								if err = service.MemberAdd(dbconn, member); err != nil {
									MsgBox("알림", "고객 등록에 실패하였습니다.")
								} else {
									dlg.Accept()
								}

							}
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
