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

	if revenueDto.Sales == revenueDto.SubPoint {
		revenueDto.PayType = "포인트"
	} else {
		if revenueDto.PayType == types.Card {
			if result, err = dbconn.SelectSettingBySettingType(types.SettingCard); err != nil {
				return err
			}
		} else if revenueDto.PayType == types.Cash {
			if result, err = dbconn.SelectSettingBySettingType(types.SettingCash); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	log.Debugf("(결제 타입 조회) 곁제 타입 :[ %s ], 적립 퍼센트 :[ %s ]", revenueDto.PayType, result+" %")

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
		log.Errorf("(포인트 적립/사용) >>>> 사용자 포인트 정보 업데이트 실패  : [%v]", err)
		return err
	}
	log.Debug("(포인트 적립/사용) >>>> 사용자 포인트 정보 업데이트")

	if err = dbconn.CreateRevenue(revenueDto.MemberId, revenueDto.Sales, revenueDto.SubPoint, addPointTemp, fixedSalesTemp, revenueDto.PayType); err != nil {
		log.Errorf("(포인트 적립/사용) >>>> 매출 정보 생성 실패 : [%v]", err)
		return err
	}

	log.Debugf("(포인트 적립/사용) >>>> memberId : [%d], sales : [%d], subPoint : [%d], addPoint : [%d], fixedSales : [%d], payType : [%s]", revenueDto.MemberId, revenueDto.Sales, revenueDto.SubPoint, addPointTemp, fixedSalesTemp, revenueDto.PayType)
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
	log.Infof("(포인트 정보 수정) >>>  memberId : [%d], subPoint : [%d], addPoint : [%d], payType : [포인트 변경]", result, subPoint, addPoint)
	return nil
}

func FindRevenueList(dbconn *query.DbConfig, startDate, endDate string, memberId int) ([]dto.RevenueDto, error) {
	var err error
	log := dbconn.Logue
	var revenueList []dto.RevenueDto

	if memberId != 0 {
		if revenueList, err = dbconn.SelectRevenuesByMember(memberId); err != nil {
			log.Debugf("(SelectRevenuesByMember) >>>  fail memberId: [%d], err : [%v]", memberId, err)
			return nil, err
		}

	} else {
		if revenueList, err = dbconn.SelectRevenuesByCustomDate(startDate, endDate); err != nil {
			log.Debugf("(SelectRevenuesByCustomDate) >>>  fail startDate : [%s], endDate : [%s]", startDate, endDate)
			return nil, err
		}
	}

	log.Debugf("(매출 정보 조회) >>> startDate : [%s], endDate : [%s], memberId :[%d]", startDate, endDate, memberId)
	return revenueList, nil
}

func FindSumSalesPoint(dbconn *query.DbConfig, startDate, endDate string, memberId int) (*dto.SumSalesPointDto, error) {
	var err error
	log := dbconn.Logue
	var result *dto.SumSalesPointDto

	if memberId != 0 {
		if result, err = dbconn.SelectSumSalesPointByMemberId(memberId); err != nil {
			log.Debugf("(SelectSumSalesPointByMemberId) >>>  fail memberId: [%d], err : [%v]", memberId, err)
			return nil, err
		}
	} else {
		if result, err = dbconn.SelectSumSalesPointByCustomDate(startDate, endDate); err != nil {
			log.Debugf("(SelectSumSalesPointByCustomDate) >>>  fail startDate : [%s], endDate : [%s]", startDate, endDate)
			return nil, err
		}
	}

	log.Debugf("(매출 정보 SUM 조회) >>> startDate : [%s], endDate : [%s], memberId :[%d]", startDate, endDate, memberId)
	return result, nil
}

func RevenueDelete(dbconn *query.DbConfig, revenueId int, memberId int, subPoint int, addPoint int) error {
	var err error
	log := dbconn.Logue

	if subPoint == 0 { // 포인트 사용 안했을시 -> 적립금 o
		if err = dbconn.UpdateMemberByDelete(memberId, addPoint); err != nil {
			log.Errorf("(매출 삭제) >>>> 사용자 포인트 정보 업데이트 실패  : [%v]", err)
			return err
		}
		if err = dbconn.DeleteRevenue(revenueId); err != nil {
			log.Errorf("(매출 삭제) >>>> 매출 삭제 실패  : [%v]", err)
			return err
		}
	} else { // 포인트 사용했을시 -> 적립금 x
		if err = dbconn.UpdateRevenue(revenueId, "결제 취소"); err != nil {
			log.Errorf("(매출 삭제) >>>> 매출 정보 업데이트 실패 : [%v]", err)
			return err
		}
	}

	log.Debugf("(매출 삭제) >>>> revenueId : [%d], memberId : [%d], subPoint : [%d], addPoint : [%d]", revenueId, memberId, subPoint, addPoint)
	return nil
}
