package test

import (
	"bufio"
	"fmt"
	"os"
	"ppoint/dto"
	"ppoint/query"
	"ppoint/types"
	"ppoint/utils"
	"strconv"
)

var scanner *bufio.Scanner

func CmdTest(dbc *query.DbConfig) {
	var err error

	scanner = bufio.NewScanner(os.Stdin)
	var isExistMember = new(dto.MemberDto)
	var temp string

	for {
		fmt.Println("=============HOME================")
		fmt.Println("1. 고객 조회")
		fmt.Println("2. 모든 고객 관리")
		fmt.Println("3. 매출 조회")
		fmt.Println("4. 설정")
		fmt.Println("5. 종료")
		fmt.Println("=================================")
		fmt.Print("입력 :")
		scanner.Scan()
		temp = scanner.Text()
		if temp == "1" { // 고객 조회
		L1:
			for {
				fmt.Println("=================================")
				fmt.Print("고객 이름 or 핸드폰 번호 전체 or 핸드폰 뒤 4자리 입력해주세요. : ")
				scanner.Scan()
				search := scanner.Text()
				fmt.Println("검색어", search) // 이름 검색 시 중복 폰번호 -> 에러
				fmt.Println("=================================")
				var memberList []types.Member
				if memberList, err = dbc.SelectMemberSearch(search); err != nil {
					panic(err.Error())
				}
				length := len(memberList)
				if length == 0 {
					fmt.Println("==> 검색 결과가 없습니다.")
					fmt.Println("=================================")
					fmt.Println("1. 신규 회원 등록하기")
					fmt.Println("2. 다시 조회하기")
					fmt.Println("=================================")
					fmt.Print("입력 :")
					scanner.Scan()
					temp = scanner.Text()
					if temp == "1" { // 회원 최초 등록하기
						fmt.Println("=================================")
						var intTemp int
						if intTemp, err = AddMember(dbc); err != nil {
							panic(err.Error())
						}
						fmt.Println("=================================")
						fmt.Println("==> 회원 등록 완료")
						if isExistMember, err = dbc.SelectMemberByMemberId(intTemp); err != nil {
							panic(err.Error())
						}
						break
					} else { // 다시 조회하기
						continue
					}
				} else {
					fmt.Println("검색 결과가 존재합니다.")
					fmt.Println("검색 결과 : ", length)
					fmt.Println("{고객번호 고객이름 폰번호 생일 등급번호 보유포인트 가입일 최근방문일")
					for i := 0; i < length; i++ {
						fmt.Println(memberList[i])
					}
					for {
						fmt.Println("=================================")
						fmt.Print("선택할 고객의 고객 번호를 입력해주세요. : ")
						scanner.Scan()
						if utils.RegExpNumber(scanner.Text()) {
							break
						}
					}
					memberIdTemp, _ := strconv.Atoi(scanner.Text())
					if isExistMember, err = dbc.SelectMemberByMemberId(memberIdTemp); err != nil {
						panic(err.Error())
					}
					break
				}
			} // for

			for {
				fmt.Println("=================================")
				fmt.Println("현재 조회 중인 회원 : ", isExistMember.PhoneNumber, isExistMember.MemberName)
				fmt.Println("1. 회원 정보 조회")
				fmt.Println("2. 결제 정보 등록")
				fmt.Println("3. 다른 회원 조회하기")
				fmt.Println("4. 홈으로")
				fmt.Println("=================================")
				fmt.Print("입력 :")
				scanner.Scan()
				temp = scanner.Text()
				if temp == "1" { // 회원 정보 조회
					fmt.Println("=================================")
					fmt.Println("회원 번호 : ", isExistMember.MemberId)
					fmt.Println("이름 : ", isExistMember.MemberName)
					fmt.Println("폰 번호 : ", isExistMember.PhoneNumber)
					fmt.Println("생년월일 : ", isExistMember.Birth)
					fmt.Println("등급 : ", isExistMember.GradeName)
					fmt.Println("보유 포인트 : ", isExistMember.TotalPoint)
					fmt.Println("등록일 : ", isExistMember.CreateDate)
					fmt.Println("수정일 : ", isExistMember.UpdateDate)
					fmt.Println("=================================")
					fmt.Println("1. 회원 정보 수정")
					fmt.Println("2. 회원 정보 삭제")
					fmt.Println("3. 뒤로")
					fmt.Println("=================================")
					fmt.Print("입력 :")
					scanner.Scan()
					temp = scanner.Text()
					if temp == "1" { // 회원 정보 수정
						// 회원 정보 수정 로직
					} else if temp == "2" { // 회원 정보 삭제
						// 결제 내역있는 회원 정보 삭제 안됨 --> 변경 필요
						if err = dbc.DeleteMember(isExistMember.MemberId); err != nil {
							panic(err.Error())
						}
						isExistMember = nil
						break
					} else { // 뒤로
						continue
					}
				} else if temp == "2" { // 결제 정보 등록
					fmt.Println("=================================")
					if err = AddRevenue(isExistMember.MemberId, isExistMember.TotalPoint, dbc); err != nil {
						panic(err.Error())
					}
					fmt.Println("=================================")
					fmt.Println("==> 결제 정보 등록 완료")
					if err = UpgradeMember(isExistMember.MemberId, dbc); err != nil {
						panic(err.Error())
					}
					// 업데이트 내역 다시 가져오기
					if isExistMember, err = dbc.SelectMemberByMemberId(isExistMember.MemberId); err != nil {
						panic(err.Error())
					}
				} else if temp == "3" { // 다른 회원 조회하기
					goto L1
				} else if temp == "4" { // 홈으로
					break
				}
			} // for
		} else if temp == "2" { // 모든 고객 관리
			// 정렬방식 : 높은 등급 순 -> 이름 순
			fmt.Println("=================================")
			var memberList []dto.MemberDto
			if memberList, err = dbc.SelectMembersOrderByGrade(); err != nil {
				panic(err.Error())
			}
			length := len(memberList)
			fmt.Println("모든 고객 수 : ", length)
			fmt.Println("{번호 등급 고객이름 폰번호 생일 보유포인트 가입일 최근방문일")
			for i := 0; i < length; i++ {
				fmt.Println(memberList[i])
			}
		} else if temp == "3" { // 매출 조회
			for {
				fmt.Println("=================================")
				fmt.Println("1. 금일 매출")
				fmt.Println("2. 특정 기간 매출")
				fmt.Println("3. 홈으로")
				fmt.Println("=================================")
				fmt.Print("입력 :")
				scanner.Scan()
				temp = scanner.Text()
				if temp == "1" { // 금일 매출
					var revenueList []types.Revenue
					if revenueList, err = dbc.SelectRevenuesByToday(utils.CurrentDay()); err != nil {
						panic(err.Error())
					}
					length := len(revenueList)
					fmt.Println("검색 결과 : ", length)
					fmt.Println("{번호 고객번호 결제금액 사용포인트 적립포인트 실제결제금액 결제방법 결제일")
					for i := 0; i < length; i++ {
						fmt.Println(revenueList[i])
					}
				} else if temp == "2" { // 특정 기간 매출
					fmt.Println("=================================")
					var startDate string
					var endDate string
					for {
						fmt.Print("시작 날짜(예 : '2010-10-10') : ")
						scanner.Scan()
						startDate = scanner.Text()
						if utils.RegExpDate(startDate) {
							break
						}
					}
					for {
						fmt.Print("종료 날짜(예 : '2010-10-10') : ")
						scanner.Scan()
						endDate := scanner.Text()
						if utils.RegExpDate(endDate) {
							break
						}
					}
					fmt.Println("=================================")
					var revenueList []types.Revenue
					if revenueList, err = dbc.SelectRevenuesByCustomDate(startDate, endDate); err != nil {
						panic(err.Error())
					}
					length := len(revenueList)
					fmt.Println("검색 결과 : ", length)
					fmt.Println("{번호 고객번호 결제금액 사용포인트 적립포인트 실제결제금액 결제방법 결제일")
					for i := 0; i < length; i++ {
						fmt.Println(revenueList[i])
					}
				} else { // 홈으로
					break
				}
			} // for
		} else if temp == "4" { // 설정

		} else { // 종료
			break
		}
	} // for
}

