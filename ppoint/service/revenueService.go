package service

import (
	"ppoint/dto"
	"ppoint/query"
	"ppoint/types"
	"strconv"
)

func RevenueAdd(dbconn *query.DbConfig, revenueDto *dto.RevenueAddDto) error {
	var err error
	log := dbconn.Logue
	var result string
	var changePoint, addPointTemp, fixedSalesTemp int

	if revenueDto.MemberId <= 0 {
		return err
	}
	if revenueDto.PayType == types.Card {
		if result, err = dbconn.SelectSettingBySettingType(types.SettingCard); err != nil {
			return err
		}
		log.Debug("[Select]SelectSettingBySettingType(card)")
	} else if revenueDto.PayType == types.Cash {
		if result, err = dbconn.SelectSettingBySettingType(types.SettingCash); err != nil {
			return err
		}
		log.Debug("[Select]SelectSettingBySettingType(cash)")
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
	if err = dbconn.UpdateMemberByPoint(revenueDto.MemberId, changePoint); err != nil {
		return err
	}
	log.Infof("[Update]UpdateMemberByPoint() // param{memberId : %d, point : %d}", revenueDto.MemberId, changePoint)
	if err = dbconn.CreateRevenue(revenueDto.MemberId, revenueDto.Sales, revenueDto.SubPoint, addPointTemp, fixedSalesTemp, revenueDto.PayType); err != nil {
		return err
	}
	log.Infof("[Create]CreateRevenue() // param{memberId : %d, sales : %d, subPoint : %d, addPoint : %d, fixedSales : %d, payType : %s}",
		revenueDto.MemberId, revenueDto.Sales, revenueDto.SubPoint, addPointTemp, fixedSalesTemp, revenueDto.PayType)
	return nil
}

func PointEdit(dbconn *query.DbConfig, memberId string, originPoint, newPoint int) error {
	var err error
	log := dbconn.Logue
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
	if err = dbconn.CreateRevenue(result, 0, subPoint, addPoint, 0, "포인트 변경"); err != nil {
		return err
	}
	log.Infof("[Create]CreateRevenue() // param{memberId : %d, subPoint : %d, addPoint : %d, payType : 포인트 변경}", result, subPoint, addPoint)
	return nil
}

func FindRevenueList(dbconn *query.DbConfig, startDate, endDate string, memberId int) ([]dto.RevenueDto, error) {
	var err error
	log := dbconn.Logue
	var revenueList []dto.RevenueDto

	if memberId != 0 {
		if revenueList, err = dbconn.SelectRevenuesByMember(memberId); err != nil {
			return nil, err
		}
		log.Debugf("[Select]SelectRevenuesByMember() // param{memberId : %d}", memberId)

	} else {
		if revenueList, err = dbconn.SelectRevenuesByCustomDate(startDate, endDate); err != nil {
			return nil, err
		}
		log.Debugf("[Select]SelectRevenuesByCustomDate() // param{startDate : %s, endDate : %s}", startDate, endDate)

	}
	return revenueList, nil
}

func FindSumSalesPoint(dbconn *query.DbConfig, startDate, endDate string, memberId int) (*dto.SumSalesPointDto, error) {
	var err error
	log := dbconn.Logue
	var result *dto.SumSalesPointDto

	if memberId != 0 {
		if result, err = dbconn.SelectSumSalesPointByMemberId(memberId); err != nil {
			return nil, err
		}
		log.Debugf("[Select]SelectSumSalesPointByMemberId() // param{memberId : %d}", memberId)

	} else {
		if result, err = dbconn.SelectSumSalesPointByCustomDate(startDate, endDate); err != nil {
			return nil, err
		}
		log.Debugf("[Select]SelectSumSalesPointByCustomDate() // param{startDate : %s, endDate : %s}", startDate, endDate)
	}

	return result, nil
}
