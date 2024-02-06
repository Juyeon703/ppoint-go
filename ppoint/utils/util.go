package utils

import (
	"log"
	"os"
	"regexp"
	"time"
)

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CurrentDay() string {
	return time.Now().Format("2006-01-02")
}

func RegExpCustom(pattern, str string) bool {
	reg, _ := regexp.MatchString(pattern, str)
	return reg
}

func RegExpNumber(str string) bool {
	reg, _ := regexp.MatchString("^[0-9]*$", str)
	return reg
}

func RegExpDate(str string) bool {
	reg, _ := regexp.MatchString("(19[0-9][0-9]|20[0-9][0-9])-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$", str)
	return reg
}

func RegExpPhoneNum(str string) bool {
	reg, _ := regexp.MatchString("^01([0|1|6|7|8|9])-([0-9]{3,4})-([0-9]{4})$", str)
	return reg
}

func PhoneNumAddHyphen(text string) string {
	log.Println(text)
	rtStr := ""
	for idx, str := range text {
		rtStr += string(str)
		if len(text) == 11 {
			if idx == 2 || idx == 6 {
				rtStr += "-"
			}
		} else if len(text) == 10 {
			if idx == 2 || idx == 5 {
				rtStr += "-"
			}
		}
	}
	log.Println(rtStr)
	return rtStr
}

func IsExistFile(fname string) bool {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateFilePath(fname string) error {
	if err := os.MkdirAll(fname, os.ModePerm); err != nil {
		return err
	}

	return nil
}
