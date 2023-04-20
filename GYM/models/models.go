package models

import (
	"time"

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

// users struct
type Users struct {
	User_id        string `json:"user_id" gorm:"default:uuid_generate_v4();uniqueIndex"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Gender         string `json:"gender" `
	Contact_Number string `json:"contact_number"`
}

// subscriptions struct
type Subscription struct {
	Payment_id string `json:"payment_id"`
	// Payment           Payment     `gorm:"references:Payment_id"` //FK
	User_id         string  `json:"user_id"`
	User            Users   `gorm:"references:User_id"` //FK
	Start_date      string  `json:"start_date"`
	Membership_type string  `json:"membership_type"`
	End_date        string  `json:"end_date"`
	Duration        float64 `json:"duration"`
	Employee_id     string  `json:"employee_id"`
	// Employee          GymEmployee `gorm:"references:Employee_id"` //FK
	Trainer_name  string  `json:"trainer_name"`
	IsDeleted     bool    `json:"isDeleted"`
	RefundedMoney float64 `json:"refundedMoney"`
}

// Payment struct
type Payment struct {
	Payment_id   string    `json:"payment_id" gorm:"default:uuid_generate_v4();uniqueIndex"`
	User_id      string    `json:"user_id"`
	User         Users     `gorm:"references:User_id"` //FK
	Amount       float64   `json:"amount"`
	Date         time.Time `json:"date"`
	Payment_type string    `json:"payment_type"`
}

// Equipment struct
type Equipment struct {
	Model_number string `json:"model_number" gorm:"default:uuid_generate_v4();unique"`
	Equip_name   string `json:"equip_name"`
	Quantity     int64  `json:"quantity"`
}

type GymEmployee struct {
	Employee_id    string `json:"employee_id" gorm:"default:uuid_generate_v4();unique"`
	First_name     string `json:"first_name"`
	Last_name      string `json:"last_name"`
	Gender         string `json:"gender"`
	Contact_Number string `json:"contact_number"`
	Role           string `json:"role"`
}

// TODO user attendance
// TODO Employee attendance
// TODO offers
