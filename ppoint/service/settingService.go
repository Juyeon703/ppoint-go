package service

import (
	"ppoint/query"
	"strconv"
)

func FindSettingValue(dbconn *query.DbConfig, settingType string) (int, error) {
	var err error
	log := dbconn.Logue
	var result string
	if result, err = dbconn.SelectSettingBySettingType(settingType); err != nil {
		return 0, err
	}
	temp, _ := strconv.Atoi(result)
	log.Debugf("[Select]SelectSettingByPayType() // param{settingType : %s}", settingType)
	return temp, nil
}

// 로그 사용 불가
func FindSettingStrValue(dbconn *query.DbConfig, settingType string) (string, error) {
	var err error
	var result string
	if result, err = dbconn.SelectSettingBySettingType(settingType); err != nil {
		return "", err
	}
	return result, nil
}
