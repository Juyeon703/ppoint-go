package utils

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	//fmt.Println(text)
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
	//fmt.Println(rtStr)
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

func MoneyConverter(v int64) string {
	sign := ""

	// Min int64 can't be negated to a usable value, so it has to be special cased.
	if v == math.MinInt64 {
		return "-9,223,372,036,854,775,808"
	}

	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = strconv.FormatInt(v%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v))
	return sign + strings.Join(parts[j:], ",")
}

func InterfaceToInt64(t interface{}) (int64, error) {
	switch t := t.(type) { // This is a type switch.
	case int64:
		return t, nil // All done if we got an int64.
	case int:
		return int64(t), nil // This uses a conversion from int to int64
	case string:
		return strconv.ParseInt(t, 10, 64)
	default:
		return 0, fmt.Errorf("type %T not supported", t)
	}
}

func MoneyReverseConverter(str string) int {
	r := strings.NewReplacer(",", "")
	result, _ := strconv.Atoi(r.Replace(str))
	return result
}
