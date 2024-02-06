package service

import (
	"fmt"
	"ppoint/query"
	"strconv"
)

func FindSettingValue(dbconn *query.DbConfig, settingType string) (int, error) {
	var err error
	var result string
	if result, err = dbconn.SelectSettingBySettingType(settingType); err != nil {
		return 0, err
	}
	temp, _ := strconv.Atoi(result)
	fmt.Println("=============> SelectSettingByPayType() 호출")
	return temp, nil
}

func FindSettingStrValue(dbconn *query.DbConfig, settingType string) (string, error) {
	var err error
	var result string
	if result, err = dbconn.SelectSettingBySettingType(settingType); err != nil {
		return "", err
	}
	return result, nil
}
