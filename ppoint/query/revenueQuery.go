package query

import (
	"mysql01/dto"
	"mysql01/types"
)

func (dbc *DbConfig) CreateRevenue(revenue *types.Revenue) error {
	_, err := dbc.DbConnection.Exec("INSERT INTO `ppoint`.`revenue` (`member_id`, `sales`, `sub_point`, `add_point`, `fixed_sales`, `pay_type`) VALUES (?, ?, ?, ?, ?, ?);",
		revenue.MemberId, revenue.Sales, revenue.SubPoint, revenue.AddPoint, revenue.FixedSales, revenue.PayType)
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

func (dbc *DbConfig) SelectRevenuesByCustomDate(startDate, endDate string) ([]types.Revenue, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`revenue` where date_format(create_date, '%Y-%m-%d') BETWEEN ? and ? order by create_date DESC;", startDate, endDate)
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

// 실제 사용금액으로 할건지,, 총 결제금액으로 할건지
func (dbc *DbConfig) SelectTotalSalesByMember(memberId int) (*dto.MemberSalesDto, error) {
	var total dto.MemberSalesDto
	err := dbc.DbConnection.QueryRow("SELECT SUM(fixed_sales), member.grade_id, grade.grade_name FROM ppoint.revenue join ppoint.member on revenue.member_id = member.member_id join grade on member.grade_id = grade.grade_id WHERE revenue.member_id=?", memberId).Scan(&total.TotalSales, &total.GradeId, &total.GradeName)
	if err != nil {
		return nil, err
	}
	return &total, nil
}
