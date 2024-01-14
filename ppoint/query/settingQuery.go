package query

import "mysql01/types"

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

func (dbc *DbConfig) SelectSettings() ([]types.Setting, error) {
	rows, err := dbc.DbConnection.Query("SELECT * FROM `ppoint`.`setting`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var settings []types.Setting
	for rows.Next() {
		var setting types.Setting
		if err = rows.Scan(&setting.SettingId, &setting.SettingName, &setting.SettingValue); err != nil {
			return nil, err
		}
		settings = append(settings, setting)
	}

	return settings, nil
}

func (dbc *DbConfig) SelectSettingByPayType(payType string) (*types.Setting, error) {
	var setting types.Setting
	err := dbc.DbConnection.QueryRow("SELECT * FROM ppoint.setting WHERE setting_type='결제 방법' And setting_name=?", payType).
		Scan(&setting.SettingId, &setting.SettingType, &setting.SettingName, &setting.SettingValue, &setting.SettingDescription)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}