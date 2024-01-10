package types

import "database/sql"

type DbConfig struct {
	DbConnection *sql.DB
}

// 임시
type Total struct {
	TotalSales int
	GradeId    int
	GradeName  string
}

type MemberDto struct {
	MemberId    int
	GradeName   string
	MemberName  string
	PhoneNumber string
	Birth       string
	TotalPoint  int
	CreateDate  string
	UpdateDate  string
}

type Users struct {
	UserId      int
	StoreName   string
	UserName    string
	PhoneNumber string
	Id          string
	Password    string
	Email       string
	CreateDate  string
	UpdateDate  string
}

type Member struct {
	MemberId    int
	MemberName  string
	PhoneNumber string
	Birth       string
	GradeId     int
	TotalPoint  int
	CreateDate  string
	UpdateDate  string
}

type Revenue struct {
	RevenueId  int
	MemberId   int
	Sales      int
	SubPoint   int
	AddPoint   int
	FixedSales int
	PayType    string
	CreateDate string
}

type Grade struct {
	GradeId    int
	GradeName  string
	GradeValue int
}

type Setting struct {
	SettingId          int
	SettingType        string
	SettingName        string
	SettingValue       int
	SettingDescription string
}
