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
	if revenueDto.PayType == types.Card {
		fmt.Println("===SelectSettingBySettingType(SettingCard) 호출")
		if result, err = dbconn.SelectSettingBySettingType(types.SettingCard); err != nil {
			return err
		}
	} else if revenueDto.PayType == types.Cash {
		fmt.Println("===SelectSettingBySettingType(SettingCash) 호출")
		if result, err = dbconn.SelectSettingBySettingType(types.SettingCash); err != nil {
			return err
		}
	} else {
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

func PointEdit(dbconn *query.DbConfig, memberId string, originPoint, newPoint int) error {
	var err error
	var revenue = new(types.Revenue)

	calc := originPoint - newPoint
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

func FindRevenueList(dbconn *query.DbConfig, startDate, endDate string, memberId int) ([]dto.RevenueDto, error) {
	var err error
	var revenueList []dto.RevenueDto

	if memberId != 0 && startDate == "" && endDate == "" {
		if revenueList, err = dbconn.SelectRevenuesByMember(memberId); err != nil {
			return nil, err
		} else {
			fmt.Println("=============> SelectRevenuesByMember() 호출")
		}
	} else {
		if revenueList, err = dbconn.SelectRevenuesByCustomDate(startDate, endDate); err != nil {
			return nil, err
		} else {
			fmt.Println("=============> SelectRevenuesByCustomDate() 호출")
		}
	}

	return revenueList, nil
}