func AddMember(dbc *query.DbConfig) (int, error) {
	var err error
	var member = new(types.Member)
	var memberId int

	fmt.Println("[회원 등록]")
	fmt.Print("이름 : ")
	scanner.Scan()
	member.MemberName = scanner.Text()
	for {
		fmt.Println("폰 번호 (예 : 01012345678) : ")
		scanner.Scan()
		if utils.RegExpCustom("010([0-9]{7,8}$)", scanner.Text()) {
			break
		}
	}
	member.PhoneNumber = scanner.Text()
	for {
		fmt.Print("생년월일(예시 2010-11-11) : ")
		scanner.Scan()
		if utils.RegExpDate(scanner.Text()) {
			break
		}
	}
	member.Birth = scanner.Text()
	if memberId, err = dbc.CreateMember(member); err != nil {
		return 0, err
	}
	return memberId, nil
}

func AddRevenue(memberId, totalPoint int, dbc *query.DbConfig) error {
	var err error
	var revenue = new(types.Revenue)

	revenue.MemberId = memberId
	fmt.Println("[결제 정보 등록]")
	for {
		fmt.Print("결제 금액 : ")
		scanner.Scan()
		if utils.RegExpNumber(scanner.Text()) {
			break
		}
	}
	sales, _ := strconv.Atoi(scanner.Text())
	revenue.Sales = sales
	fmt.Println("결제 방법을 선택해주세요.")
	fmt.Println("1. 현금")
	fmt.Println("2. 카드")
	var payTypeTemp string
	for {
		fmt.Print("입력 : ")
		scanner.Scan()
		scanTemp := scanner.Text()
		if scanTemp == "1" {
			payTypeTemp = "현금"
		} else if scanTemp == "2" {
			payTypeTemp = "카드"
		} else {
			fmt.Println("입력값이 잘못되었습니다.")
			continue
		}
		break
	}
	var subPointTemp int
	for {
		fmt.Println("현재 보유 포인트 : ", totalPoint, "포인트")
		for {
			fmt.Print("사용 포인트 : ")
			scanner.Scan()
			if utils.RegExpNumber(scanner.Text()) {
				break
			}
		}
		subPointTemp, _ = strconv.Atoi(scanner.Text())
		if totalPoint < subPointTemp {
			fmt.Println("현재 보유 포인트 이상 입력할 수 없습니다.")
		} else {
			break
		}
	}
	revenue.SubPoint = subPointTemp
	fixedSalesTemp := sales - subPointTemp
	revenue.FixedSales = fixedSalesTemp
	var addPointTemp int
	revenue.PayType = payTypeTemp
	var setting = new(types.Setting)
	if setting, err = dbc.SelectSettingByPayType(payTypeTemp); err != nil {
		return err
	}
	addPointTemp = fixedSalesTemp * setting.SettingValue / 100
	revenue.AddPoint = addPointTemp
	if err = dbc.UpdateMemberByPoint(memberId, totalPoint-subPointTemp+addPointTemp); err != nil {
		return err
	}
	// 포인트 적립 어떻게 할건지?
	// -> 총 결제 금액인지 실제 사용금액인지?
	// -> 포인트 사용 시 나머지 금액에 대한 적립 할건지?
	// 현재 로직 : 포인트 사용해도 실제 결제 금액으로 포인트 적립 중
	if err = dbc.CreateRevenue(revenue); err != nil {
		return err
	}

	return nil
}

func UpgradeMember(memberId int, dbc *query.DbConfig) error {
	var err error
	var howTotalSales = new(dto.MemberSalesDto)
	var grade = new(types.Grade)

	if howTotalSales, err = dbc.SelectTotalSalesByMember(memberId); err != nil {
		return err
	}
	if grade, err = dbc.SelectGradeByTotalSales(howTotalSales.TotalSales); err != nil {
		return err
	}
	// 등급업 시 고지 어떻게?
	if howTotalSales.GradeId != grade.GradeId {
		if err = dbc.UpdateMemberByGrade(memberId, grade.GradeId); err != nil {
			return err
		}
	}
	return nil
}
