package utils

import (
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
