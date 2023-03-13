package models

import (
	"gorm.io/gorm"
)

type MemberInfo struct {
	Id             string  `json:"id" gorm:"default:uuid_generate_v4()"`
	Name           string  `json:"name"`
	Membership     string  `json:"membership"`
	StartDate      string  `json:"startDate"`
	EndDate        string  `json:"endDate"`
	MoneySubmitted float64 `json:"moneySubmitted"`
	MonthlyPrice   float64 `json:"monthlyPrice"`
	Duration       float64 `json:"duration"`
	IsDeleted      bool    `json:"isDeleted"`
	RefundedMoney  float64 `json:"refundedMoney"`
	DeletedAt      gorm.DeletedAt
}

type Price struct {
	Name  string
	Price float64
}
