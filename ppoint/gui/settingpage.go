package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/types"
)

type SettingPage struct {
	*walk.Composite
}

func newSettingPage(parent walk.Container) (Page, error) {
	p := new(SettingPage)

	var pCash, pCard, pPointLimit, pDbBackupPath *walk.LineEdit

	var strCash string
	var strCard string
	var strPointLimit string
	var strDbBackupPath string

	var err error
	var settingList []types.Setting

	if settingList, err = dbconn.SelectSettings(); err == nil {
		for _, sett := range settingList {
			if sett.SettingType == "db_backup_path" {
				strDbBackupPath = sett.SettingValue
			} else if sett.SettingType == "point_limit" {
				strPointLimit = sett.SettingValue
			} else if sett.SettingType == "pay_type_card" {
				strCard = sett.SettingValue
			} else if sett.SettingType == "pay_type_cash" {
				strCash = sett.SettingValue
			}
		}
	}

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "설정 페이지",
		Layout:   VBox{},
		OnBoundsChanged: func() {
			pCash.SetText(strCash)
			pCard.SetText(strCard)
			pPointLimit.SetText(strPointLimit)
			pDbBackupPath.SetText(strDbBackupPath)
		},
		Children: []Widget{
			Label{Text: "설정 페이지"},
			HSpacer{},
			HSplitter{
				Children: []Widget{
					//left box
					//VSplitter{},

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
										AssignTo: &pCash,
										Name:     "text_cash_point_percent",
										Text:     Bind("cash_point_percent"),
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
										AssignTo: &pCard,
										Name:     "text_card_point_percent",
										Text:     Bind("card_point_percent"),
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
										AssignTo: &pPointLimit,
										Name:     "text_usable_point_limit",
										Text:     Bind("usable_point_limit"),
									},
									Label{
										Text: "포인트 이상",
									},
								},
							},

							HSplitter{
								Children: []Widget{
									Label{
										Name: "db_bakup_path",
										Text: "데이터 백업 경로",
									},
									LineEdit{
										AssignTo: &pDbBackupPath,
										Name:     "text_db_bakup_path",
										Text:     Bind("db_bakup_path"),
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
					var settingGrp []string
					settingGrp = []string{"pay_type_cash", "pay_type_card", "point_limit", "db_backup_path"}

					if pCash.Text() != "" && pCard.Text() != "" && pPointLimit.Text() != "" && pDbBackupPath.Text() != "" {
						for _, settType := range settingGrp {
							var inputValue string
							if settType == "pay_type_cash" {
								inputValue = pCash.Text()
							}
							if settType == "pay_type_card" {
								inputValue = pCard.Text()
							}
							if settType == "point_limit" {
								inputValue = pPointLimit.Text()
							}
							if settType == "db_backup_path" {
								inputValue = pDbBackupPath.Text()
							}

							fmt.Println(settType, inputValue)
							if err = dbconn.UpdateSettingByType(settType, inputValue); err != nil {
								//err
							} else {
								fmt.Println("success")
							}
						}
					}

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
