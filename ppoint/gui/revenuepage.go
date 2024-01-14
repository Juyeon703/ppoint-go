package gui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type SalesPage struct {
	*walk.Composite
}

func newSalesPage(parent walk.Container) (Page, error) {
	p := new(SalesPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "매출 관리 페이지",
		Layout:   HBox{},
		Children: []Widget{
			Label{Text: "매출 관리 페이지"},
			HSpacer{},
			VSplitter{
				Children: []Widget{
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
