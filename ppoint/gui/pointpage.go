package gui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type PointPage struct {
	*walk.Composite
}

/*
VSpacer{
	ColumnSpan: 2,
	Size:       8,
},
*/

func newPointPage(parent walk.Container) (Page, error) {
	p := new(PointPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "포인트 관리",
		Layout:   VBox{},
		Border:   true,
		Children: []Widget{
			Label{Text: "포인트 관리 페이지"},
			VSplitter{
				Children: []Widget{
					HSplitter{
						Children: []Widget{
							Label{
								Name: "search_phone_number",
								Text: "핸드폰 번호",
							},
							LineEdit{
								Text: Bind("search_phone_number"),
							},
						},
					},
					VSplitter{
						Children: []Widget{
							RadioButtonGroupBox{
								ColumnSpan: 2,
								Title:      "포인트 적립/사용 구분",
								Layout:     HBox{},
								Buttons: []RadioButton{
									{Text: "적립", Value: "types.ADD_POINT"},
									{Text: "사용", Value: "types.SUB_POINT"},
								},
							},
							RadioButtonGroupBox{
								ColumnSpan: 2,
								Title:      "결제 유형",
								Layout:     HBox{},
								Buttons: []RadioButton{
									{Text: "카드", Value: "types.TYPE_CARD"},
									{Text: "현금", Value: "types.TYPE_CASH"},
								},
							},
							Label{
								Name: "payment",
								Text: "결제 금액",
							},
							NumberEdit{
								Value:    Bind("payment", Range{1, 999999999}),
								Suffix:   " 원",
								Decimals: 0,
							},
						},
					},
				},
			},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}
