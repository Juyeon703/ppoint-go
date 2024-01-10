package types

import (
	"mysql01/utils"
)

func (dbc *DbConfig) CreateMember(member *Member) (int, error) {
	result, err := dbc.DbConnection.Exec("INSERT INTO `ppoint`.`member` (`member_name`, `phone_number`, `birth`) VALUES (?, ?, ?);", member.MemberName, member.PhoneNumber, member.Birth)
	memberId, err := result.LastInsertId()
	return int(memberId), err
}

func (dbc *DbConfig) UpdateMemberByPhoneNumber(id int, newphoneNumber, newName string) error {
	_, err := dbc.DbConnection.Exec("UPDATE `ppoint`.`member` SET phone_number=?, member_name=?, update_date=? WHERE member_id=?", newphoneNumber, newName, utils.CurrentTime(), id)
	return err
}

func (dbc *DbConfig) UpdateMemberByPoint(id, updatePoint int) error {
	_, err := dbc.DbConnection.Exec("UPDATE `ppoint`.`member` SET total_point=?, update_date=? WHERE member_id=?", updatePoint, utils.CurrentTime(), id)
	return err
}

func (dbc *DbConfig) UpdateMemberByGrade(id, newGradeId int) error {
	_, err := dbc.DbConnection.Exec("UPDATE `ppoint`.`member` SET grade_id=? WHERE member_id=?", newGradeId, id)
	return err
}

func (dbc *DbConfig) SelectMembers() ([]Member, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`member`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member
		if err = rows.Scan(&member.MemberId, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.GradeId, &member.TotalPoint, &member.CreateDate, &member.UpdateDate); err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}

