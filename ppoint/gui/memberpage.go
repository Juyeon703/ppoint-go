package gui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MemberPage struct {
	*walk.Composite
}

func newMemberPage(parent walk.Container) (Page, error) {
	p := new(MemberPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "고객 관리 페이지",
		Layout:   HBox{},
		Children: []Widget{
			Label{Text: "고객 관리 페이지"},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}
