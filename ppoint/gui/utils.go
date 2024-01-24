package gui

import (
	"fmt"
	"github.com/lxn/walk"
)

func MsgBox(title, msg string) {
	walk.MsgBox(winMain, title, msg, walk.MsgBoxOK|walk.MsgBoxIconInformation)
}

func PhoneNumAddHyphen(text string) string {
	fmt.Println(text)
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
	fmt.Println(rtStr)
	return rtStr
}
