package query

import "ppoint/types"

func (dbc *DbConfig) CreateSetting(settingType, settingValue, settingDescription string) error {
	_, err := dbc.DbConnection.Exec("INSERT INTO `ppoint`.`setting` (`setting_type`, `setting_value`, `setting_description`) VALUES (?, ?, ?);", settingType, settingValue, settingDescription)
	return err
}
func (dbc *DbConfig) UpdateSettingById(settingType, newValue, newDescription string) error {
	_, err := dbc.DbConnection.Exec("UPDATE `ppoint`.`setting` SET setting_value=?, setting_description=? WHERE setting_type=?", newValue, newDescription, settingType)
	return err
}
func (dbc *DbConfig) DeleteSetting(settingType string) error {
	_, err := dbc.DbConnection.Exec("DELETE FROM `ppoint`.`setting` WHERE setting_type=?;", settingType)
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
		if err = rows.Scan(&setting.SettingType, &setting.SettingValue, &setting.SettingDescription); err != nil {
			return nil, err
		}
		settings = append(settings, setting)
	}

	return settings, nil
}

func (dbc *DbConfig) SelectSettingBySettingType(settingType string) (string, error) {
	var result string
	err := dbc.DbConnection.QueryRow("SELECT setting_value FROM ppoint.setting WHERE setting_type=?", settingType).
		Scan(&result)
	if err != nil {
		return "", err
	}
	return result, nil
}
