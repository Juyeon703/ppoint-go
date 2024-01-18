package service

import (
	"fmt"
	"ppoint/dto"
	"ppoint/query"
	"ppoint/types"
	"strconv"
)

func RevenueAdd(dbconn *query.DbConfig, revenueDto *dto.RevenueAddDto) error {
	var err error
	var result int
	var revenue = new(types.Revenue)
	var changePoint int
	if revenueDto.MemberId <= 0 {
		return err
	}
	fmt.Println("===SelectSettingByPayType() 호출")
	if result, err = dbconn.SelectSettingByPayType(revenueDto.PayType); err != nil {
		return err
	}
	revenue.MemberId = revenueDto.MemberId
	revenue.Sales = revenueDto.Sales
	revenue.SubPoint = revenueDto.SubPoint
	revenue.PayType = revenueDto.PayType
	if revenueDto.SubPoint == 0 {
		changePoint = revenueDto.Sales * result / 100
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

func PointEdit(dbconn *query.DbConfig, memberId, originPoint, newPoint string) error {
	var err error
	var revenue = new(types.Revenue)

	originP, _ := strconv.Atoi(originPoint)
	newP, _ := strconv.Atoi(newPoint)
	calc := originP - newP
	if calc < 0 {
		revenue.SubPoint = 0
		revenue.AddPoint = -calc
	} else if calc > 0 {
		revenue.SubPoint = calc
		revenue.AddPoint = 0
	} else {
		return err
	}
	result, _ := strconv.Atoi(memberId)
	revenue.MemberId = result
	revenue.Sales = 0
	revenue.FixedSales = 0
	revenue.PayType = "포인트 변경"
	fmt.Println("===CreateRevenue() 호출")
	if err = dbconn.CreateRevenue(revenue); err != nil {
		return err
	}

	return nil
}
