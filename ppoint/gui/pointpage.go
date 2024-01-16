package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/types"
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

	var searchPhoneNumber *walk.LineEdit
	searchPhoneNumber = new(walk.LineEdit)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "포인트 관리",
		Layout:   VBox{},
		Border:   true,
		Children: []Widget{
			Label{Text: "포인트 관리 페이지"},
			VSpacer{
				ColumnSpan: 2,
				Size:       10,
			},
			Composite{
				Border: true,
				Layout: Grid{Columns: 10},
				Children: []Widget{
					Label{
						Name: "search_phone_number",
						Text: "핸드폰 번호",
					},
					LineEdit{
						AssignTo: &searchPhoneNumber,
						Text:     Bind("search_phone_number"),
					},
					PushButton{
						Text: "고객 조회",
						OnClicked: func() {
							var memberList []types.Member
							var err error

							if searchPhoneNumber.Text() == "" {
								MsgBox("고객 조회 에러", "핸드폰번호를 입력해주세요")
								return
							}

							if memberList, err = dbconn.SelectMemberSearch(fmt.Sprintf(searchPhoneNumber.Text())); err != nil {
								panic(err.Error())
							} else {
								if len(memberList) <= 0 {
									MsgBox("고객 조회 에러", "존재하지 않는 고객입니다.")
								} else {
									fmt.Printf("%v", memberList)
								}
							}
						},
					},
				},
			},
			VSpacer{
				ColumnSpan: 2,
				Size:       50,
			},

			//조회 결과
			HSplitter{
				Children: []Widget{
					Composite{
						Border: true,
						Layout: Grid{Columns: 10},
						Children: []Widget{
							Label{
								Name: "member_number",
								Text: "고객 번호",
							},
							LineEdit{
								Text: Bind("member_number"),
							},
							Label{
								Name: "member_name",
								Text: "고객 이름",
							},
							LineEdit{
								Text: Bind("member_name"),
							},
							Label{
								Name: "member_phone_number",
								Text: "휴대폰 번호",
							},
							LineEdit{
								Text: Bind("member_phone_number"),
							},
						},
					},
					Composite{
						Border: true,
						Layout: Grid{Columns: 10},
						Children: []Widget{
							Label{
								ColumnSpan: 2,
								Name:       "search_phone_number",
								Text:       "핸드폰 번호 조회",
							},
							LineEdit{
								Text: Bind("search_phone_number"),
							},
						},
					},
				},
			}, // end of VSplitter

			VSplitter{
				Children: []Widget{

					HSplitter{
						Children: []Widget{
							Label{
								ColumnSpan: 2,
								Name:       "search_phone_number",
								Text:       "핸드폰 번호",
							},
							LineEdit{
								Text: Bind("search_phone_number"),
							},
						},
					},
				},
			},
			VSplitter{
				Children: []Widget{
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
