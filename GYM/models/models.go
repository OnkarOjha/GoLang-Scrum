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


/*

func ShowUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var show []models.Payment
	DB.Joins("Payments").Joins("User").Find(&show)
	json.NewEncoder(w).Encode(&show)
}

func UserShowHandler(w http.ResponseWriter, r *http.Request) {
	// this will show user info full and user susbcriptons information

	// this will show user info full and user susbcriptons information
	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)
	w.Header().Set("Content-Type", "application/json")

	var user models.Subscription
	DB.Joins("User").Joins("Payment").Omit("User").Where("subscriptions.user_id = ?", id).First(&user)

	json.NewEncoder(w).Encode(&user)
	// id := r.URL.Query().Get("id")
	// fmt.Println("id: ", id)
	// w.Header().Set("Content-Type", "application/json")
	// var user models.Users
	// var subscriptions []models.Subscription
	// DB.Where("user_id =?", id).First(&user)

	// // retrieve the payment details for the user
	// var payment models.Payment
	// DB.Omit("User").Where("user_id = ?", id).First(&payment)

	// // retrieve the full subscription details for the user
	// DB.Omit("User","Payments").Where("user_id = ?", id).Find(&subscriptions)

	// // create a map to hold the user, payment, and subscription details
	// userMap := make(map[string]interface{})
	// userMap["user"] = user
	// userMap["payment"] = payment
	// userMap["subscriptions"] = subscriptions

	// // encode the map as JSON and send it in the response
	// json.NewEncoder(w).Encode(userMap)

}

r := http.NewServeMux()
	r.HandleFunc("/user", UserEnrollment)
	r.HandleFunc("/user/subscription", SubscriptionHandler)
	r.HandleFunc("/user/endSubs", EndSubscriptionHandler)
	r.HandleFunc("/employee", EmployeeHandler)
	r.HandleFunc("/user/payment", PaymentHandler)
	r.HandleFunc("/equipment", EquipementHandler)
	r.HandleFunc("/users/show", ShowUsersHandler)
	r.HandleFunc("/users/showByID", UserShowHandler)

	r.HandleFunc("/enroll", EnrollHandler)
	r.HandleFunc("/members", ShowMembersHandler)
	r.HandleFunc("/end", EndMembershipHandler)
	r.HandleFunc("/priceshow", MembershipPriceShowHandler)
	r.HandleFunc("/priceupdate", MembershipPriceUpdateHandler)

	*/

	