package gui

import "github.com/lxn/walk"

func MsgBox(title, msg string) {
	walk.MsgBox(winMain, title, msg, walk.MsgBoxOK|walk.MsgBoxIconInformation)
}
