package gui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"ppoint/types"
)

type SettingPage struct {
	*walk.Composite
}

func newSettingPage(parent walk.Container) (Page, error) {
	p := new(SettingPage)

	var pCash, pCard, pPointLimit, pDbBackupPath, pLogPath *walk.LineEdit

	var strCash string
	var strCard string
	var strPointLimit string
	var strDbBackupPath string
	var strLogPath string

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
			} else if sett.SettingType == "log_path" {
				strLogPath = sett.SettingValue
			}
		}
	}

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "설정 페이지",
		Layout:   VBox{},
		Border:   true,
		MinSize:  Size{subWidth, subHeight},
		OnBoundsChanged: func() {
			pCash.SetText(strCash)
			pCard.SetText(strCard)
			pPointLimit.SetText(strPointLimit)
			pDbBackupPath.SetText(strDbBackupPath)
			pLogPath.SetText(strLogPath)
		},

		Children: []Widget{
			Label{Text: "설정 페이지"},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Composite{
						Layout: HBox{},
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
							HSpacer{
								Size: subWidth / 2,
							},
						},
					},
				},
			},

			Composite{
				Layout: HBox{},
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
					HSpacer{
						Size: subWidth / 2,
					},
				},
			},

			Composite{
				Layout: HBox{},
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
					HSpacer{
						Size: subWidth / 2,
					},
				},
			},
			Composite{
				Layout: HBox{},
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
					HSpacer{
						Size: subWidth / 2,
					},
				},
			},

			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Name: "log_path",
						Text: "로그 파일 경로",
					},
					LineEdit{
						AssignTo: &pLogPath,
						Name:     "text_log_path",
						Text:     Bind("log_path"),
					},
					HSpacer{
						Size: subWidth / 2,
					},
				},
			},

			Composite{
				Layout: HBox{},
				Children: []Widget{
					VSpacer{
						Size: subHeight / 2,
					},
				},
			},

			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						MaxSize: Size{subWidth / 4, subHeight},
						Text:    "저장",
						OnClicked: func() {
							var settingGrp []string
							settingGrp = []string{"pay_type_cash", "pay_type_card", "point_limit", "db_backup_path", "log_path"}

							if pCash.Text() != "" && pCard.Text() != "" && pPointLimit.Text() != "" && pDbBackupPath.Text() != "" && pLogPath.Text() != "" {
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
									if settType == "log_path" {
										inputValue = pLogPath.Text()
									}

									if err = dbconn.UpdateSettingByType(settType, inputValue); err != nil {
										//err
										MsgBox("에러", "저장에 실패하였습니다")
									} else {
										log.Infof("[Update]UpdateSettingByType() // param{settingType : %s, settingValue : %s}", settType, inputValue)

									}
								}

								MsgBox("알림", "저장되었습니다.")
							} else {
								MsgBox("알림", "값을 입력해주세요.")
							}

							/*if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()*/
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
