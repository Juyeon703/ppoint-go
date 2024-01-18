package query

import (
	"ppoint/dto"
	"ppoint/types"
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

func (dbc *DbConfig) SelectTotalSalesByMember(memberId int) (*dto.MemberSumSalesDto, error) {
	var total dto.MemberSumSalesDto
	err := dbc.DbConnection.QueryRow("SELECT SUM(Sales), SUM(add_point) FROM ppoint.revenue WHERE revenue.member_id=?", memberId).
		Scan(&total.TotalSales, &total.TotalPoint)
	if err != nil {
		return &total, err
	}
	return &total, nil
}
