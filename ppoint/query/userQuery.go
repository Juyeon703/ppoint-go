package query

import (
	"mysql01/types"
	"mysql01/utils"
)

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

func (dbc *DbConfig) SelectUsers() ([]types.Users, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`users`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.Users
	for rows.Next() {
		var user types.Users
		if err = rows.Scan(&user.UserId, &user.StoreName, &user.UserName, &user.PhoneNumber, &user.Id, &user.Password, &user.Email, &user.CreateDate, &user.UpdateDate); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
