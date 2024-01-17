package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/dto"
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
							var memberList []dto.MemberDto
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
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Composite{
						Border: true,
						Layout: VBox{},
						Children: []Widget{
							Label{
								Text: "조회 고객",
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "고객 번호 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "고객 이름 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "핸드폰 번호 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "누적 결제 금액 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "누적 적립 포인트 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "누적 사용 포인트 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "보유 포인트 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "방문 횟수: ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "최근 방문일 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									PushButton{
										Text: "등록",
										OnClicked: func() {

										},
									},
									PushButton{
										Text: "수정",
										OnClicked: func() {

										},
									},
								},
							},
						},
					},
					Composite{
						Border: true,
						Layout: VBox{},
						Children: []Widget{
							Label{
								Text: "결제",
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "상품 금액 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "결제 방법 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "사용 포인트 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "적립 전 포인트 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "적립 후 포인트 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "결제 금액 : ",
									},
									LineEdit{
										//AssignTo: &,
										ReadOnly: true,
									},
								},
							},
							PushButton{
								Text: "확인/적립",
								OnClicked: func() {

								},
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
