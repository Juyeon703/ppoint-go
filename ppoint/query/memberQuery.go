package query

import (
	"ppoint/dto"
	"ppoint/types"
	"ppoint/utils"
)

func (dbc *DbConfig) CreateMember(member *dto.MemberAddDto) (int, error) {
	result, err := dbc.DbConnection.Exec("INSERT INTO `ppoint`.`member` (`member_name`, `phone_number`, `birth`) VALUES (?, ?, ?);", member.MemberName, member.PhoneNumber, member.Birth.Format("2006-01-02"))
	memberId, err := result.LastInsertId()
	return int(memberId), err
}

func (dbc *DbConfig) UpdateMemberByDto(updateMember *dto.MemberUpdateDto) error {
	_, err := dbc.DbConnection.Exec("UPDATE `ppoint`.`member` SET member_name=?, phone_number=?, birth=?, total_point=?, visit_count=?, update_date=? WHERE member_id=?",
		updateMember.MemberName, updateMember.PhoneNumber, updateMember.Birth, updateMember.TotalPoint, updateMember.VisitCount, utils.CurrentTime(), updateMember.MemberId)
	return err
}

func (dbc *DbConfig) UpdateMemberByPoint(id, changePoint int) error {
	_, err := dbc.DbConnection.Exec("UPDATE `ppoint`.`member` SET total_point=total_point+?, visit_count=visit_count + 1, update_date=? WHERE member_id=?", changePoint, utils.CurrentTime(), id)
	return err
}

// test 2일 -- 수정 필요
func (dbc *DbConfig) Update0PointMemberNoVisitFor3Month() error {
	_, err := dbc.DbConnection.Exec("UPDATE ppoint.member SET total_point=0 WHERE total_point != 0 and update_date <= DATE_ADD(now(), INTERVAL -2 DAY)")
	return err
}

func (dbc *DbConfig) SelectMembers() ([]types.Member, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`member`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []types.Member
	for rows.Next() {
		var member types.Member
		if err = rows.Scan(&member.MemberId, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.TotalPoint, &member.VisitCount, &member.CreateDate, &member.UpdateDate); err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}

func (dbc *DbConfig) SelectMemberByPhone(phoneNumber string) (*dto.MemberDto, error) {
	var member dto.MemberDto
	err := dbc.DbConnection.QueryRow("SELECT * FROM ppoint.member WHERE phone_number=?", phoneNumber).
		Scan(&member.MemberId, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.TotalPoint, &member.VisitCount, &member.CreateDate, &member.UpdateDate)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (dbc *DbConfig) SelectUpdateMemberByPhone(phoneNumber, memberId string) (*dto.MemberDto, error) {
	var member dto.MemberDto
	err := dbc.DbConnection.QueryRow("SELECT * FROM ppoint.member WHERE phone_number=? AND member_id != ?", phoneNumber, memberId).
		Scan(&member.MemberId, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.TotalPoint, &member.VisitCount, &member.CreateDate, &member.UpdateDate)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (dbc *DbConfig) SelectMemberByPhoneAndName(phoneNumber string, memberName string) (*dto.MemberDto, error) {
	var member dto.MemberDto
	err := dbc.DbConnection.QueryRow("SELECT * FROM ppoint.member WHERE phone_number=? AND member_name=?", phoneNumber, memberName).
		Scan(&member.MemberId, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.TotalPoint, &member.VisitCount, &member.CreateDate, &member.UpdateDate)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (dbc *DbConfig) SelectMemberSearch(search string) ([]dto.MemberDto, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM ppoint.member WHERE phone_number LIKE ? or member_name like ?;", "%"+search+"%", "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []dto.MemberDto
	for rows.Next() {
		var member dto.MemberDto
		if err = rows.Scan(&member.MemberId, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.TotalPoint, &member.VisitCount, &member.CreateDate, &member.UpdateDate); err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}

func (dbc *DbConfig) SelectMemberByMemberId(memberId int) (*dto.MemberDto, error) {
	var member dto.MemberDto
	err := dbc.DbConnection.QueryRow("SELECT member_id, member_name, phone_number, birth, total_point, visit_count, create_date, update_date FROM `ppoint`.`member` WHERE member_id=?", memberId).
		Scan(&member.MemberId, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.TotalPoint, &member.VisitCount, &member.CreateDate, &member.UpdateDate)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (dbc *DbConfig) SelectMembersDto() ([]dto.MemberDto, error) {
	rows, err := dbc.DbConnection.Query("SELECT member_id, member_name, phone_number, birth, total_point, visit_count, create_date, update_date FROM `ppoint`.`member`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []dto.MemberDto
	for rows.Next() {
		var member dto.MemberDto
		if err = rows.Scan(&member.MemberId, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.TotalPoint, &member.VisitCount, &member.CreateDate, &member.UpdateDate); err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}

func (dbc *DbConfig) DeleteMember(id int) error {
	_, err := dbc.DbConnection.Exec("DELETE FROM `ppoint`.`member` WHERE member_id=?;", id)
	return err
}
