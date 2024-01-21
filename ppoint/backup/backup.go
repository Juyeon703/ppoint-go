package backup

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"ppoint/query"
	"ppoint/types"
	"time"
)

func DbBackup(DbConf *query.DbConfig) {
	var err error
	var memberList []types.Member
	var revenueList []types.Revenue
	var settingList []types.Setting
	var saveStr string
	var memStr, rvnStr, SettStr string

	prefixMemStr := "INSERT INTO `member` (`member_id`, `member_name`, `phone_number`, `birth`, `total_point`, `visit_count`, `create_date`, `update_date`) VALUES "
	prefixRvnStr := "INSERT INTO `revenue` (`member_id`, `sales`, `sub_point`, `add_point`, `fixed_sales`, `pay_type`, `create_date`) VALUES "
	prefixSettStr := "INSERT INTO `ppoint`.`setting` (`setting_type`, `setting_value`, `setting_description`) VALUES "

	currentDate := time.Now()
	pastDate := currentDate.AddDate(0, -1, 0)

	dbBackupPath, err := DbConf.SelectSettingBySettingType("db_backup_path")
	if err != nil {
		log.Fatalln(err)
	}

	clearOldBackupFile(dbBackupPath, pastDate)

	if memberList, err = DbConf.SelectMembers(); err != nil {
		panic(err.Error())
	}

	for idx, mem := range memberList {
		mStr := fmt.Sprintf("(%d,'%s','%s','%s',%d,%d,'%s','%s')", mem.MemberId, mem.MemberName, mem.PhoneNumber, mem.Birth, mem.TotalPoint, mem.VisitCount, mem.CreateDate, mem.UpdateDate)
		memStr += mStr

		if idx == len(memberList)-1 {
			memStr += ";"
		} else {
			memStr += ", \n"
		}
	}

	if revenueList, err = DbConf.SelectRevenues(); err != nil {
		panic(err.Error())
	}
	for idx, rvn := range revenueList {
		rStr := fmt.Sprintf("( %d, %d, %d, %d, %d, '%s','%s')", rvn.MemberId, rvn.Sales, rvn.SubPoint, rvn.AddPoint, rvn.FixedSales, rvn.PayType, rvn.CreateDate)
		rvnStr += rStr

		if idx == len(revenueList)-1 {
			rvnStr += ";"
		} else {
			rvnStr += ", \n"
		}
	}

	if settingList, err = DbConf.SelectSettings(); err != nil {
		panic(err.Error())
	}
	for idx, sett := range settingList {
		sStr := fmt.Sprintf("('%s','%s','%s')", sett.SettingType, sett.SettingValue, sett.SettingDescription)
		SettStr += sStr

		if idx == len(settingList)-1 {
			SettStr += ";"
		} else {
			SettStr += ", \n"
		}
	}

	saveStr = prefixMemStr + memStr + "\n\n" + prefixRvnStr + rvnStr + "\n\n" + prefixSettStr + SettStr + "\n\n"
	createBackupFile(dbBackupPath, currentDate.Format("2006_01_02"), saveStr)

}

func createBackupFile(backupPath, batchFileName, inputStr string) {
	backupFileFullPath := filepath.Join(backupPath, batchFileName)
	//batch file 존재 여부 확인
	if _, err := os.Stat(backupFileFullPath); err == nil {
		fmt.Println("file 존재")
		return
	} else {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file 없음")
			b := []byte(inputStr)

			f5, err := os.Create(backupFileFullPath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			n, err := io.WriteString(f5, string(b))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("wrote %d bytes\n", n)

			fmt.Println("backup 파일 생성 완료")
		} else {
			panic(err)
		}
	}
}

func clearOldBackupFile(backupPath string, pastDate time.Time) {
	var err error
	var backupFileFullPath string

	backupFileFullPath = filepath.Join(backupPath, pastDate.Format("2006_01_02"))

	//batch file 존재 여부 확인
	if _, err = os.Stat(backupFileFullPath); err == nil {
		fmt.Println("file 존재")
		if err = os.Remove(backupFileFullPath); err != nil {
			return
		}

		fmt.Println("오래된 backup 파일 정리 완료")
	} else {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file 없음")
			return

		} else {

			panic(err)
		}
	}
}
