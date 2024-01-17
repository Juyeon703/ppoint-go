package service

import (
	"fmt"
	"ppoint/dto"
	"ppoint/query"
	"ppoint/types"
)

func RevenueAdd(dbconn *query.DbConfig, revenueDto *dto.RevenueAddDto) error {
	var err error
	var setting = new(types.Setting)
	var revenue = new(types.Revenue)
	var changePoint int

	fmt.Println("===SelectSettingByPayType() 호출")
	if setting, err = dbconn.SelectSettingByPayType(revenueDto.PayType); err != nil {
		return err
	}
	revenue.MemberId = revenueDto.MemberId
	revenue.Sales = revenueDto.Sales
	revenue.SubPoint = revenueDto.SubPoint
	revenue.PayType = revenueDto.PayType
	if revenueDto.SubPoint == 0 {
		changePoint = revenueDto.Sales * setting.SettingValue / 100
		revenue.AddPoint = changePoint
		revenue.FixedSales = revenueDto.Sales
	} else {
		if revenueDto.SubPoint < 0 {
			return err
		}
		revenue.AddPoint = 0
		revenue.FixedSales = revenueDto.Sales - revenueDto.SubPoint
		changePoint = -(revenueDto.SubPoint)
	}
	fmt.Println("===UpdateMemberByPoint() 호출")
	if err = dbconn.UpdateMemberByPoint(revenueDto.MemberId, changePoint); err != nil {
		return err
	}
	fmt.Println("===CreateRevenue() 호출")
	if err = dbconn.CreateRevenue(revenue); err != nil {
		return err
	}
	return nil
}
