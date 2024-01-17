package gui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
)

func RunMemberEditDialog(owner walk.Form, member *MDto) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "회원 정보 수정",
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
						Text: "이름 :",
					},
					LineEdit{
						Text: Bind("Name", SelRequired{}),
					},
					Label{
						Text: "등급 :",
					},
					ComboBox{
						Value: Bind("GradeName", SelRequired{}),
						//BindingMember: "GradeId",
						//DisplayMember: "GradeName",
						//Model:         KnownGrades(),
						Model: []string{"잠재", "신규", "우수", "평생"},
					},
					Label{
						Text: "핸드폰 번호:",
					},
					LineEdit{
						Text: Bind("PhoneNumber", SelRequired{}),
					},
					RadioButtonGroupBox{
						ColumnSpan: 2,
						Title:      "성별",
						Layout:     HBox{},
						DataMember: "Sex",
						Buttons: []RadioButton{
							{Text: "남", Value: "SexMale"},
							{Text: "여", Value: "SexFemale"},
						},
					},
					Label{
						Text: "생일:",
					},
					DateEdit{
						Date: Bind("Birth"),
					},
					Label{
						Text: "보유 포인트:",
					},
					NumberEdit{
						Value:  Bind("Point", SelRequired{}),
						Suffix: " 포인트",
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

type MDto struct {
	Name        string
	GradeName   string
	PhoneNumber string
	Sex         string
	Birth       string
	Point       int
}