func (dbc *DbConfig) SelectMemberByPhoneNumber(phoneNumber string) (*Member, error) {
	var member Member
	err := dbc.DbConnection.QueryRow("SELECT * FROM ppoint.member WHERE phone_number=?", phoneNumber).
		Scan(&member.MemberId, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.GradeId, &member.TotalPoint, &member.CreateDate, &member.UpdateDate)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (dbc *DbConfig) SelectMemberSearch(search string) ([]Member, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM ppoint.member WHERE phone_number LIKE ? or member_name like ?;", "%"+search, "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member
		if err = rows.Scan(&member.MemberId, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.GradeId, &member.TotalPoint, &member.CreateDate, &member.UpdateDate); err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}

func (dbc *DbConfig) SelectMemberByMemberId(memberId int) (*MemberDto, error) {
	var member MemberDto
	err := dbc.DbConnection.QueryRow("SELECT member.member_id, grade.grade_name, member.member_name, member.phone_number, member.birth, member.total_point, member.create_date, member.update_date FROM `ppoint`.`member` join `ppoint`.`grade` on member.grade_id = grade.grade_id WHERE member_id=?", memberId).
		Scan(&member.MemberId, &member.GradeName, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.TotalPoint, &member.CreateDate, &member.UpdateDate)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (dbc *DbConfig) SelectMembersOrderByGrade() ([]MemberDto, error) {
	rows, err := dbc.DbConnection.Query("SELECT member.member_id, grade.grade_name, member.member_name, member.phone_number, member.birth, member.total_point, member.create_date, member.update_date FROM `ppoint`.`member` join `ppoint`.`grade` on member.grade_id = grade.grade_id order by member.grade_id DESC, member_name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []MemberDto
	for rows.Next() {
		var member MemberDto
		if err = rows.Scan(&member.MemberId, &member.GradeName, &member.MemberName, &member.PhoneNumber, &member.Birth, &member.TotalPoint, &member.CreateDate, &member.UpdateDate); err != nil {
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

// user
func (dbc *DbConfig) CreateUser(storeName, userName, phoneNumber, id, password, email string) error {
	_, err := dbc.DbConnection.Exec("INSERT INTO `ppoint`.`users` (`store_name`, `user_name`, `phone_number`, `id`, `password`, `email`) VALUES (?, ?, ?, ?, ?, ?);", storeName, userName, phoneNumber, id, password, email)
	return err
}
func (dbc *DbConfig) UpdateUserByPhoneNumber(id int, newPhoneNumber, newstoreName string) error {
	_, err := dbc.DbConnection.Exec("UPDATE `ppoint`.`users` SET phone_number=?, store_name=?, update_date=? WHERE user_id=?", newPhoneNumber, newstoreName, utils.CurrentTime(), id)
	return err
}

func (dbc *DbConfig) DeleteUser(id int) error {
	_, err := dbc.DbConnection.Exec("DELETE FROM `ppoint`.`users` WHERE user_id=?;", id)
	return err
}

func (dbc *DbConfig) SelectUsers() ([]Users, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`users`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []Users
	for rows.Next() {
		var user Users
		if err = rows.Scan(&user.UserId, &user.StoreName, &user.UserName, &user.PhoneNumber, &user.Id, &user.Password, &user.Email, &user.CreateDate, &user.UpdateDate); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// revenue
func (dbc *DbConfig) CreateRevenue(revenue *Revenue) error {
	_, err := dbc.DbConnection.Exec("INSERT INTO `ppoint`.`revenue` (`member_id`, `sales`, `sub_point`, `add_point`, `fixed_sales`, `pay_type`) VALUES (?, ?, ?, ?, ?, ?);",
		revenue.MemberId, revenue.Sales, revenue.SubPoint, revenue.AddPoint, revenue.FixedSales, revenue.PayType)
	return err
}
func (dbc *DbConfig) DeleteRevenue(id int) error {
	_, err := dbc.DbConnection.Exec("DELETE FROM `ppoint`.`revenue` WHERE revenue_id=?;", id)
	return err
}

func (dbc *DbConfig) SelectRevenues() ([]Revenue, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`revenue`;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var revenues []Revenue
	for rows.Next() {
		var revenue Revenue
		if err = rows.Scan(&revenue.RevenueId, &revenue.MemberId, &revenue.Sales, &revenue.SubPoint, &revenue.AddPoint, &revenue.FixedSales, &revenue.PayType, &revenue.CreateDate); err != nil {
			return nil, err
		}
		revenues = append(revenues, revenue)
	}

	return revenues, nil
}

func (dbc *DbConfig) SelectRevenuesByToday(today string) ([]Revenue, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`revenue` where date_format(create_date, '%Y-%m-%d') = ? order by create_date;", today)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var revenues []Revenue
	for rows.Next() {
		var revenue Revenue
		if err = rows.Scan(&revenue.RevenueId, &revenue.MemberId, &revenue.Sales, &revenue.SubPoint, &revenue.AddPoint, &revenue.FixedSales, &revenue.PayType, &revenue.CreateDate); err != nil {
			return nil, err
		}
		revenues = append(revenues, revenue)
	}
	return revenues, nil
}

func (dbc *DbConfig) SelectRevenuesByCustomDate(startDate, endDate string) ([]Revenue, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`revenue` where date_format(create_date, '%Y-%m-%d') BETWEEN ? and ? order by create_date DESC;", startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var revenues []Revenue
	for rows.Next() {
		var revenue Revenue
		if err = rows.Scan(&revenue.RevenueId, &revenue.MemberId, &revenue.Sales, &revenue.SubPoint, &revenue.AddPoint, &revenue.FixedSales, &revenue.PayType, &revenue.CreateDate); err != nil {
			return nil, err
		}
		revenues = append(revenues, revenue)
	}
	return revenues, nil
}

// 실제 사용금액으로 할건지,, 총 결제금액으로 할건지
func (dbc *DbConfig) SelectTotalSalesByMember(memberId int) (*Total, error) {
	var total Total
	err := dbc.DbConnection.QueryRow("SELECT SUM(fixed_sales), member.grade_id, grade.grade_name FROM ppoint.revenue join ppoint.member on revenue.member_id = member.member_id join grade on member.grade_id = grade.grade_id WHERE revenue.member_id=?", memberId).Scan(&total.TotalSales, &total.GradeId, &total.GradeName)
	if err != nil {
		return nil, err
	}
	return &total, nil
}

// grade
func (dbc *DbConfig) CreateGrade(gradeName string) error {
	_, err := dbc.DbConnection.Exec("INSERT INTO `ppoint`.`grade` (`grade_name`) VALUES (?);", gradeName)
	return err
}
func (dbc *DbConfig) DeleteGrade(id int) error {
	_, err := dbc.DbConnection.Exec("DELETE FROM `ppoint`.`grade` WHERE grade_id=?;", id)
	return err
}

func (dbc *DbConfig) SelectGrades() ([]Grade, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`grade`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []Grade
	for rows.Next() {
		var grade Grade
		if err = rows.Scan(&grade.GradeId, &grade.GradeName); err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return grades, nil
}

func (dbc *DbConfig) SelectGradeByTotalSales(totalSales int) (*Grade, error) {
	var grade Grade
	err := dbc.DbConnection.QueryRow("SELECT * FROM ppoint.grade WHERE grade_value<=? order by grade_value DESC Limit 1", totalSales).Scan(&grade.GradeId, &grade.GradeName, &grade.GradeValue)
	if err != nil {
		return nil, err
	}
	return &grade, nil
}

// setting
func (dbc *DbConfig) CreateSetting(settingName, settingValue string) error {
	_, err := dbc.DbConnection.Exec("INSERT INTO `ppoint`.`setting` (`setting_name`, `setting_value`) VALUES (?, ?);", settingName, settingValue)
	return err
}
func (dbc *DbConfig) UpdateSettingById(id int, newName, newValue string) error {
	_, err := dbc.DbConnection.Exec("UPDATE `ppoint`.`setting` SET setting_name=?, setting_value=? WHERE setting_id=?", newName, newValue, id)
	return err
}
func (dbc *DbConfig) DeleteSetting(id int) error {
	_, err := dbc.DbConnection.Exec("DELETE FROM `ppoint`.`setting` WHERE setting_id=?;", id)
	return err
}

func (dbc *DbConfig) SelectSettings() ([]Setting, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`setting`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var settings []Setting
	for rows.Next() {
		var setting Setting
		if err = rows.Scan(&setting.SettingId, &setting.SettingName, &setting.SettingValue); err != nil {
			return nil, err
		}
		settings = append(settings, setting)
	}

	return settings, nil
}

func (dbc *DbConfig) SelectSettingByPayType(payType string) (*Setting, error) {
	var setting Setting
	err := dbc.DbConnection.QueryRow("SELECT * FROM ppoint.setting WHERE setting_type='결제 방법' And setting_name=?", payType).
		Scan(&setting.SettingId, &setting.SettingType, &setting.SettingName, &setting.SettingValue, &setting.SettingDescription)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}
