package service

import (
	"fmt"
	"ppoint/query"
)

func FindSettingValue(dbconn *query.DbConfig, settingType string) (int, error) {
	var err error
	var result int
	if result, err = dbconn.SelectSettingBySettingType(settingType); err != nil {
		return 0, err
	}
	fmt.Println("=============> SelectSettingByPayType() 호출")
	return result, nil
}
