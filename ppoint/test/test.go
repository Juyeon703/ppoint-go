package test

import (
	"fmt"
	"ppoint/query"
	"ppoint/types"
)

func Test(DbConf *query.DbConfig) {
	var err error
	var memberList []types.Member

	var memberId int
	var member = new(types.Member)
	member.MemberName = "bbbb"
	member.PhoneNumber = "01022223333"
	member.Birth = "2017-11-11"
	if memberId, err = DbConf.CreateMember(member); err != nil {
		panic("createUser err")
	}
	fmt.Println("CREATE_MEMBER", memberId)

	if err = DbConf.UpdateMemberByPhoneNumber(2, "01033333333", "cccccc"); err != nil {
		panic(err.Error())
	}
	fmt.Println("UPDATE_MEMBER")

	if memberList, err = DbConf.SelectMembers(); err != nil {
		panic(err.Error())
	}
	fmt.Println("SELECT_MEMBER")
	fmt.Printf("%v", memberList)

	if err = DbConf.DeleteMember(2); err != nil {
		panic(err.Error())
	}
	fmt.Println("DELETE_MEMBER")

	if err = DbConf.CreateUser("새우", "빨감", "01035353535", "shrimp", "1111", "shrimp@gmail.com"); err != nil {
		panic(err.Error())
	}
	fmt.Println("CREATE_USER")

	if err = DbConf.UpdateUserByPhoneNumber(1, "01033333333", "sleep"); err != nil {
		panic(err.Error())
	}
	fmt.Println("UPDATE_USER")

	var userList []types.Users
	if userList, err = DbConf.SelectUsers(); err != nil {
		panic(err.Error())
	}
	fmt.Println("SELECT_USERS")
	fmt.Printf("%v", userList)

	if err = DbConf.DeleteUser(1); err != nil {
		panic(err.Error())
	}
	fmt.Println("DELETE_USER")

	var revenue = new(types.Revenue)
	revenue.MemberId = 1
	revenue.Sales = 10000
	revenue.SubPoint = 0
	revenue.AddPoint = 100
	revenue.FixedSales = 10000
	revenue.PayType = "현금"
	if err = DbConf.CreateRevenue(revenue); err != nil {
		panic(err.Error())
	}
	fmt.Println("CREATE_REVENUE")

	var revenueList []types.Revenue
	if revenueList, err = DbConf.SelectRevenues(); err != nil {
		panic(err.Error())
	}
	fmt.Println("SELECT_REVENUE")
	fmt.Printf("%v", revenueList)

	if err = DbConf.DeleteRevenue(4); err != nil {
		panic(err.Error())
	}
	fmt.Println("DELETE_USER")

	if err = DbConf.CreateGrade("임시"); err != nil {
		panic(err.Error())
	}
	fmt.Println("CREATE_GRADE")

	var gradeList []types.Grade
	if gradeList, err = DbConf.SelectGrades(); err != nil {
		panic(err.Error())
	}
	fmt.Println("SELECT_GRADE")
	fmt.Printf("%v", gradeList)

	if err = DbConf.DeleteGrade(5); err != nil {
		panic(err.Error())
	}
	fmt.Println("DELETE_GRADE")

	if err = DbConf.CreateSetting("등급업", "30000"); err != nil {
		panic(err.Error())
	}
	fmt.Println("CREATE_SETTING")

	var settingList []types.Setting
	if settingList, err = DbConf.SelectSettings(); err != nil {
		panic(err.Error())
	}
	fmt.Println("SELECT_SETTING")
	fmt.Printf("%v", settingList)

	if err = DbConf.DeleteSetting(1); err != nil {
		panic(err.Error())
	}
	fmt.Println("DELETE_SETTING")
}
