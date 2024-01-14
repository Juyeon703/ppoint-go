package query

import "ppoint/types"

func (dbc *DbConfig) CreateGrade(gradeName string) error {
	_, err := dbc.DbConnection.Exec("INSERT INTO `ppoint`.`grade` (`grade_name`) VALUES (?);", gradeName)
	return err
}
func (dbc *DbConfig) DeleteGrade(id int) error {
	_, err := dbc.DbConnection.Exec("DELETE FROM `ppoint`.`grade` WHERE grade_id=?;", id)
	return err
}

func (dbc *DbConfig) SelectGrades() ([]types.Grade, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`grade`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []types.Grade
	for rows.Next() {
		var grade types.Grade
		if err = rows.Scan(&grade.GradeId, &grade.GradeName); err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return grades, nil
}

func (dbc *DbConfig) SelectGradeByTotalSales(totalSales int) (*types.Grade, error) {
	var grade types.Grade
	err := dbc.DbConnection.QueryRow("SELECT * FROM ppoint.grade WHERE grade_value<=? order by grade_value DESC Limit 1", totalSales).Scan(&grade.GradeId, &grade.GradeName, &grade.GradeValue)
	if err != nil {
		return nil, err
	}
	return &grade, nil
}
