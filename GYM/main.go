package main


import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	models "main/models"
)


var DB *gorm.DB

// ^DUMMY MEMBERS
var memberships []models.MemberInfo = []models.MemberInfo{
	{Id: "1", Name: "Onkar", Membership: "Gold", StartDate: "2023-01-01", EndDate: "2023-03-01", MoneySubmitted: 12000, MonthlyPrice: 2000, Duration: 4, IsDeleted: false},
	{Id: " 2", Name: "Aman", Membership: "Silver", StartDate: "2023-02-01", EndDate: "2023-05-01", MoneySubmitted: 6000, MonthlyPrice: 1000, Duration: 5},
}

func EnrollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data!!!")

	}
	var price models.Price
	var member  models.MemberInfo
	_ = json.NewDecoder(r.Body).Decode(&member)

	//^ membership check
	if member.Membership == "Gold" {
		// monthly price and duration set krna hai
		DB.Where("name = ?", "Gold").First(&price)

		fmt.Println("price: ", price.Price)
		member.Duration = member.MoneySubmitted / price.Price
		member.MonthlyPrice = price.Price
	} else if member.Membership == "Silver" {
		// monthly price and duration set krna hai
		DB.Where("name = ?", "Silver").First(&price)
		fmt.Println("price: ", price.Price)
		member.Duration = member.MoneySubmitted / price.Price
		member.MonthlyPrice = price.Price
	} else {
		http.Error(w, "Membership could be either Gold or Silver!!!", http.StatusMethodNotAllowed)
		return
	}
	// ^ startdate sahi format mei lo
	dateStr := time.Now().Truncate(time.Hour)
	member.StartDate = dateStr.Format("2006-01-02")

	// ^ set end date according to duration and startdate
	startTime, err := time.Parse("2006-01-02", member.StartDate)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return
	}
	//^ add duration to start time
	endTime := startTime.AddDate(0, 0, int(member.Duration))
	member.EndDate = endTime.Format("2006-01-02")

	// ^ membership ke paise agar km de rha hai to

	if member.MoneySubmitted < 1000 {
		http.Error(w, "Money submitted is less than both Gold and silver membership", http.StatusMethodNotAllowed)
		return
	} else if member.MoneySubmitted > 1000 && member.MoneySubmitted < 2000 {
		http.Error(w, "You can either submit 1000 or 2000 or more cause due to our Gold and Silver memerbership rules", http.StatusMethodNotAllowed)
		return

	}

	// ^ membership ke paise agar jyd ade rha hai to balance mei dal do
	if int(member.Duration)*int(member.MonthlyPrice) < int(member.MoneySubmitted) {
		member.RefundedMoney = member.MoneySubmitted - float64(member.Duration)*member.MonthlyPrice
		fmt.Println("refunded: ", member.RefundedMoney)
	}

	// inserting the newly created member into the databaseimport (

	result := DB.Create(&member)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	memberships = append(memberships, member)

	// Handle enrollment logic here
	fmt.Fprintf(w, "Enrollment successful")

}
func ShowMembersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	sort.Slice(memberships, func(i, j int) bool {
		return memberships[i].Id < memberships[j].Id
	})

	// Reading from DB
	result := DB.Unscoped().Find(&memberships)
	if result.Error != nil {
		http.Error(w, "Reading users failed", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(memberships) // this will send response as a json value

	fmt.Fprintf(w, "Reading successful")
}

func EndMembershipHandler(w http.ResponseWriter, r *http.Request) {
	// agar money submitted 12000 hai duration 6 months hai GOLD membership mein
	//  member ka mn badla usne kaha 3 months mei return jan hai
	// to uske 3 months ka khrcha hua 6000 baki bache 6000
	// maine 6000 wapas anhi krne uska 50% RETURN krna hai
	// select membre by ID
	fmt.Println("You are going to end your membership!!")
	fmt.Println("Please enter your Id in params")

	params := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(params)
	fmt.Println("id: ", id)
	now := time.Now().Truncate(24 * time.Hour)

	for i, member := range memberships {

		if member.Id == params {

			//^ money return calculation
			startDate, err := time.Parse("2006-01-02", member.StartDate)
			if err != nil {
				log.Fatal(err)
			}

			member.IsDeleted = true
			member.EndDate = now.Format("2006-01-02")
			duration := now.Sub(startDate)
			days := int(duration.Hours() / 24)
			member.Duration = float64(days / 30)
			MoneyRefund := member.MoneySubmitted - member.Duration*member.MonthlyPrice
			member.RefundedMoney += MoneyRefund / 2
			fmt.Fprintf(w, "Your return balance is: %.2f", member.RefundedMoney)

			memberships = append(memberships, member)
			memberships = append(memberships[:i], memberships[i+1:]...)

			DB.Where("id =?", params).Updates(& models.MemberInfo{RefundedMoney: MoneyRefund / 2, Duration: float64(days / 30), EndDate: now.Format("2006-01-02"), IsDeleted: true})
			DB.Where("id=?", params).Delete(&member)

			json.NewEncoder(w).Encode(member)

			break

		}

	}

}

func MembershipPriceShowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hy Welcome here are our Membership Prices")
	var memPrice []models.Price
	result := DB.Find(&memPrice)
	if result.Error != nil {
		http.Error(w, "Reading users failed", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(memPrice)
}

func MembershipPriceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Now you can update membership prices!!")

	var memPrice  models.Price
	err := json.NewDecoder(r.Body).Decode(&memPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	DB.Model(&models.Price{}).Where("name=?", memPrice.Name).Updates(&memPrice)
	json.NewEncoder(w).Encode(memPrice)

	fmt.Fprint(w, "Membership prices updated")

}


func main() {
	fmt.Println("Implementing GYM membership Task...")

	// implementing DB
	dsn := "host=localhost port=5432 user=postgres password=Onkar17082001 dbname=gym sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error connecting to the database")
	}

	db.AutoMigrate(&models.MemberInfo{}, &models.Price{})
	DB = db

	r := http.NewServeMux()
	r.HandleFunc("/enroll", EnrollHandler)
	r.HandleFunc("/members", ShowMembersHandler)
	r.HandleFunc("/end", EndMembershipHandler)
	r.HandleFunc("/priceshow", MembershipPriceShowHandler)
	r.HandleFunc("/priceupdate", MembershipPriceUpdateHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
