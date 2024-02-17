package gui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func RunRevenueDeleteDialog(owner walk.Form) (int, error) {
	var dlg *walk.Dialog
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "매출 삭제",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       Size{200, 200},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "삭제하시겠습니까?",
					},
					Label{
						Text: "* 삭제된 내역은 되살릴 수 없습니다.",
						//TextColor: walk.RGB(137, 0, 0),
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
