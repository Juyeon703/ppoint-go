package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/dto"
	"strconv"
)

func RunMemberSearchDialog(owner walk.Form, memberList []dto.MemberDto, memberNameLE, phoneNumberLE, birthLE, udtLE *walk.LineEdit, memberIdNE, pointNE, countNE *walk.NumberEdit) (int, error) {
	var dlg *walk.Dialog
	var acceptPB, cancelPB *walk.PushButton
	var tv *walk.TableView
	model := NewSearchMembersModel(memberList)

	return Dialog{
		AssignTo:      &dlg,
		Title:         "회원 검색",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       Size{300, 300},
		Layout:        VBox{},
		Children: []Widget{
			TableView{
				Name:             "searchTable",
				AssignTo:         &tv,
				AlternatingRowBG: true,
				ColumnsOrderable: true,
				MultiSelection:   true,
				MinSize:          Size{200, 200},
				Columns: []TableViewColumn{
					{Title: "번호", DataMember: "MemberId", Hidden: true},
					{Title: "이름", DataMember: "MemberName"},
					{Title: "핸드폰번호", DataMember: "PhoneNumber"},
					{Title: "생일", DataMember: "Birth", Hidden: true},
					{Title: "보유포인트", DataMember: "TotalPoint", Alignment: AlignFar, Hidden: true},
					{Title: "방문 횟수", DataMember: "VisitCount", Hidden: true},
					{Title: "가입일", DataMember: "CreateDate", Width: 150, Hidden: true},
					{Title: "최근방문일", DataMember: "UpdateDate", Width: 150, Hidden: true},
				},
				Model: model,
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							index := tv.SelectedIndexes()
							fmt.Println("클릭한 인덱스 : ", index, fmt.Sprintf("%v", model.Value(index[0], 0)), fmt.Sprintf("%v", model.Value(index[0], 1)))
							if len(index) > 0 {
								memberIdFl, _ := strconv.ParseFloat(fmt.Sprintf("%v", model.Value(index[0], 0)), 64)
								memberIdNE.SetValue(memberIdFl)
								memberNameLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 1)))
								phoneNumberLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 2)))
								birthLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 3)))
								pointFL, _ := strconv.ParseFloat(fmt.Sprintf("%v", model.Value(index[0], 4)), 64)
								pointNE.SetValue(pointFL)
								countFL, _ := strconv.ParseFloat(fmt.Sprintf("%v", model.Value(index[0], 4)), 64)
								countNE.SetValue(countFL)
								udtLE.SetText(fmt.Sprintf("%v", model.Value(index[0], 7)))
							}
							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}

func NewSearchMembersModel(memberList []dto.MemberDto) *SearchMembersModel {

	m := &SearchMembersModel{members: make([]*dto.MemberDto, len(memberList))}

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

	return m
}

type SearchMembersModel struct {
	walk.SortedReflectTableModelBase
	members []*dto.MemberDto
}

func (m *SearchMembersModel) Items() interface{} {
	return m.members
}
