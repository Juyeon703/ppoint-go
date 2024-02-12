package backup

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"ppoint/logue"
	"ppoint/query"
	"ppoint/types"
	"sync"
	"time"
)

var log *logue.Logbook

func DbBackup(DbConf *query.DbConfig, wg *sync.WaitGroup, backupResult chan bool) {
	var err error
	var memberList []types.Member
	var revenueList []types.Revenue
	var settingList []types.Setting
	log = DbConf.Logue

	prefixMemStr := "INSERT INTO `member` (`member_id`, `member_name`, `phone_number`, `birth`, `total_point`, `visit_count`, `create_date`, `update_date`) VALUES "
	prefixRvnStr := "INSERT INTO `revenue` (`member_id`, `sales`, `sub_point`, `add_point`, `fixed_sales`, `pay_type`, `create_date`) VALUES "
	prefixSettStr := "INSERT INTO `ppoint`.`setting` (`setting_type`, `setting_value`, `setting_description`) VALUES "

	currentDate := time.Now()
	pastDate := currentDate.AddDate(0, -1, 0)

	dbBackupPath, err := DbConf.SelectSettingBySettingType("db_backup_path")
	if err != nil {
		log.Errorf("(데이터 백업) >>> DB_BACKUP_PATH 조회 실패 : [%v]", err)
		panic(err)
	}

	log.Debug("(백업 데이터 정리) >>> START")
	log.Debugf("(백업 데이터 정리) >>> %s 이전 파일 삭제", pastDate.Format("2006_01_02"))
	clearOldBackupFile(dbBackupPath, pastDate)
	log.Debug("(백업 데이터 정리) >>> END")

	log.Debug("(데이터 백업) >>> START")
	var bFile *os.File
	backupFileFullPath := filepath.Join(dbBackupPath, currentDate.Format("2006_01_02"))
	if bFile, err = os.OpenFile(backupFileFullPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0777)); err != nil {
		//file create fail
		log.Errorf("(데이터 백업) >>> DB_BACKUP FILE OPEN 실패 : [%v]", err)
		return
	}

	//////////////////////////////////////////////
	///고객
	//////////////////////////////////////////////
	_, err = bFile.Write([]byte(prefixMemStr))
	if err != nil {
		log.Errorf("(데이터 백업) >>> DB_BACKUP FILE PREFIX MEMBER WRITE 실패 : [%v]", err)
		return
	}

	if memberList, err = DbConf.SelectMembers(); err != nil {
		log.Errorf("(데이터 백업) >>> DB_BACKUP 고객 데이터 조회 실패 : [%v]", err)
		panic(err.Error())
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
			log.Errorf("(데이터 백업) >>> DB_BACKUP FILE MEMBER WRITE 실패 : [%v]", err)
			return
		}
	}

	//////////////////////////////////////////////
	///매출
	//////////////////////////////////////////////
	_, err = bFile.Write([]byte(prefixRvnStr))
	if err != nil {
		log.Errorf("(데이터 백업) >>> DB_BACKUP FILE PREFIX REVENUES WRITE 실패 : [%v]", err)
		return
	}

	if revenueList, err = DbConf.SelectRevenues(); err != nil {
		log.Errorf("(데이터 백업) >>> DB_BACKUP 매출 데이터 조회 실패 : [%v]", err)
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
			log.Errorf("(데이터 백업) >>> DB_BACKUP FILE REVENUES WRITE 실패 : [%v]", err)
			return
		}
	}

	//////////////////////////////////////////////
	///설정
	//////////////////////////////////////////////
	_, err = bFile.Write([]byte(prefixSettStr))
	if err != nil {
		log.Errorf("(데이터 백업) >>> DB_BACKUP FILE PREFIX SETTING WRITE 실패 : [%v]", err)
		return
	}

	if settingList, err = DbConf.SelectSettings(); err != nil {
		log.Errorf("(데이터 백업) >>> DB_BACKUP 설정 데이터 조회 실패 : [%v]", err)
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
			log.Errorf("(데이터 백업) >>> DB_BACKUP FILE REVENUES WRITE 실패 : [%v]", err)
			return
		}
	}

	bFile.Close()

	//테스트용 sleep
	//time.Sleep(5 * time.Second)
	log.Debug("(데이터 백업) >>> END")

	backupResult <- true
	wg.Done()
	return
}

func clearOldBackupFile(backupPath string, pastDate time.Time) {
	var err error
	var backupFileFullPath string
	var tempPastDate time.Time
	backupFileFullPath = filepath.Join(backupPath, pastDate.Format("2006_01_02"))

	tempPastDate = pastDate
	for i := 0; i < 7; i++ {
		backupFileFullPath = filepath.Join(backupPath, tempPastDate.Format("2006_01_02"))

		//batch file 존재 여부 확인
		if _, err = os.Stat(backupFileFullPath); err == nil {
			//log.Debug("file 존재")
			if err = os.Remove(backupFileFullPath); err != nil {
				log.Errorf("(백업 데이터 정리) >>> 파일 삭제 실패 [%v]", err)
				return
			}
			log.Debugf("(백업 데이터 정리) >>> 파일 삭제 : [%s]", backupFileFullPath)
		} else {
			if errors.Is(err, os.ErrNotExist) {
				log.Debugf("(백업 데이터 정리) >>> 파일 없음 : [%s]", backupFileFullPath)
			} else {
				log.Errorf(err.Error())
				panic(err)
			}
		}

		tempPastDate = tempPastDate.AddDate(0, 0, -1)
	}

	return
}
