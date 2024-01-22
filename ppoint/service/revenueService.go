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
	var result string
	var changePoint, addPointTemp, fixedSalesTemp int

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

	settingValue, _ := strconv.Atoi(result)
	if revenueDto.SubPoint == 0 {
		changePoint = revenueDto.Sales * settingValue / 100
		addPointTemp = changePoint
		fixedSalesTemp = revenueDto.Sales
	} else {
		if revenueDto.SubPoint < 0 {
			return err
		}
		addPointTemp = 0
		fixedSalesTemp = revenueDto.Sales - revenueDto.SubPoint
		changePoint = -(revenueDto.SubPoint)
	}
	fmt.Println("===UpdateMemberByPoint() 호출")
	if err = dbconn.UpdateMemberByPoint(revenueDto.MemberId, changePoint); err != nil {
		return err
	}
	fmt.Println("===CreateRevenue() 호출")
	if err = dbconn.CreateRevenue(revenueDto.MemberId, revenueDto.Sales, revenueDto.SubPoint, addPointTemp, fixedSalesTemp, revenueDto.PayType); err != nil {
		return err
	}
	return nil
}

func PointEdit(dbconn *query.DbConfig, memberId string, originPoint, newPoint int) error {
	var err error
	var subPoint, addPoint int

	calc := originPoint - newPoint
	if calc < 0 {
		subPoint = 0
		addPoint = -calc
	} else if calc > 0 {
		subPoint = calc
		addPoint = 0
	} else {
		return err
	}
	result, _ := strconv.Atoi(memberId)
	fmt.Println("===CreateRevenue() 호출")
	if err = dbconn.CreateRevenue(result, 0, subPoint, addPoint, 0, "포인트 변경"); err != nil {
		return err
	}

	return nil
}

func FindRevenueList(dbconn *query.DbConfig, startDate, endDate string, memberId int) ([]dto.RevenueDto, error) {
	var err error
	var revenueList []dto.RevenueDto

	if memberId != 0 {
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

	fmt.Println(len(revenueList), startDate, endDate, memberId)
	return revenueList, nil
}

func FindSumSalesPoint(dbconn *query.DbConfig, startDate, endDate string, memberId int) (*dto.SumSalesPointDto, error) {
	var err error
	var result *dto.SumSalesPointDto

	if memberId != 0 {
		if result, err = dbconn.SelectSumSalesPointByMemberId(memberId); err != nil {
			return nil, err
		} else {
			fmt.Println("=============> SelectSumSalesPointByMemberId() 호출")
		}
	} else {
		if result, err = dbconn.SelectSumSalesPointByCustomDate(startDate, endDate); err != nil {
			return nil, err
		} else {
			fmt.Println("=============> SelectSumSalesPointByCustomDate() 호출")
		}
	}

	return result, nil
}
