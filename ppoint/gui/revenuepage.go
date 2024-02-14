package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/dto"
	"ppoint/service"
	"ppoint/types"
	"ppoint/utils"
	"sort"
	"strconv"
	"time"
)

type SalesPage struct {
	*walk.Composite
}

func newSalesPage(parent walk.Container) (Page, error) {
	var err error
	p := new(SalesPage)
	var tv *walk.TableView
	var datedb *walk.DataBinder
	var tvResultLabel *walk.Label
	var startDateSearchDE, endDateSearchDE *walk.DateEdit
	var sumNEcc, sumNEcard, sumNEcash, sumNEaddP, sumNEsubP *walk.LineEdit
	var dateSearch = &SearchDate{Sdt: time.Now(), Edt: time.Now()}
	dateNow := time.Now().Format("2006-01-02")
	model := NewRevenuesModel(dateNow, dateNow, moveId)
	var sumDto *dto.SumSalesPointDto
	if sumDto, err = service.FindSumSalesPoint(dbconn, dateNow, dateNow, moveId); err != nil {
		return nil, err
	}

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "매출 관리",
		Layout:   VBox{},
		Border:   true,
		MinSize:  Size{subWidth, subHeight},
		Children: []Widget{
			Label{Text: "매출 관리 페이지"},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text: "기간별 조회 :",
					},
					PushButton{
						Text: "일별",
						OnClicked: func() {
							now := time.Now()
							startDateSearchDE.SetDate(now)
							endDateSearchDE.SetDate(now)
							model = tvRevenueReloading(dateSearch, 0, tv, tvResultLabel, datedb, sumNEcc, sumNEcard, sumNEcash, sumNEaddP, sumNEsubP)
						},
					},
					PushButton{
						Text: "월별",
						OnClicked: func() {
							now := time.Now()
							startDateSearchDE.SetDate(time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local))
							endDateSearchDE.SetDate(time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, time.Local))
							model = tvRevenueReloading(dateSearch, 0, tv, tvResultLabel, datedb, sumNEcc, sumNEcard, sumNEcash, sumNEaddP, sumNEsubP)
						},
					},
					PushButton{
						Text: "년도별",
						OnClicked: func() {
							now := time.Now()
							startDateSearchDE.SetDate(time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local))
							endDateSearchDE.SetDate(time.Date(now.Year(), 12, 31, 0, 0, 0, 0, time.Local))
							model = tvRevenueReloading(dateSearch, 0, tv, tvResultLabel, datedb, sumNEcc, sumNEcard, sumNEcash, sumNEaddP, sumNEsubP)
						},
					},
				},
			},
			Composite{
				Layout: HBox{},
				DataBinder: DataBinder{
					AssignTo:   &datedb,
					Name:       "dateSearch",
					DataSource: dateSearch,
				},
				Children: []Widget{
					Label{
						Text: "검색 : ",
					},
					DateEdit{
						AssignTo: &startDateSearchDE,
						Date:     Bind("Sdt"),
						OnBoundsChanged: func() {
							startDateSearchDE.SetDate(dateSearch.Sdt)
						},
					},
					Label{
						Text: " ~ ",
					},
					DateEdit{
						AssignTo: &endDateSearchDE,
						Date:     Bind("Edt"),
						OnBoundsChanged: func() {
							endDateSearchDE.SetDate(dateSearch.Edt)
						},
					},
					PushButton{
						Text: "검색",
						OnClicked: func() {
							model = tvRevenueReloading(dateSearch, 0, tv, tvResultLabel, datedb, sumNEcc, sumNEcard, sumNEcash, sumNEaddP, sumNEsubP)
						},
					},
					HSpacer{
						Size: 400,
					},
					Label{
						Text:     "검색 수 : " + strconv.Itoa(model.RowCount()),
						AssignTo: &tvResultLabel,
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
					{Title: "번호", DataMember: "index", Width: 40},
					{Title: "매출번호", DataMember: "RevenueId", Hidden: true},
					{Title: "이름", DataMember: "MemberName"},
					{Title: "핸드폰번호", DataMember: "PhoneNumber"},
					{Title: "결제금액", DataMember: "Sales", Alignment: AlignFar, FormatFunc: func(value interface{}) string {
						r, _ := utils.InterfaceToInt64(value)
						return utils.MoneyConverter(r)
					}},
					{Title: "사용포인트", DataMember: "SubPoint", Alignment: AlignFar, FormatFunc: func(value interface{}) string {
						r, _ := utils.InterfaceToInt64(value)
						return utils.MoneyConverter(r)
					}},
					{Title: "적립포인트", DataMember: "AddPoint", Alignment: AlignFar, FormatFunc: func(value interface{}) string {
						r, _ := utils.InterfaceToInt64(value)
						return utils.MoneyConverter(r)
					}},
					{Title: "실제결제금액", DataMember: "FixedSales", Alignment: AlignFar, FormatFunc: func(value interface{}) string {
						r, _ := utils.InterfaceToInt64(value)
						return utils.MoneyConverter(r)
					}},
					{Title: "결제방법", DataMember: "PayType"},
					{Title: "결제일", DataMember: "CreateDate", Width: 150},
					{Title: "#", Width: 40, FormatFunc: func(value interface{}) string {
						return "삭제"
					}},
				},
				Model: model,
				//TODO:삭제 글씨 클릭시로 변경
				OnSelectedIndexesChanged: func() {
					var point int
					index := tv.SelectedIndexes()

					if len(index) > 0 {
						revenueId, _ := strconv.Atoi(fmt.Sprintf("%v", model.Value(index[0], 1)))
						memberId, _ := strconv.Atoi(fmt.Sprintf("%v", model.Value(index[0], 10)))
						subPoint, _ := strconv.Atoi(fmt.Sprintf("%v", model.Value(index[0], 5)))
						addPoint, _ := strconv.Atoi(fmt.Sprintf("%v", model.Value(index[0], 6)))

						payType := fmt.Sprintf("%v", model.Value(index[0], 8))

						fmt.Println("//////")
						fmt.Println(tv.SelectedIndexes())
						fmt.Println(revenueId, memberId, subPoint, addPoint)

						if payType == types.Card || payType == types.Cash || payType == "포인트" {
							if cmd, err := RunRevenueDeleteDialog(winMain); err != nil {
								log.Error(err)
							} else if cmd == walk.DlgCmdOK {
								if point, err = dbconn.SelectMemberPoint(memberId); err != nil {
									log.Error(err)
								}
								log.Debugf("(매출 삭제) 회원 보유 포인트 조회 >>>> memberId : [%d], Point : [%d]", memberId, point)
								if point-addPoint < 0 {
									MsgBox("오류", "해당 내역으로 적립된 포인트가 이미 사용되어 해당 내역을 삭제할 수 없습니다.")
								} else {
									if err := service.RevenueDelete(dbconn, revenueId, memberId, subPoint, addPoint); err != nil {
										log.Error(err.Error())
										panic(err)
									}
									model = tvRevenueReloading(dateSearch, moveId, tv, tvResultLabel, datedb, sumNEcc, sumNEcard, sumNEcash, sumNEaddP, sumNEsubP)
								}
							}
						} else {
							MsgBox("매출 삭제", "삭제할 수 없는 내역입니다.")
						}
					}
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
						Layout: Grid{Columns: 8},
						Children: []Widget{
							Label{
								Text: "총 매출 금액(현금+카드)",
							},
							LineEdit{
								AssignTo: &sumNEcc,
								//Suffix:   " 원",
								ReadOnly: true,
							},
							HSpacer{
								Size: 20,
							},
							Label{
								Text: "총 매출 금액(카드)",
							},
							LineEdit{
								AssignTo: &sumNEcard,
								//Suffix:   " 원",
								ReadOnly: true,
							},
							HSpacer{
								Size: 20,
							},
							Label{
								Text: "총 매출 금액(현금)",
							},
							LineEdit{
								AssignTo: &sumNEcash,
								//Suffix:   " 원",
								ReadOnly: true,
							},
							Label{
								Text: "총 적립 포인트",
							},
							LineEdit{
								AssignTo: &sumNEaddP,
								//Suffix:   " p",
								ReadOnly: true,
							},
							HSpacer{
								Size: 20,
							},
							Label{
								Text: "총 사용 포인트(소멸 제외)",
							},
							LineEdit{
								AssignTo: &sumNEsubP,
								//Suffix:   " p",
								ReadOnly: true,
								OnBoundsChanged: func() {
									/*sumNEcc.SetValue(float64(sumDto.SumSales))
									sumNEcard.SetValue(float64(sumDto.SumCard))
									sumNEcash.SetValue(float64(sumDto.SumCash))
									sumNEaddP.SetValue(float64(sumDto.SumAddP))
									sumNEsubP.SetValue(float64(sumDto.SumSubP))*/
									sumNEcc.SetText(utils.MoneyConverter(sumDto.SumSales) + " 원")
									sumNEcard.SetText(utils.MoneyConverter(sumDto.SumCard) + " 원")
									sumNEcash.SetText(utils.MoneyConverter(sumDto.SumCash) + " 원")
									sumNEaddP.SetText(utils.MoneyConverter(sumDto.SumAddP) + " P")
									sumNEsubP.SetText(utils.MoneyConverter(sumDto.SumSubP) + " P")
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

type SearchDate struct {
	Sdt time.Time
	Edt time.Time
}

func tvRevenueReloading(dateSearch *SearchDate, memberId int, tv *walk.TableView, tvResultLabel *walk.Label, datedb *walk.DataBinder, sumNEcc, sumNEcard, sumNEcash, sumNEaddP, sumNEsubP *walk.LineEdit) *RevenuesModel {
	if err := datedb.Submit(); err != nil {
		log.Error(err.Error())
		return nil
	}
	startDate := dateSearch.Sdt.Format("2006-01-02")
	endDate := dateSearch.Edt.Format("2006-01-02")

	model := NewRevenuesModel(startDate, endDate, memberId)
	tv.SetModel(model)
	tvResultLabel.SetText("검색 수 : " + strconv.Itoa(model.RowCount()))
	if err := SumInfoLoading(startDate, endDate, memberId, sumNEcc, sumNEcard, sumNEcash, sumNEaddP, sumNEsubP); err != nil {
		log.Error(err.Error())
		panic(err.Error())
	}
	return model
}

func SumInfoLoading(startDate, endDate string, memberId int, sumNEcc, sumNEcard, sumNEcash, sumNEaddP, sumNEsubP *walk.LineEdit) error {
	var err error
	var result *dto.SumSalesPointDto
	if result, err = service.FindSumSalesPoint(dbconn, startDate, endDate, memberId); err != nil {
		return err
	}
	/*	sumNEcc.SetText(utils.MoneyConverter(result.SumSales))
		sumNEcard.SetText(utils.MoneyConverter(result.SumCard))
		sumNEcash.SetText(utils.MoneyConverter(result.SumCash))
		sumNEaddP.SetText(utils.MoneyConverter(result.SumAddP))
		sumNEsubP.SetText(utils.MoneyConverter(result.SumSubP))*/

	sumNEcc.SetText(utils.MoneyConverter(result.SumSales) + " 원")
	sumNEcard.SetText(utils.MoneyConverter(result.SumCard) + " 원")
	sumNEcash.SetText(utils.MoneyConverter(result.SumCash) + " 원")
	sumNEaddP.SetText(utils.MoneyConverter(result.SumAddP) + " P")
	sumNEsubP.SetText(utils.MoneyConverter(result.SumSubP) + " P")
	return nil
}

type RevenuesModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	revenues   []*RevenueTV
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
		return revenue.index

	case 1:
		return revenue.RevenueId

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

	case 10:
		return revenue.delete
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
			return c(a.index < b.index)

		case 1:
			return c(a.RevenueId < b.RevenueId)

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

		case 10:
			return c(a.delete < b.delete)
		}

		panic("unreachable")
	})

	return r.SorterBase.Sort(col, order)
}

func (r *RevenuesModel) ResetRows(startDate, endDate string, memberId int) {
	var err error
	var revenueList []dto.RevenueDto

	if revenueList, err = service.FindRevenueList(dbconn, startDate, endDate, memberId); err != nil {
		log.Error(err.Error())
		panic(err.Error())
	}
	r.revenues = make([]*RevenueTV, len(revenueList))
	for i := range revenueList {
		r.revenues[i] = &RevenueTV{
			index:       i + 1,
			RevenueId:   revenueList[i].RevenueId,
			MemberName:  revenueList[i].MemberName,
			PhoneNumber: revenueList[i].PhoneNumber,
			Sales:       revenueList[i].Sales,
			SubPoint:    revenueList[i].SubPoint,
			AddPoint:    revenueList[i].AddPoint,
			FixedSales:  revenueList[i].FixedSales,
			PayType:     revenueList[i].PayType,
			CreateDate:  revenueList[i].CreateDate,
			delete:      revenueList[i].MemberId,
		}
	}

	// Notify TableView and other interested parties about the reset.
	r.PublishRowsReset()

	r.Sort(r.sortColumn, r.sortOrder)
}

type RevenueTV struct {
	index       int
	RevenueId   int
	MemberName  string
	PhoneNumber string
	Sales       int
	SubPoint    int
	AddPoint    int
	FixedSales  int
	PayType     string
	CreateDate  string
	delete      int
}
