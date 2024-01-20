package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/dto"
	"ppoint/service"
	"sort"
	"strconv"
)

type SalesPage struct {
	*walk.Composite
}

func newSalesPage(parent walk.Container) (Page, error) {
	p := new(SalesPage)
	var tv *walk.TableView
	var startDate, endDate string
	var startDateSearchDE, endDateSearchDE *walk.DateEdit
	var sumNEcc, sumNEcard, sumNEcash, sumNEaddP, sumNEsubP *walk.NumberEdit
	startDate = "2017-01-01"
	endDate = "2024-01-31"
	model := NewRevenuesModel(startDate, endDate, 0)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "매출 관리",
		Layout:   VBox{},
		Border:   true,
		Children: []Widget{
			Label{Text: "매출 관리 페이지"},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "기간별 조회 :",
					},
					PushButton{
						Text: "전체",
						OnClicked: func() {

						},
					},
					PushButton{
						Text: "일별",
						OnClicked: func() {

						},
					},
					PushButton{
						Text: "월별",
						OnClicked: func() {

						},
					},
					PushButton{
						Text: "년도별",
						OnClicked: func() {

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
					DateEdit{
						AssignTo: &startDateSearchDE,
						Date:     Bind("startDate"),
						//OnDateChanged: func() {
						//	fmt.Println()
						//},
					},
					Label{
						Text: " ~ ",
					},
					DateEdit{
						AssignTo: &endDateSearchDE,
						Date:     Bind("endDate"),
					},
					PushButton{
						Text: "검색",
						OnClicked: func() {
							fmt.Println(startDateSearchDE.Date())
							fmt.Println(endDateSearchDE.Date())
						},
					},
					HSpacer{
						Size: 300,
					},
					Label{
						Text: "조회 기간 : " + startDate + " ~ " + endDate,
					},
					HSpacer{
						Size: 50,
					},
					Label{
						Text: "검색 수 : " + strconv.Itoa(model.RowCount()),
					},
				},
			},
			TableView{
				Name:             "revenueTable",
				AssignTo:         &tv,
				AlternatingRowBG: true,
				ColumnsOrderable: true,
				MultiSelection:   true,
				MinSize:          Size{300, 300},
				Columns: []TableViewColumn{
					{Title: "번호", DataMember: "RevenueId"},
					{Title: "고객번호", DataMember: "MemberId", Hidden: true},
					{Title: "이름", DataMember: "MemberName"},
					{Title: "핸드폰번호", DataMember: "PhoneNumber"},
					{Title: "결제금액", DataMember: "Sales", Alignment: AlignFar},
					{Title: "사용포인트", DataMember: "SubPoint", Alignment: AlignFar},
					{Title: "적립포인트", DataMember: "AddPoint", Alignment: AlignFar},
					{Title: "실제결제금액", DataMember: "FixedSales", Alignment: AlignFar},
					{Title: "결제방법", DataMember: "PayType"},
					{Title: "결제일", DataMember: "CreateDate", Width: 150},
				},
				Model: model,
				OnSelectedIndexesChanged: func() {
					var index []int
					index = tv.SelectedIndexes()
					fmt.Println(fmt.Sprintf("%v", model.Value(index[0], 0)))
				},
			},
			Composite{
				Layout: VBox{},
				Border: true,
				Children: []Widget{
					Label{
						Text: "통계",
					},
					Composite{
						Layout: HBox{},
						Children: []Widget{
							Label{
								Text: "총 매출 금액(현금+카드)",
							},
							NumberEdit{
								AssignTo: &sumNEcc,
								Suffix:   " 원",
							},
							Label{
								Text: "총 매출 금액(카드)",
							},
							NumberEdit{
								AssignTo: &sumNEcard,
								Suffix:   " 원",
							},
							Label{
								Text: "총 매출 금액(현금)",
							},
							NumberEdit{
								AssignTo: &sumNEcash,
								Suffix:   " 원",
							},
							Label{
								Text: "총 적립 포인트",
							},
							NumberEdit{
								AssignTo: &sumNEaddP,
								Suffix:   " p",
							},
							Label{
								Text: "총 사용 포인트",
							},
							NumberEdit{
								AssignTo: &sumNEsubP,
								Suffix:   " p",
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

type RevenuesModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	revenues   []*dto.RevenueDto
}

func NewRevenuesModel(startDate, endDate string, memberId int) *RevenuesModel {
	r := new(RevenuesModel)
	r.ResetRows(startDate, endDate, memberId)
	return r
}

func (r *RevenuesModel) RowCount() int {
	return len(r.revenues)
}

func (r *RevenuesModel) Value(row, col int) interface{} {
	revenue := r.revenues[row]

	switch col {
	case 0:
		return revenue.RevenueId

	case 1:
		return revenue.MemberId

	case 2:
		return revenue.MemberName

	case 3:
		return revenue.PhoneNumber

	case 4:
		return revenue.Sales

	case 5:
		return revenue.SubPoint

	case 6:
		return revenue.AddPoint

	case 7:
		return revenue.FixedSales

	case 8:
		return revenue.PayType

	case 9:
		return revenue.CreateDate
	}

	panic("unexpected col")
}

func (r *RevenuesModel) Sort(col int, order walk.SortOrder) error {
	r.sortColumn, r.sortOrder = col, order

	sort.SliceStable(r.revenues, func(i, j int) bool {
		a, b := r.revenues[i], r.revenues[j]

		c := func(ls bool) bool {
			if r.sortOrder == walk.SortAscending {
				return ls
			}

			return !ls
		}

		switch r.sortColumn {
		case 0:
			return c(a.RevenueId < b.RevenueId)

		case 1:
			return c(a.MemberId < b.MemberId)

		case 2:
			return c(a.MemberName < b.MemberName)

		case 3:
			return c(a.PhoneNumber < b.PhoneNumber)

		case 4:
			return c(a.Sales < b.Sales)

		case 5:
			return c(a.SubPoint < b.SubPoint)

		case 6:
			return c(a.AddPoint < b.AddPoint)

		case 7:
			return c(a.FixedSales < b.FixedSales)

		case 8:
			return c(a.PayType < b.PayType)

		case 9:
			return c(a.CreateDate < b.CreateDate)
		}

		panic("unreachable")
	})

	return r.SorterBase.Sort(col, order)
}

func (r *RevenuesModel) ResetRows(startDate, endDate string, memberId int) {
	var err error
	var revenueList []dto.RevenueDto

	if revenueList, err = service.FindRevenueList(dbconn, startDate, endDate, memberId); err != nil {
		panic(err.Error())
	}
	r.revenues = make([]*dto.RevenueDto, len(revenueList))
	for i := range revenueList {
		r.revenues[i] = &dto.RevenueDto{
			RevenueId:   revenueList[i].RevenueId,
			MemberId:    revenueList[i].MemberId,
			MemberName:  revenueList[i].MemberName,
			PhoneNumber: revenueList[i].PhoneNumber,
			Sales:       revenueList[i].Sales,
			SubPoint:    revenueList[i].SubPoint,
			AddPoint:    revenueList[i].AddPoint,
			FixedSales:  revenueList[i].FixedSales,
			PayType:     revenueList[i].PayType,
			CreateDate:  revenueList[i].CreateDate,
		}
	}

	// Notify TableView and other interested parties about the reset.
	r.PublishRowsReset()

	r.Sort(r.sortColumn, r.sortOrder)
}
