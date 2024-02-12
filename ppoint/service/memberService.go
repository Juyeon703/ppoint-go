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
	log.Debugf("(고객 등록) >>> 고객 입력 정보 : [%v],  memberId : [%d]", member, memberId)
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
	log.Infof("(고객 정보 수정) >>> 고객 수정 정보 : [%v]", updateMember)
	return nil
}

func FindMemberList(dbconn *query.DbConfig, search string) ([]dto.MemberDto, error) {
	var err error
	log := dbconn.Logue
	var memberList []dto.MemberDto

	if search != "" {
		if memberList, err = dbconn.SelectMemberSearch(search); err != nil {
			log.Errorf("(포인트 페이지) >>> 고객 조회 실패 [%v]", err)
			return nil, err
		}
		log.Debugf("(포인트 페이지) >>> 고객 조회 검색어 : [%s]", search)
	} else {
		if memberList, err = dbconn.SelectMembersDto(); err != nil {
			log.Debugf("(포인트 페이지) >>> 조회 고객 정보 출력 실패 : [%v]", err)
			return nil, err
		}
		//log.Debugf("(포인트 페이지) >>> 조회 고객 정보 출력 : [%v]", memberList)
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
	log.Debugf("(고객 정보 찾기) >>> 핸드폰 번호 : [%s], 이름 : [%s]", phoneNumber, name)
	return member, nil
}

func FindMemberPhoneNumber(dbconn *query.DbConfig, phoneNumber string) (*dto.MemberDto, error) {
	var err error
	log := dbconn.Logue
	var member = new(dto.MemberDto)
	if member, err = dbconn.SelectMemberByPhone(phoneNumber); err != nil {
		return nil, err
	}
	log.Debugf("(고객 정보 찾기) >>> 핸드폰 번호 : [%s]", phoneNumber)
	return member, nil
}

func FindUpdateMemberPhoneNumber(dbconn *query.DbConfig, phoneNumber, memberId string) (*dto.MemberDto, error) {
	var err error
	log := dbconn.Logue
	var member = new(dto.MemberDto)
	if member, err = dbconn.SelectUpdateMemberByPhone(phoneNumber, memberId); err != nil {
		return nil, err
	}
	log.Debugf("(고객 정보 수정) >>>  핸드폰 번호 : [%s], memberId : [%s]", phoneNumber, memberId)
	return member, nil
}

func FindSumSalesOfMember(dbconn *query.DbConfig, memberId int) (*dto.MemberSumSalesDto, error) {
	var err error
	var result = new(dto.MemberSumSalesDto)
	if result, err = dbconn.SelectTotalSalesByMember(memberId); err != nil {
		return nil, err
	}

	return result, nil
}

func ChangePointNoVisitFor3Month(dbconn *query.DbConfig) error {
	var err error
	log := dbconn.Logue
	if err = dbconn.CreateRevenueChangePointNoVisitFor3Month(); err != nil {
		log.Errorf("(ChangePointNoVisitFor3Month) 포인트 소멸 데이터 추가 실패 >>> %v", err)
		return err
	}

	if err = dbconn.Update0PointMemberNoVisitFor3Month(); err != nil {
		log.Errorf("(ChangePointNoVisitFor3Month) >>> 멤버 업데이트 실패 >>> %v", err)
		return err
	}

	return nil
}
