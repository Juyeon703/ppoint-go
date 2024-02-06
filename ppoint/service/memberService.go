package service

import (
	"ppoint/dto"
	"ppoint/query"
)

func MemberAdd(dbconn *query.DbConfig, member *dto.MemberAddDto) error {
	var err error
	log := dbconn.Logue
	var memberId int
	if memberId, err = dbconn.CreateMember(member); err != nil {
		return err
	}
	log.Infof("[Create]CreateMember() // param{member : %v} // result{memberId : %d}", member, memberId)
	return nil
}

func MemberUpdate(dbconn *query.DbConfig, updateMember *dto.MemberUpdateDto, originTotalPoint int) error {
	var err error
	log := dbconn.Logue
	if updateMember.TotalPoint != originTotalPoint {
		if err = PointEdit(dbconn, updateMember.MemberId, originTotalPoint, updateMember.TotalPoint); err != nil {
			return err
		}
	}
	if err = dbconn.UpdateMemberByDto(updateMember); err != nil {
		return err
	}
	log.Infof("[Update]UpdateMemberByDto() // param{member : %v}", updateMember)
	return nil
}

func FindMemberList(dbconn *query.DbConfig, search string) ([]dto.MemberDto, error) {
	var err error
	log := dbconn.Logue
	var memberList []dto.MemberDto

	if search != "" {
		if memberList, err = dbconn.SelectMemberSearch(search); err != nil {
			return nil, err
		}
		log.Debugf("[Select]SelectMemberSearch() // param{search : %s}", search)
	} else {
		if memberList, err = dbconn.SelectMembersDto(); err != nil {
			return nil, err
		}
		log.Debug("[Select]SelectMembersDto()")
	}
	return memberList, nil
}

func FindMember(dbconn *query.DbConfig, name, phoneNumber string) (*dto.MemberDto, error) {
	var err error
	log := dbconn.Logue
	var member = new(dto.MemberDto)
	if member, err = dbconn.SelectMemberByPhoneAndName(phoneNumber, name); err != nil {
		return nil, err
	}
	log.Debugf("[Select]SelectMemberByPhoneAndName() // param{phoneNumber : %s, name : %s}", phoneNumber, name)
	return member, nil
}

func FindMemberPhoneNumber(dbconn *query.DbConfig, phoneNumber string) (*dto.MemberDto, error) {
	var err error
	log := dbconn.Logue
	var member = new(dto.MemberDto)
	if member, err = dbconn.SelectMemberByPhone(phoneNumber); err != nil {
		return nil, err
	}
	log.Debugf("[Select]SelectMemberByPhone() // param{phoneNumber : %s}", phoneNumber)
	return member, nil
}

func FindUpdateMemberPhoneNumber(dbconn *query.DbConfig, phoneNumber, memberId string) (*dto.MemberDto, error) {
	var err error
	log := dbconn.Logue
	var member = new(dto.MemberDto)
	if member, err = dbconn.SelectUpdateMemberByPhone(phoneNumber, memberId); err != nil {
		return nil, err
	}
	log.Debugf("[Select]SelectUpdateMemberByPhone() // param{memberId : %s, phoneNumber : %s}", memberId, phoneNumber)
	return member, nil
}

func FindSumSalesOfMember(dbconn *query.DbConfig, memberId int) (*dto.MemberSumSalesDto, error) {
	var err error
	log := dbconn.Logue
	var result = new(dto.MemberSumSalesDto)
	if result, err = dbconn.SelectTotalSalesByMember(memberId); err != nil {
		return nil, err
	}
	log.Debugf("[Select]SelectTotalSalesByMember() // param{memberId : %d}", memberId)

	return result, nil
}

func ChangePointNoVisitFor3Month(dbconn *query.DbConfig) error {
	var err error
	log := dbconn.Logue
	if err = dbconn.CreateRevenueChangePointNoVisitFor3Month(); err != nil {
		return err
	}
	log.Info("[Create]CreateRevenueChangePointNoVisitFor3Month()")
	if err = dbconn.Update0PointMemberNoVisitFor3Month(); err != nil {
		return err
	}
	log.Info("[Update]Update0PointMemberNoVisitFor3Month()")
	return nil
}
