package gui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/dto"
	"ppoint/service"
	"ppoint/types"
	"ppoint/utils"
	"strconv"
	"strings"
)

type PointPage struct {
	*walk.Composite
}

var cashSV, cardSV, nPointLimit int

func newPointPage(parent walk.Container) (Page, error) {
	p := new(PointPage)
	var err error
	var ppdb, mudb *walk.DataBinder
	var searchMember, memberIdLE, memberNameLE, phoneNumberLE, birthLE, udtLE *walk.LineEdit
	var memberIdNE, pointNE, countNE, salesNE, subPointNE, beforePointNE, afterPointNE, fixedSalesNE, totalSalesNE, totalPointNE, addPointNE *walk.NumberEdit
	var payTypeRBtn *walk.GroupBox
	var radioCardBtn, radioCashBtn *walk.RadioButton
	var updateBtn, cancelBtn *walk.PushButton
	var clickedPT = types.Card
	var newMember *dto.MemberDto
	revenue := new(dto.RevenueAddDto)
	revenue = &dto.RevenueAddDto{PayType: types.Card}
	var updateMember = new(dto.MemberUpdateDto)
	const updateTitle = "수정"
	const okTitle = "확인"
	var nameTemp, phoneTemp, birthTemp string
	var pointTemp, countTemp float64

	if cashSV, err = service.FindSettingValue(dbconn, types.SettingCash); err != nil {
		log.Error(err.Error())
		panic(err)
	}
	if cardSV, err = service.FindSettingValue(dbconn, types.SettingCard); err != nil {
		log.Error(err.Error())
		panic(err)
	}
	if nPointLimit, err = service.FindSettingValue(dbconn, types.SettingPointLimit); err != nil {
		log.Error(err.Error())
		panic(err)
	}
	log.Debugf("현금 적립 %d, 카드 적립 %d, 포인트 사용 제한 : %dp", cashSV, cardSV, nPointLimit)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "포인트 관리",
		Layout:   VBox{},
		Border:   true,
		MinSize:  Size{subWidth, subHeight},
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
								OnEditingFinished: func() {
									str := searchMember.Text()
									if strings.HasPrefix(str, "010") {
										if len(str) == 11 || len(str) == 10 {
											searchMember.SetText(utils.PhoneNumAddHyphen(str))
										}
									}
								},
							},
							PushButton{
								Text: "검색",
								OnClicked: func() {
									if searchMember.Text() == "" {
										MsgBox("검색 에러", "검색어를 입력해주세요.")
										return
									} else {
										if memberList, err := service.FindMemberList(dbconn, searchMember.Text()); err != nil {
											log.Error(err.Error())
											panic(err)
										} else {
											if len(memberList) <= 0 {
												MsgBox("검색 결과 없음", "검색 결과가 없습니다.\n신규 회원을 등록해주세요.")
												addMember := new(dto.MemberAddDto)
												if cmd, err := RunMemberAddDialog(winMain, addMember); err != nil {
													log.Error(err)
												} else if cmd == walk.DlgCmdOK {
													memberInfoClear(memberIdLE, memberNameLE, phoneNumberLE, birthLE, udtLE,
														memberIdNE, pointNE, countNE, beforePointNE, afterPointNE, totalSalesNE, totalPointNE)
													clickedPT = revenueInfoClear(salesNE, subPointNE, fixedSalesNE, addPointNE, radioCardBtn, radioCashBtn)
													if newMember, err = service.FindMember(dbconn, addMember.MemberName, addMember.PhoneNumber); err != nil {
														log.Debug(err.Error())
														panic(err)
													}
													memberIdNE.SetValue(float64(newMember.MemberId))
													memberIdLE.SetText(strconv.Itoa(newMember.MemberId))
													memberNameLE.SetText(newMember.MemberName)
													phoneNumberLE.SetText(newMember.PhoneNumber)
													birthLE.SetText(newMember.Birth)
													udtLE.SetText(newMember.UpdateDate)
												}
												searchMember.SetText("")
											} else { // 검색 결과 있을 시 새로운 창 호출
												if cmd, err := RunMemberSearchDialog(winMain, memberList, memberIdLE, memberNameLE,
													phoneNumberLE, birthLE, udtLE, memberIdNE, pointNE, countNE, beforePointNE,
													afterPointNE, totalSalesNE, totalPointNE); err != nil {
													log.Error(err)
												} else if cmd == walk.DlgCmdOK {
													searchMember.SetText("")
												}
											}
										}
									}
								},
							},
							PushButton{
								Text: "초기화",
								OnClicked: func() {
									searchMember.SetText("")
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
						MaxSize: Size{Width: (subWidth / 2) - 100, Height: subHeight},
						DataBinder: DataBinder{
							AssignTo:   &mudb,
							Name:       "updateMember",
							DataSource: updateMember,
							//ErrorPresenter: ToolTipErrorPresenter{},
						},
						Children: []Widget{
							Composite{
								Layout: HBox{},
								Children: []Widget{
									HSpacer{MaxSize: Size{
										Width: subWidth / 2,
									}},
									Label{
										Text:       "[  조회 고객  ]",
										ColumnSpan: 10,
										Font:       Font{Bold: true},
									},
									HSpacer{MaxSize: Size{
										Width: subWidth / 2,
									}},
									PushButton{
										Text:    "초기화",
										MaxSize: Size{Width: 50},
										OnClicked: func() {
											if updateBtn.Text() == okTitle {
												MsgBox("초기화 오류", "수정 중엔 초기화할 수 없습니다.")
											} else {
												memberInfoClear(memberIdLE, memberNameLE, phoneNumberLE, birthLE, udtLE,
													memberIdNE, pointNE, countNE, beforePointNE, afterPointNE, totalSalesNE, totalPointNE)
											}
										},
									},
								},
							},
							LineEdit{
								AssignTo: &memberIdLE,
								Visible:  false,
								Text:     Bind("MemberId", SelRequired{}),
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
										Text:     Bind("MemberName", SelRequired{}),
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
										Text:     Bind("PhoneNumber", Regexp{Pattern: "^01([0|1|6|7|8|9])-([0-9]{3,4})-([0-9]{4})|01([0|1|6|7|8|9])([0-9]{3,4})([0-9]{4})$"}, SelRequired{}),
										OnEditingFinished: func() {
											phoneNumberLE.SetText(utils.PhoneNumAddHyphen(phoneNumberLE.Text()))
										},
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
										Text:     Bind("Birth", Regexp{Pattern: "(19[0-9][0-9]|20[0-9][0-9])-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$"}),
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "누적 매출 금액 : ",
									},
									NumberEdit{
										AssignTo: &totalSalesNE,
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
										AssignTo: &totalPointNE,
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
										Value:    Bind("TotalPoint"),
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
										Value:    Bind("VisitCount"),
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
										AssignTo: &updateBtn,
										Text:     updateTitle,
										OnClicked: func() {
											if memberIdLE.Text() != "" {
												if updateBtn.Text() == updateTitle {
													nameTemp = memberNameLE.Text()
													phoneTemp = phoneNumberLE.Text()
													birthTemp = birthLE.Text()
													pointTemp = pointNE.Value()
													countTemp = countNE.Value()
													memberNameLE.SetReadOnly(false)
													phoneNumberLE.SetReadOnly(false)
													birthLE.SetReadOnly(false)
													pointNE.SetReadOnly(false)
													countNE.SetReadOnly(false)
													updateBtn.SetText(okTitle)
													cancelBtn.SetVisible(true)
												} else if updateBtn.Text() == okTitle && !memberNameLE.ReadOnly() {
													if err := mudb.Submit(); err != nil {
														log.Error(err.Error())
														panic(err)
													}
													if nameTemp != updateMember.MemberName || phoneTemp != updateMember.PhoneNumber || birthTemp != updateMember.Birth ||
														int(pointTemp) != updateMember.TotalPoint || int(countTemp) != updateMember.VisitCount {

														if existMember, err := service.FindUpdateMemberPhoneNumber(dbconn, updateMember.PhoneNumber, memberIdLE.Text()); existMember != nil {
															if err != nil {
																log.Errorf("(고객 수정) >>> 중복 조회 실패 >>> [%v]", err)
															} else {
																MsgBox("알림", "이미 존재하는 핸드폰 번호 입니다.")
															}
														} else {
															if err := service.MemberUpdate(dbconn, updateMember, int(pointTemp)); err != nil {
																MsgBox("알림", "고객 등록에 실패하였습니다.")
																log.Error(err)
															} else {
																udtLE.SetText(utils.CurrentTime())
																memberNameLE.SetReadOnly(true)
																phoneNumberLE.SetReadOnly(true)
																birthLE.SetReadOnly(true)
																pointNE.SetReadOnly(true)
																countNE.SetReadOnly(true)
																updateBtn.SetText(updateTitle)
																cancelBtn.SetVisible(false)
																MsgBox("수정 완료", "회원 정보가 변경되었습니다.")
																beforePointNE.SetValue(float64(updateMember.TotalPoint))
																afterPointNE.SetValue(float64(updateMember.TotalPoint))
																clickedPT = revenueInfoClear(salesNE, subPointNE, fixedSalesNE, addPointNE, radioCardBtn, radioCashBtn)
															}
														}
													}
												}
											}
										},
									},
									PushButton{
										AssignTo: &cancelBtn,
										Text:     "취소",
										Visible:  false,
										OnClicked: func() {
											if cancelBtn.Visible() {
												memberNameLE.SetText(nameTemp)
												phoneNumberLE.SetText(phoneTemp)
												birthLE.SetText(birthTemp)
												pointNE.SetValue(pointTemp)
												countNE.SetValue(countTemp)
												memberNameLE.SetReadOnly(true)
												phoneNumberLE.SetReadOnly(true)
												birthLE.SetReadOnly(true)
												pointNE.SetReadOnly(true)
												countNE.SetReadOnly(true)
												updateBtn.SetText(updateTitle)
												cancelBtn.SetVisible(false)
											}
										},
									},
								},
							},
						},
					},
					Composite{
						MaxSize: Size{Width: 1, Height: subHeight},
						Layout:  VBox{},
						Children: []Widget{
							Label{
								Visible: false,
							},
						},
					},
					Composite{
						Border:  true,
						Layout:  VBox{},
						MaxSize: Size{Width: subWidth, Height: subHeight},
						DataBinder: DataBinder{
							AssignTo:   &ppdb,
							Name:       "revenue",
							DataSource: revenue,
							//ErrorPresenter: ToolTipErrorPresenter{},
						},
						Children: []Widget{
							Composite{
								Layout: HBox{},
								Children: []Widget{
									HSpacer{MaxSize: Size{
										Width: subWidth / 4,
									}},
									Label{
										Text:       "[  결제  ]",
										ColumnSpan: 10,
										Font:       Font{Bold: true},
									},
									HSpacer{MaxSize: Size{
										Width: subWidth / 2,
									}},
									PushButton{
										Text:    "초기화",
										MaxSize: Size{Width: 50},
										OnClicked: func() {
											clickedPT = revenueInfoClear(salesNE, subPointNE, fixedSalesNE, addPointNE, radioCardBtn, radioCashBtn)
										},
									},
								},
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
										Text: "매출 금액 : ",
									},
									NumberEdit{
										AssignTo: &salesNE,
										Value:    Bind("Sales", SelRequired{}),
										Suffix:   " 원",
										OnValueChanged: func() {
											revenueInfoCalc(clickedPT, salesNE, subPointNE, fixedSalesNE, addPointNE, beforePointNE, afterPointNE)
										},
									},
								},
							},
							RadioButtonGroupBox{
								ColumnSpan: 2,
								Title:      "결제 방법",
								Layout:     HBox{},
								DataMember: "PayType",
								AssignTo:   &payTypeRBtn,
								OnBoundsChanged: func() {
									radioCardBtn.SetChecked(true)
								},
								Buttons: []RadioButton{
									{Text: "카드", AssignTo: &radioCardBtn, Value: types.Card, OnClicked: func() {
										clickedPT = types.Card
										revenueInfoCalc(clickedPT, salesNE, subPointNE, fixedSalesNE, addPointNE, beforePointNE, afterPointNE)
									}},
									{Text: "현금", AssignTo: &radioCashBtn, Value: types.Cash, OnClicked: func() {
										clickedPT = types.Cash
										revenueInfoCalc(clickedPT, salesNE, subPointNE, fixedSalesNE, addPointNE, beforePointNE, afterPointNE)
									}},
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
										Value:    Bind("SubPoint"),
										Suffix:   " p",
										OnValueChanged: func() {
											revenueInfoCalc(clickedPT, salesNE, subPointNE, fixedSalesNE, addPointNE, beforePointNE, afterPointNE)
										},
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
										AssignTo: &beforePointNE,
										ReadOnly: true,
										Suffix:   " p",
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "적립 포인트 : ",
									},
									NumberEdit{
										AssignTo: &addPointNE,
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
										AssignTo: &afterPointNE,
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
										AssignTo: &fixedSalesNE,
										ReadOnly: true,
										Suffix:   " 원",
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									VSpacer{
										Size: 21,
									},
								},
							},
							PushButton{
								Text: "확인/적립",
								OnClicked: func() {
									if memberIdLE.Text() == "" {
										MsgBox("선택된 회원 없음", "선택된 회원이 없습니다.")
									} else {
										if err := ppdb.Submit(); err != nil {
											log.Error(err)
											return
										}

										//if int(pointNE.Value()) <= 0 {
										//	MsgBox("알림", "사용가능한 보유 포인트가 없습니다.")
										//	return
										//}

										//if int(pointNE.Value()) < nPointLimit {
										//	MsgBox("알림", "사용가능한 포인트가 "+strconv.Itoa(nPointLimit)+"p 보다 많아야 합니다.")
										//	return
										//}

										if revenue.Sales <= 0 {
											MsgBox("알림", "매출 금액을 입력해주세요.")
											return
										}

										if revenue.SubPoint > int(pointNE.Value()) {
											MsgBox("알림", "보유 포인트가 부족합니다.")
											return
										}

										if revenue.Sales <= 0 || revenue.SubPoint > int(pointNE.Value()) {
											MsgBox("error", "error") //////////////////////////////////////////////////////////////////////////
										} else {
											if err := service.RevenueAdd(dbconn, revenue); err != nil {
												log.Error(err.Error())
												panic(err)
											}
											totalSalesNE.SetValue(totalSalesNE.Value() + salesNE.Value())
											totalPointNE.SetValue(totalPointNE.Value() + addPointNE.Value())
											pointNE.SetValue(afterPointNE.Value())
											countNE.SetValue(countNE.Value() + 1)
											beforePointNE.SetValue(afterPointNE.Value())
											udtLE.SetText(utils.CurrentTime())

											if int(subPointNE.Value()) == 0 {
												MsgBox("결제 완료", "포인트 "+strconv.Itoa(int(addPointNE.Value()))+"이 적립되었습니다.")
											} else {
												MsgBox("결제 완료", "포인트 "+strconv.Itoa(int(subPointNE.Value()))+"이 사용되었습니다.")
											}
											clickedPT = revenueInfoClear(salesNE, subPointNE, fixedSalesNE, addPointNE, radioCardBtn, radioCashBtn)
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

func memberInfoClear(memberIdLE, memberNameLE, phoneNumberLE, birthLE, udtLE *walk.LineEdit,
	memberIdNE, pointNE, countNE, beforePointNE, afterPointNE, totalSalesNE, totalPointNE *walk.NumberEdit) {
	memberIdLE.SetText("")
	memberNameLE.SetText("")
	phoneNumberLE.SetText("")
	birthLE.SetText("")
	udtLE.SetText("")
	pointNE.SetValue(0)
	countNE.SetValue(0)
	memberIdNE.SetValue(0)
	beforePointNE.SetValue(0)
	afterPointNE.SetValue(0)
	totalSalesNE.SetValue(0)
	totalPointNE.SetValue(0)
}

func revenueInfoClear(salesNE, subPointNE, fixedSalesNE, addPointNE *walk.NumberEdit, radioCardBtn, radioCashBtn *walk.RadioButton) string {
	salesNE.SetValue(0)
	subPointNE.SetValue(0)
	fixedSalesNE.SetValue(0)
	addPointNE.SetValue(0)
	radioCardBtn.SetChecked(true)
	radioCashBtn.SetChecked(false)
	return types.Card
}

func revenueInfoCalc(clickedPT string, salesNE, subPointNE, fixedSalesNE, addPointNE, beforePointNE, afterPointNE *walk.NumberEdit) {
	creditValue := cardSV
	if clickedPT == types.Cash {
		creditValue = cashSV
	}
	fixedSalesNE.SetValue(salesNE.Value() - subPointNE.Value())
	if subPointNE.Value() == 0 {
		addPointNE.SetValue(salesNE.Value() * float64(creditValue) / 100)
		afterPointNE.SetValue(beforePointNE.Value() + (salesNE.Value() * float64(creditValue) / 100))
	} else {
		addPointNE.SetValue(0)
		afterPointNE.SetValue(beforePointNE.Value() - subPointNE.Value())
	}
}
