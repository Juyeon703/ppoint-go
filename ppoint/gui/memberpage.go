package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"ppoint/dto"
	"sort"
	"strconv"
)

type MemberPage struct {
	*walk.Composite
}

func newMemberPage(parent walk.Container) (Page, error) {
	p := new(MemberPage)
	var tv *walk.TableView
	var memberIdLE, memberNameLE, phonenumLE, birthLE, pointLE, countLE, cdtLE, udtLE *walk.LineEdit
	model := NewMembersModel()

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "고객 관리",
		Layout:   VBox{},
		Border:   true,
		Children: []Widget{
			Label{Text: "고객 관리 페이지"},
			Composite{
				Layout: VBox{Margins: Margins{150, 10, 150, 10}},
				Children: []Widget{
					PushButton{
						Text: "신규 고객 등록",
						OnClicked: func() {
							member := new(MDto)
							if cmd, err := RunMemberEditDialog(winMain, member); err != nil {
								fmt.Println("====err=====")
								log.Print(err)
							} else if cmd == walk.DlgCmdOK {
								fmt.Println("====수정 완료=====")
								//outTE.SetText(fmt.Sprintf("%+v", animal))
							}
						},
					},
					Label{
						Text: "조회 고객",
					},
					Composite{
						Layout: Grid{Columns: 4},
						Border: true,
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
										Text: "보유 포인트 : ",
									},
									LineEdit{
										AssignTo: &pointLE,
										ReadOnly: true,
									},
								},
							},
							Composite{
								Layout: HBox{},
								Children: []Widget{
									Label{
										Text: "방문 횟수 : ",
									},
									LineEdit{
										AssignTo: &countLE,
										ReadOnly: true,
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
									},
								},
							},
							PushButton{
								Text: "수정",
								OnClicked: func() {
									member := new(MDto)
									if cmd, err := RunMemberEditDialog(winMain, member); err != nil {
										fmt.Println("====err=====")
										log.Print(err)
									} else if cmd == walk.DlgCmdOK {
										fmt.Println("====수정 완료=====")
										//outTE.SetText(fmt.Sprintf("%+v", animal))
									}
								},
							},
							PushButton{
								Text: "결제",
								OnClicked: func() {

								},
							},
							PushButton{
								Text: "매출 이력 조회",
								OnClicked: func() {

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
					LineEdit{},
					PushButton{
						Text: "검색",
						OnClicked: func() {

						},
					},
					HSpacer{
						Size: 500,
					},
					Label{
						Text: "검색 수 : " + strconv.Itoa(model.RowCount()),
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
					{Title: "등급", DataMember: "GradeName"},
					{Title: "이름", DataMember: "MemberName"},
					{Title: "핸드폰번호", DataMember: "PhoneNumber"},
					{Title: "생일", DataMember: "Birth"},
					{Title: "보유포인트", DataMember: "TotalPoint", Alignment: AlignFar},
					{Title: "가입일", DataMember: "CreateDate", Width: 150},
					{Title: "최근방문일", DataMember: "UpdateDate", Width: 150},
				},
				Model: model,
				OnSelectedIndexesChanged: func() {
					// 에러..?
					fmt.Printf("SelectedIndexes: %v\n", tv.SelectedIndexes())
					var index []int
					index = tv.SelectedIndexes()
					fmt.Println(model.Value(index[0], 0), model.Value(index[0], 2))
					memberIdLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 0)))
					memberNameLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 1)))
					phonenumLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 2)))
					birthLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 3)))
					pointLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 4)))
					countLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 5)))
					cdtLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 6)))
					udtLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 7)))
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

type MembersModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	members    []*dto.MemberDto
}

func NewMembersModel() *MembersModel {
	m := new(MembersModel)
	m.ResetRows()
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

func (m *MembersModel) ResetRows() {
	var err error
	var memberList []dto.MemberDto
	if memberList, err = dbconn.SelectMembersOrderByGrade(); err != nil {
		panic(err.Error())
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
