package gui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type SettingPage struct {
	*walk.Composite
}

func newSettingPage(parent walk.Container) (Page, error) {
	p := new(SettingPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "설정 페이지",
		Layout:   VBox{},
		Children: []Widget{
			Label{Text: "설정 페이지"},
			HSpacer{},
			HSplitter{
				Children: []Widget{
					//left box
					VSplitter{
						Children: []Widget{
							Label{
								ColumnSpan: 10,
								Text:       "* 잠재 고객 : 포인트 적립을 한번도 하지 않은 고객",
							},
							HSpacer{},
							Label{
								Text: "* 신규 고객 : 포인트 적립을 처음으로 1회 한 고객",
							},
							HSpacer{},
							Label{
								Name: "woosoo",
								Text: "* 우수 고객 ",
							},
							HSpacer{},
							HSplitter{
								Children: []Widget{
									Label{
										Name: "woosoo_sales",
										Text: "매출액",
									},
									LineEdit{
										Name: "text_woosoo_sales",
										Text: Bind("woosoo"),
									},
									Label{
										Text: "원 이상",
									},
								},
							},
							HSplitter{
								Children: []Widget{
									Label{
										Name: "woosoo_visit_count",
										Text: "방문횟수",
									},
									LineEdit{
										Name: "text_woosoo_visit_count",
										Text: Bind("woosoo_visit_count"),
									},
									Label{
										Text: "회 이상",
									},
								},
							},
							HSpacer{},
							Label{
								Name: "forever",
								Text: "* 평생 고객 ",
							},
							HSpacer{},
							HSplitter{
								Children: []Widget{
									Label{
										Name: "forever_sales",
										Text: "매출액",
									},
									LineEdit{
										Name: "text_forever_sales",
										Text: Bind("forever"),
									},
									Label{
										Text: "원 이상",
									},
								},
							},
							HSplitter{
								Children: []Widget{
									Label{
										Name: "forever_visit_count",
										Text: "방문횟수",
									},
									LineEdit{
										Name: "text_forever_visit_count",
										Text: Bind("forever_visit_count"),
									},
									Label{
										Text: "회 이상",
									},
								},
							},
						},
					},

					//right box
					VSplitter{
						Children: []Widget{
							HSplitter{
								Children: []Widget{
									Label{
										Name: "cash_point_percent",
										Text: "현금 매출 액의 ",
									},
									LineEdit{
										Name: "text_cash_point_percent",
										Text: Bind("cash_point_percent"),
									},
									Label{
										Text: "% 적립",
									},
								},
							},
							HSplitter{
								Children: []Widget{
									Label{
										Name: "card_point_percent",
										Text: "카드 매출 액의 ",
									},
									LineEdit{
										Name: "text_card_point_percent",
										Text: Bind("card_point_percent"),
									},
									Label{
										Text: "% 적립",
									},
								},
							},
							HSplitter{
								Children: []Widget{
									Label{
										Name: "usable_point_limit",
										Text: "현금처럼 사용 ",
									},
									LineEdit{
										Name: "text_usable_point_limit",
										Text: Bind("usable_point_limit"),
									},
									Label{
										Text: "포인트 이상",
									},
								},
							},
						},
					},
				},
			},
			PushButton{
				Text: "저장",
				OnClicked: func() {
					/*if err := db.Submit(); err != nil {
						log.Print(err)
						return
					}

					dlg.Accept()*/
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
