package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/dto"
	"ppoint/service"
	"ppoint/utils"
	"sort"
	"strconv"
	"strings"
)

type MemberPage struct {
	*walk.Composite
}

var moveId int

func newMemberPage(parent walk.Container) (Page, error) {
	p := new(MemberPage)
	var tv *walk.TableView
	var mudb *walk.DataBinder
	var memberIdLE, memberNameLE, phonenumLE, birthLE, cdtLE, udtLE, mpSearchLE *walk.LineEdit
	var pointNE, countNE *walk.NumberEdit
	var tvResultLabel *walk.Label
	var mpSearchBtn, updateBtn, selectBtn *walk.PushButton
	var updateMember = new(dto.MemberUpdateDto)
	model := NewMembersModel("")
	var isExistMember = new(dto.IsExistMember)
	const updateTitle = "수정"
	const okTitle = "확인"
	const cancelTitle = "취소"
	const selectTitle = "매출 이력 조회"

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "고객 관리",
		Layout:   VBox{},
		Border:   true,
		MinSize:  Size{subWidth, subHeight},
		Children: []Widget{
			Label{Text: "고객 관리 페이지"},
			Composite{
				Layout: VBox{Margins: Margins{130, 10, 130, 10}},
				Children: []Widget{
					PushButton{
						Text: "신규 고객 등록",
						OnClicked: func() {
							addMember := new(dto.MemberAddDto)
							if cmd, err := RunMemberAddDialog(winMain, addMember); err != nil {
								log.Error(err)
							} else if cmd == walk.DlgCmdOK {
								model = tvReloading("", tv, tvResultLabel)
								tv.SetCurrentIndex(model.RowCount() - 1)
							}
						},
					},
					Label{
						Text: "조회 고객",
					},
					Composite{
						Layout: Grid{Columns: 4},
						Border: true,
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
									Label{
										Text: "번호 : ",
									},
									LineEdit{
										AssignTo: &memberIdLE,
										ReadOnly: true,
										Text:     Bind("MemberId", SelRequired{}),
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "이름 : ",
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
										AssignTo: &phonenumLE,
										ReadOnly: true,
										Text:     Bind("PhoneNumber", Regexp{Pattern: "^01([0|1|6|7|8|9])-([0-9]{3,4})-([0-9]{4})|01([0|1|6|7|8|9])([0-9]{3,4})([0-9]{4})$"}, SelRequired{}),
										OnEditingFinished: func() {
											phonenumLE.SetText(utils.PhoneNumAddHyphen(phonenumLE.Text()))
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
										Text: "보유 포인트 : ",
									},
									NumberEdit{
										AssignTo: &pointNE,
										ReadOnly: true,
										Suffix:   " p",
										Value:    Bind("TotalPoint"),
										MinSize:  Size{Width: 60},
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
										MinSize:  Size{Width: 40},
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "가입일 : ",
									},
									LineEdit{
										AssignTo: &cdtLE,
										ReadOnly: true,
										MinSize:  Size{Width: 100},
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "수정일 : ",
									},
									LineEdit{
										AssignTo: &udtLE,
										ReadOnly: true,
										MinSize:  Size{Width: 100},
									},
								},
							},
							PushButton{
								Text:     updateTitle,
								AssignTo: &updateBtn,
								OnClicked: func() {
									if memberIdLE.Text() != "" {
										if updateBtn.Text() == updateTitle {
											memberNameLE.SetReadOnly(false)
											phonenumLE.SetReadOnly(false)
											birthLE.SetReadOnly(false)
											pointNE.SetReadOnly(false)
											countNE.SetReadOnly(false)
											updateBtn.SetText(okTitle)
											selectBtn.SetText(cancelTitle)
										} else if updateBtn.Text() == okTitle && !memberNameLE.ReadOnly() {
											log.Debug("(고객 관리 페이지) >>> 고객 정보 수정 START")
											if err := mudb.Submit(); err != nil {
												log.Error(err.Error())
												panic(err)
											}
											if isExistMember.MemberName != updateMember.MemberName || isExistMember.PhoneNumber != updateMember.PhoneNumber ||
												isExistMember.Birth != updateMember.Birth || isExistMember.TotalPoint != updateMember.TotalPoint ||
												isExistMember.VisitCount != updateMember.VisitCount {

												if existMember, err := service.FindUpdateMemberPhoneNumber(dbconn, updateMember.PhoneNumber, memberIdLE.Text()); existMember != nil {
													if err != nil {
														log.Errorf("(고객 수정) >>> 중복 조회 실패 >>> [%v]", err)
													} else {
														MsgBox("알림", "이미 존재하는 핸드폰 번호 입니다.")
													}
												} else {
													if err := service.MemberUpdate(dbconn, updateMember, isExistMember.TotalPoint); err != nil {
														MsgBox("알림", "고객 등록에 실패하였습니다.")
														log.Error(err)
													} else {
														isExistMember.MemberName = updateMember.MemberName
														isExistMember.PhoneNumber = updateMember.PhoneNumber
														isExistMember.Birth = updateMember.Birth
														isExistMember.TotalPoint = updateMember.TotalPoint
														isExistMember.VisitCount = updateMember.VisitCount
														isExistMember.UpdateDate = utils.CurrentTime()
														udtLE.SetText(utils.CurrentTime())
														memberNameLE.SetReadOnly(true)
														phonenumLE.SetReadOnly(true)
														birthLE.SetReadOnly(true)
														pointNE.SetReadOnly(true)
														countNE.SetReadOnly(true)
														updateBtn.SetText(updateTitle)
														selectBtn.SetText(selectTitle)
														model = tvReloading("", tv, tvResultLabel)

														log.Debugf("(고객 관리 페이지) >>> 고객 정보 수정 : [%v]", updateMember)
														MsgBox("수정 완료", "["+updateMember.MemberName+"] 회원 정보가 변경되었습니다.")
														log.Debug("(고객 관리 페이지) >>> 고객 정보 수정 END")
													}
												}
											}

										}
									}
								},
							},
							PushButton{
								AssignTo: &selectBtn,
								Text:     selectTitle,
								OnClicked: func() {
									if selectBtn.Text() == selectTitle {
										if memberIdLE.Text() != "" {
											moveId, _ = strconv.Atoi(isExistMember.MemberId)
											winMain.pageActions[1].SetChecked(false)
											if err := multiPageMainWindow.setCurrentAction(winMain.pageActions[2]); err != nil {
												log.Error(err.Error())
												panic(err)
											}
											moveId = 0
										}

									} else if selectBtn.Text() == cancelTitle {
										memberInfoReloading(isExistMember, memberIdLE, memberNameLE, phonenumLE, birthLE, cdtLE, udtLE, pointNE, countNE)
										memberNameLE.SetReadOnly(true)
										phonenumLE.SetReadOnly(true)
										birthLE.SetReadOnly(true)
										pointNE.SetReadOnly(true)
										countNE.SetReadOnly(true)
										updateBtn.SetText(updateTitle)
										selectBtn.SetText(selectTitle)
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
					Label{
						Text: "검색 : ",
					},
					LineEdit{
						AssignTo: &mpSearchLE,
						OnEditingFinished: func() {
							str := mpSearchLE.Text()
							if strings.HasPrefix(str, "010") {
								if len(str) == 11 || len(str) == 10 {
									mpSearchLE.SetText(utils.PhoneNumAddHyphen(str))
								}
							}
						},
					},
					PushButton{
						AssignTo: &mpSearchBtn,
						Text:     "검색",
						OnClicked: func() {
							if mpSearchLE.Text() != "" {
								// 이름, 폰 번호
								model = tvReloading(mpSearchLE.Text(), tv, tvResultLabel)
							}
						},
					},
					PushButton{
						Text: "초기화",
						OnClicked: func() {
							if mpSearchLE.Text() != "" {
								mpSearchLE.SetText("")
								model = tvReloading("", tv, tvResultLabel)
							}
						},
					},
					HSpacer{
						Size: 500,
					},
					Label{
						Text: "검색 수 : ",
					},
					Label{
						Text:     strconv.Itoa(model.RowCount()),
						AssignTo: &tvResultLabel,
					},
				},
			},
			TableView{
				Name:             "memberTable",
				AssignTo:         &tv,
				AlternatingRowBG: true,
				ColumnsOrderable: true,
				MultiSelection:   true,
				MinSize:          Size{300, 300},
				Columns: []TableViewColumn{
					{Title: "번호", DataMember: "MemberId", Hidden: true},
					{Title: "이름", DataMember: "MemberName"},
					{Title: "핸드폰번호", DataMember: "PhoneNumber"},
					{Title: "생일", DataMember: "Birth"},
					{Title: "보유포인트", DataMember: "TotalPoint", Alignment: AlignFar, FormatFunc: func(value interface{}) string {
						r, _ := utils.InterfaceToInt64(value)
						return utils.MoneyConverter(r)
					}},
					{Title: "방문 횟수", DataMember: "VisitCount", Alignment: AlignFar, FormatFunc: func(value interface{}) string {
						r, _ := utils.InterfaceToInt64(value)
						return utils.MoneyConverter(r)
					}},
					{Title: "가입일", DataMember: "CreateDate", Width: 150},
					{Title: "최근방문일", DataMember: "UpdateDate", Width: 150},
				},
				Model: model,
				OnSelectedIndexesChanged: func() {
					// 에러..?
					index := tv.SelectedIndexes()
					if len(index) > 0 {
						isExistMember.MemberId = fmt.Sprintf("%v", model.Value(index[0], 0))
						isExistMember.MemberName = fmt.Sprintf("%v", model.Value(index[0], 1))
						isExistMember.PhoneNumber = fmt.Sprintf("%v", model.Value(index[0], 2))
						isExistMember.Birth = fmt.Sprintf("%v", model.Value(index[0], 3))
						pointTemp, _ := strconv.Atoi(fmt.Sprintf("%v", model.Value(index[0], 4)))
						isExistMember.TotalPoint = pointTemp
						countTemp, _ := strconv.Atoi(fmt.Sprintf("%v", model.Value(index[0], 5)))
						isExistMember.VisitCount = countTemp
						isExistMember.CreateDate = fmt.Sprintf("%v", model.Value(index[0], 6))
						isExistMember.UpdateDate = fmt.Sprintf("%v", model.Value(index[0], 7))
						memberInfoReloading(isExistMember, memberIdLE, memberNameLE, phonenumLE, birthLE, cdtLE, udtLE, pointNE, countNE)
					}
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

func memberInfoReloading(isExistMember *dto.IsExistMember, memberIdLE, memberNameLE, phonenumLE, birthLE,
	cdtLE, udtLE *walk.LineEdit, pointNE, countNE *walk.NumberEdit) {
	memberIdLE.SetText(isExistMember.MemberId)
	memberNameLE.SetText(isExistMember.MemberName)
	phonenumLE.SetText(isExistMember.PhoneNumber)
	birthLE.SetText(isExistMember.Birth)
	pointNE.SetValue(float64(isExistMember.TotalPoint))
	countNE.SetValue(float64(isExistMember.VisitCount))
	cdtLE.SetText(isExistMember.CreateDate)
	udtLE.SetText(isExistMember.UpdateDate)
}

func tvReloading(search string, tv *walk.TableView, tvResultLabel *walk.Label) *MembersModel {
	model := NewMembersModel(search)
	tv.SetModel(model)
	tvResultLabel.SetText(strconv.Itoa(model.RowCount()))
	return model
}

type MembersModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	members    []*dto.MemberDto
}

func NewMembersModel(search string) *MembersModel {
	m := new(MembersModel)
	m.ResetRows(search)
	return m
}

func (m *MembersModel) RowCount() int {
	return len(m.members)
}

func (m *MembersModel) Value(row, col int) interface{} {
	member := m.members[row]

	switch col {
	case 0:
		return member.MemberId

	case 1:
		return member.MemberName

	case 2:
		return member.PhoneNumber

	case 3:
		return member.Birth

	case 4:
		return member.TotalPoint

	case 5:
		return member.VisitCount

	case 6:
		return member.CreateDate

	case 7:
		return member.UpdateDate
	}

	panic("unexpected col")
}

func (m *MembersModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.SliceStable(m.members, func(i, j int) bool {
		a, b := m.members[i], m.members[j]

		c := func(ls bool) bool {
			if m.sortOrder == walk.SortAscending {
				return ls
			}

			return !ls
		}

		switch m.sortColumn {
		case 0:
			return c(a.MemberId < b.MemberId)

		case 1:
			return c(a.MemberName < b.MemberName)

		case 2:
			return c(a.PhoneNumber < b.PhoneNumber)

		case 3:
			return c(a.Birth < b.Birth)

		case 4:
			return c(a.TotalPoint < b.TotalPoint)

		case 5:
			return c(a.VisitCount < b.VisitCount)

		case 6:
			return c(a.CreateDate < b.CreateDate)

		case 7:
			return c(a.UpdateDate < b.UpdateDate)
		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}

func (m *MembersModel) ResetRows(search string) {
	var err error
	var memberList []dto.MemberDto

	if memberList, err = service.FindMemberList(dbconn, search); err != nil {
		log.Error(err.Error())
		panic(err)
	}
	if search != "" && len(memberList) <= 0 {
		MsgBox("검색 에러", "검색 결과가 없습니다.")
	}
	m.members = make([]*dto.MemberDto, len(memberList))
	for i := range memberList {
		m.members[i] = &dto.MemberDto{
			MemberId:    memberList[i].MemberId,
			MemberName:  memberList[i].MemberName,
			PhoneNumber: memberList[i].PhoneNumber,
			Birth:       memberList[i].Birth,
			TotalPoint:  memberList[i].TotalPoint,
			VisitCount:  memberList[i].VisitCount,
			CreateDate:  memberList[i].CreateDate,
			UpdateDate:  memberList[i].UpdateDate,
		}
	}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}
