package service

import (
	"fmt"
	"ppoint/dto"
	"ppoint/query"
)

func MemberAdd(dbconn *query.DbConfig, member *dto.MemberAddDto) error {
	var err error
	var memberId int
	if memberId, err = dbconn.CreateMember(member); err != nil {
		return err
	}
	fmt.Println("======> CreateMember() 호출")
	fmt.Println("등록된 회원 번호 : ", memberId)
	return nil
}

func MemberUpdate(dbconn *query.DbConfig, updateMember *dto.MemberUpdateDto, originTotalPoint int) error {
	var err error
	if updateMember.TotalPoint != originTotalPoint {
		if err = PointEdit(dbconn, updateMember.MemberId, originTotalPoint, updateMember.TotalPoint); err != nil {
			return err
		}
	}
	if err = dbconn.UpdateMemberByDto(updateMember); err != nil {
		return err
	}
	fmt.Println("=============> UpdateMemberByDto() 호출")
	return nil
}

func FindMemberList(dbconn *query.DbConfig, search string) ([]dto.MemberDto, error) {
	var err error
	var memberList []dto.MemberDto

	if search != "" {
		if memberList, err = dbconn.SelectMemberSearch(search); err != nil {
			return nil, err
		} else {
			fmt.Println("=============> SelectMemberSearch() 호출")
		}
	} else {
		if memberList, err = dbconn.SelectMembersDto(); err != nil {
			return nil, err
		} else {
			fmt.Println("=============> SelectMembersDto() 호출")
		}
	}
	return memberList, nil
}

func FindMember(dbconn *query.DbConfig, name, phoneNumber string) (*dto.MemberDto, error) {
	var err error
	var member = new(dto.MemberDto)
	if member, err = dbconn.SelectMemberByPhoneAndName(phoneNumber, name); err != nil {
		return nil, err
	}

	return member, nil
}

func FindMemberPhoneNumber(dbconn *query.DbConfig, phoneNumber string) (*dto.MemberDto, error) {
	var err error
	var member = new(dto.MemberDto)
	if member, err = dbconn.SelectMemberByPhone(phoneNumber); err != nil {
		return nil, err
	}

	return member, nil
}

func FindUpdateMemberPhoneNumber(dbconn *query.DbConfig, phoneNumber, memberId string) (*dto.MemberDto, error) {
	var err error
	var member = new(dto.MemberDto)
	if member, err = dbconn.SelectUpdateMemberByPhone(phoneNumber, memberId); err != nil {
		return nil, err
	}

	return member, nil
}

func FindSumSalesOfMember(dbconn *query.DbConfig, memberId int) (*dto.MemberSumSalesDto, error) {
	var err error
	var result = new(dto.MemberSumSalesDto)
	if result, err = dbconn.SelectTotalSalesByMember(memberId); err != nil {
		return nil, err
	}
	fmt.Println("=============> SelectTotalSalesByMember() 호출")

	return result, nil
}

func ChangePointNoVisitFor3Month(dbconn *query.DbConfig) error {
	var err error
	if err = dbconn.CreateRevenueChangePointNoVisitFor3Month(); err != nil {
		return err
	}
	fmt.Println("=============> CreateRevenueChangePointNoVisitFor3Month() 호출")
	if err = dbconn.Update0PointMemberNoVisitFor3Month(); err != nil {
		return err
	}
	fmt.Println("=============> Update0PointMemberNoVisitFor3Month() 호출")
	return nil
}
