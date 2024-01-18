package query

import (
	"ppoint/dto"
	"ppoint/types"
	"ppoint/utils"
)

func (dbc *DbConfig) CreateMember(member *dto.MemberAddDto) (int, error) {
	result, err := dbc.DbConnection.Exec("INSERT INTO `ppoint`.`member` (`member_name`, `phone_number`, `birth`) VALUES (?, ?, ?);", member.MemberName, member.PhoneNumber, member.Birth)
	memberId, err := result.LastInsertId()
	return int(memberId), err
}

func (dbc *DbConfig) UpdateMemberByPhoneNumber(id int, newphoneNumber, newName string) error {
	_, err := dbc.DbConnection.Exec("UPDATE `ppoint`.`member` SET phone_number=?, member_name=?, update_date=? WHERE member_id=?", newphoneNumber, newName, utils.CurrentTime(), id)
	return err
}

func (dbc *DbConfig) UpdateMemberByPoint(id, changePoint int) error {
	_, err := dbc.DbConnection.Exec("UPDATE `ppoint`.`member` SET total_point=total_point+?, visit_count=visit_count + 1, update_date=? WHERE member_id=?", changePoint, utils.CurrentTime(), id)
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
	rows, err := dbc.DbConnection.Query("SELECT * FROM ppoint.member WHERE phone_number LIKE ? or member_name like ?;", "%"+search, "%"+search+"%")
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
