package backup

import (
	"errors"
	"fmt"
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
		log.Fatalln(err)
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	var bFile *os.File
	backupFileFullPath := filepath.Join(dbBackupPath, currentDate.Format("2006_01_02"))
	if bFile, err = os.OpenFile(backupFileFullPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0777)); err != nil {
		//file create fail
		return
	}
	defer bFile.Close()

	//////////////////////////////////////////////
	///고객
	//////////////////////////////////////////////
	_, err = bFile.Write([]byte(prefixMemStr))
	if err != nil {
		fmt.Println("file write fail")
		return
	}

	for idx, mem := range memberList {
		mStr := fmt.Sprintf("(%d,'%s','%s','%s',%d,%d,'%s','%s')", mem.MemberId, mem.MemberName, mem.PhoneNumber, mem.Birth, mem.TotalPoint, mem.VisitCount, mem.CreateDate, mem.UpdateDate)
		if idx == len(memberList)-1 {
			mStr += ";"
		} else {
			mStr += ", \n"
		}

		_, err = bFile.Write([]byte(mStr))
		if err != nil {
			fmt.Println("file write fail")
			return
		}
	}

	//////////////////////////////////////////////
	///매출
	//////////////////////////////////////////////
	_, err = bFile.Write([]byte(prefixRvnStr))
	if err != nil {
		fmt.Println("file write fail")
		return
	}

	if revenueList, err = DbConf.SelectRevenues(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
	for idx, rvn := range revenueList {
		rStr := fmt.Sprintf("( %d, %d, %d, %d, %d, '%s','%s')", rvn.MemberId, rvn.Sales, rvn.SubPoint, rvn.AddPoint, rvn.FixedSales, rvn.PayType, rvn.CreateDate)

		if idx == len(revenueList)-1 {
			rStr += ";"
		} else {
			rStr += ", \n"
		}

		_, err = bFile.Write([]byte(rStr))
		if err != nil {
			fmt.Println("file write fail")
			return
		}
	}

	//////////////////////////////////////////////
	///설정
	//////////////////////////////////////////////
	_, err = bFile.Write([]byte(prefixSettStr))
	if err != nil {
		fmt.Println("file write fail")
		return
	}

	if settingList, err = DbConf.SelectSettings(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
	for idx, sett := range settingList {
		sStr := fmt.Sprintf("('%s','%s','%s')", sett.SettingType, sett.SettingValue, sett.SettingDescription)

		if idx == len(settingList)-1 {
			sStr += ";"
		} else {
			sStr += ", \n"
		}

		_, err = bFile.Write([]byte(sStr))
		if err != nil {
			fmt.Println("file write fail")
			return
		}
	}

	bFile.Close()

}

func clearOldBackupFile(backupPath string, pastDate time.Time) {
	var err error
	var backupFileFullPath string

	backupFileFullPath = filepath.Join(backupPath, pastDate.Format("2006_01_02"))

	//batch file 존재 여부 확인
	if _, err = os.Stat(backupFileFullPath); err == nil {
		log.Println("file 존재")
		if err = os.Remove(backupFileFullPath); err != nil {
			return
		}

		log.Println("오래된 backup 파일 정리 완료")
	} else {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("file 없음")
			return

		} else {
			log.Fatalln(err.Error())
			panic(err)
		}
	}
}
