package dto

import "time"

type MemberAddDto struct {
	MemberName  string
	PhoneNumber string
	Birth       time.Time
}
