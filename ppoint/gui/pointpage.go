package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"ppoint/dto"
	"ppoint/service"
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
	var ppdb *walk.DataBinder
	var searchMember, memberNameLE, phoneNumberLE, birthLE, udtLE *walk.LineEdit
	var memberIdNE, pointNE, countNE, salesNE, subPointNE *walk.NumberEdit
	var payTypeRBtn *walk.GroupBox
	revenue := new(dto.RevenueAddDto)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "포인트 관리",
		Layout:   VBox{},
		Border:   true,
		Children: []Widget{
			Label{Text: "포인트 관리 페이지"},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Composite{
						Border:  true,
						Layout:  HBox{},
						MaxSize: Size{Width: 500},
						Children: []Widget{
							Label{
								Text: "고객 조회 : ",
							},
							LineEdit{
								AssignTo: &searchMember,
								Text:     Bind("searchMember"),
							},
							PushButton{
								Text: "검색",
								OnClicked: func() {
									if searchMember.Text() == "" {
										MsgBox("검색 에러", "검색어를 입력해주세요.")
										return
									} else {
										if memberList, err := dbconn.SelectMemberSearch(searchMember.Text()); err != nil {
											panic(err.Error())
										} else {
											fmt.Println("=============> SelectMemberSearch() 호출")
											if len(memberList) <= 0 {
												MsgBox("검색 결과 없음", "검색 결과가 없습니다.\n신규 회원을 등록해주세요.")
												searchMember.SetText("")

												// 신규 회원 등록====================================== 구현 중

											} else { // 검색 결과 있을 시 새로운 창 호출
												if cmd, err := RunMemberSearchDialog(winMain, memberList, memberNameLE, phoneNumberLE, birthLE, udtLE, memberIdNE, pointNE, countNE); err != nil {
													log.Print(err)
												} else if cmd == walk.DlgCmdOK {
													fmt.Println("====회원 검색 창 호출 확인 버튼=====")
													fmt.Println("======>", memberIdNE.Value())
													searchMember.SetText("")
													// 누적 매출 같은거 가져오기 ============================= 구현 중
												}
											}
										}
									}
								},
							},
						},
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Composite{
						Border:  true,
						Layout:  VBox{},
						MaxSize: Size{Width: 300, Height: 1000},
						Children: []Widget{
							Label{
								Text: "조회 고객",
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "고객 이름 : ",
									},
									LineEdit{
										AssignTo: &memberNameLE,
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
										AssignTo: &phoneNumberLE,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "생일 : ",
									},
									LineEdit{
										AssignTo: &birthLE,
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
									NumberEdit{
										//AssignTo: &,
										ReadOnly: true,
										Suffix:   " 원",
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "누적 적립 포인트 : ",
									},
									NumberEdit{
										//AssignTo: &,
										ReadOnly: true,
										Suffix:   " p",
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "누적 사용 포인트 : ",
									},
									NumberEdit{
										//AssignTo: &,
										ReadOnly: true,
										Suffix:   " p",
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "보유 포인트 : ",
									},
									NumberEdit{
										AssignTo: &pointNE,
										ReadOnly: true,
										Suffix:   " p",
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "방문 횟수: ",
									},
									NumberEdit{
										AssignTo: &countNE,
										ReadOnly: true,
										Suffix:   " 회",
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
										AssignTo: &udtLE,
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
						Border:  true,
						Layout:  VBox{},
						MaxSize: Size{Width: 300, Height: 1000},
						DataBinder: DataBinder{
							AssignTo:       &ppdb,
							Name:           "revenue",
							DataSource:     revenue,
							ErrorPresenter: ToolTipErrorPresenter{},
						},
						Children: []Widget{
							Label{
								Text: "결제",
							},
							NumberEdit{
								AssignTo: &memberIdNE,
								Visible:  false,
								Value:    Bind("MemberId"),
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "상품 금액 : ",
									},
									NumberEdit{
										AssignTo: &salesNE,
										Value:    Bind("Sales", SelRequired{}),
										Suffix:   " 원",
									},
								},
							},
							RadioButtonGroupBox{
								ColumnSpan: 2,
								Title:      "결제 방법",
								Layout:     HBox{},
								DataMember: "PayType",
								AssignTo:   &payTypeRBtn,
								Buttons: []RadioButton{
									{Text: "카드", Value: "카드"},
									{Text: "현금", Value: "현금"},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "사용 포인트 : ",
									},
									NumberEdit{
										AssignTo: &subPointNE,
										Value:    Bind("SubPoint"), //Range{0, pointNE.Value()}
										Suffix:   " p",
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "적립 전 포인트 : ",
									},
									NumberEdit{
										//AssignTo: &,
										ReadOnly: true,
										Suffix:   " p",
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "적립 후 포인트 : ",
									},
									NumberEdit{
										//AssignTo: &,
										ReadOnly: true,
										Suffix:   " p",
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "결제 금액 : ",
									},
									NumberEdit{
										//AssignTo: &,
										ReadOnly: true,
										Suffix:   " 원",
									},
								},
							},
							PushButton{
								Text: "확인/적립",
								OnClicked: func() {
									if memberIdNE.Value() == 0 {
										MsgBox("선택된 회원 없음", "선택된 회원이 없습니다.")
									} else {
										if err := ppdb.Submit(); err != nil {
											log.Print(err)
											return
										}
										fmt.Println(revenue)
										if revenue.Sales > 0 {
											if err := service.RevenueAdd(dbconn, revenue); err != nil {
												panic(err)
											}
											memberInfoClear(memberNameLE, phoneNumberLE, birthLE, udtLE, memberIdNE, pointNE, countNE)
											revenueInfoClear(salesNE, subPointNE)
											// 라디오버튼 초기화..?
											// 고객 정보는 초기화 안하고 포인트 적립 후 정보 보여주면 좋을지도?
											MsgBox("결제 완료", "포인트 적립이 완료되었습니다.")
										}
									}
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

func memberInfoClear(memberNameLE, phoneNumberLE, birthLE, udtLE *walk.LineEdit, memberIdNE, pointNE, countNE *walk.NumberEdit) {
	memberNameLE.SetText("")
	phoneNumberLE.SetText("")
	birthLE.SetText("")
	udtLE.SetText("")
	memberIdNE.SetValue(0)
	pointNE.SetValue(0)
	countNE.SetValue(0)
}

func revenueInfoClear(salesNE, subPointNE *walk.NumberEdit) {
	salesNE.SetValue(0)
	subPointNE.SetValue(0)
}
