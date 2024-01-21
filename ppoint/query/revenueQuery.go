package query

import (
	"ppoint/dto"
	"ppoint/types"
)

func (dbc *DbConfig) CreateRevenue(memberId, sales, subPoint, addPoint, fixedSales int, payType string) error {
	_, err := dbc.DbConnection.Exec("INSERT INTO `ppoint`.`revenue` (`member_id`, `sales`, `sub_point`, `add_point`, `fixed_sales`, `pay_type`) VALUES (?, ?, ?, ?, ?, ?);",
		memberId, sales, subPoint, addPoint, fixedSales, payType)
	return err
}
func (dbc *DbConfig) DeleteRevenue(id int) error {
	_, err := dbc.DbConnection.Exec("DELETE FROM `ppoint`.`revenue` WHERE revenue_id=?;", id)
	return err
}

func (dbc *DbConfig) SelectRevenues() ([]types.Revenue, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`revenue`;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var revenues []types.Revenue
	for rows.Next() {
		var revenue types.Revenue
		if err = rows.Scan(&revenue.RevenueId, &revenue.MemberId, &revenue.Sales, &revenue.SubPoint, &revenue.AddPoint, &revenue.FixedSales, &revenue.PayType, &revenue.CreateDate); err != nil {
			return nil, err
		}
		revenues = append(revenues, revenue)
	}

	return revenues, nil
}

func (dbc *DbConfig) SelectRevenuesByToday(today string) ([]types.Revenue, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`revenue` where date_format(create_date, '%Y-%m-%d') = ? order by create_date;", today)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var revenues []types.Revenue
	for rows.Next() {
		var revenue types.Revenue
		if err = rows.Scan(&revenue.RevenueId, &revenue.MemberId, &revenue.Sales, &revenue.SubPoint, &revenue.AddPoint, &revenue.FixedSales, &revenue.PayType, &revenue.CreateDate); err != nil {
			return nil, err
		}
		revenues = append(revenues, revenue)
	}
	return revenues, nil
}

func (dbc *DbConfig) SelectRevenuesByCustomDate(startDate, endDate string) ([]dto.RevenueDto, error) {
	rows, err := dbc.DbConnection.Query("SELECT revenue.revenue_id, revenue.member_id, member.member_name, member.phone_number, revenue.sales, revenue.sub_point, revenue.add_point, revenue.fixed_sales, revenue.pay_type, revenue.create_date FROM `ppoint`.`revenue` join `ppoint`.`member` on revenue.member_id = member.member_id where date_format(revenue.create_date, '%Y-%m-%d') BETWEEN ? and ? order by revenue.create_date DESC;", startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var revenues []dto.RevenueDto
	for rows.Next() {
		var revenue dto.RevenueDto
		if err = rows.Scan(&revenue.RevenueId, &revenue.MemberId, &revenue.MemberName, &revenue.PhoneNumber,
			&revenue.Sales, &revenue.SubPoint, &revenue.AddPoint, &revenue.FixedSales, &revenue.PayType, &revenue.CreateDate); err != nil {
			return nil, err
		}
		revenues = append(revenues, revenue)
	}
	return revenues, nil
}

func (dbc *DbConfig) SelectRevenuesByMember(memberId int) ([]dto.RevenueDto, error) {
	rows, err := dbc.DbConnection.Query("SELECT revenue.revenue_id, revenue.member_id, member.member_name, member.phone_number, revenue.sales, revenue.sub_point, revenue.add_point, revenue.fixed_sales, revenue.pay_type, revenue.create_date FROM `ppoint`.`revenue` join `ppoint`.`member` on revenue.member_id = member.member_id where revenue.member_id=? order by revenue.create_date DESC;", memberId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var revenues []dto.RevenueDto
	for rows.Next() {
		var revenue dto.RevenueDto
		if err = rows.Scan(&revenue.RevenueId, &revenue.MemberId, &revenue.MemberName, &revenue.PhoneNumber,
			&revenue.Sales, &revenue.SubPoint, &revenue.AddPoint, &revenue.FixedSales, &revenue.PayType, &revenue.CreateDate); err != nil {
			return nil, err
		}
		revenues = append(revenues, revenue)
	}
	return revenues, nil
}

func (dbc *DbConfig) SelectTotalSalesByMember(memberId int) (*dto.MemberSumSalesDto, error) {
	var total dto.MemberSumSalesDto
	err := dbc.DbConnection.QueryRow("SELECT SUM(Sales), SUM(add_point) FROM ppoint.revenue WHERE revenue.member_id=?", memberId).
		Scan(&total.TotalSales, &total.TotalPoint)
	if err != nil {
		return &total, err
	}
	return &total, nil
}
