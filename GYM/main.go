package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type MemberInfo struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Membership     string  `json:"membership"`
	StartDate      string  `json:"startDate"`
	EndDate        string  `json:"endDate"`
	MoneySubmitted float64 `json:"moneySubmitted"`
	MonthlyPrice   float64 `json:"monthlyPrice"`
	Duration       float64 `json:"duration"`
	IsDeleted      bool    `json:"isDeleted"`
	RefundedMoney  float64 `json:"refundedMoney"`
}

type Price struct {
	Gold   float64
	Silver float64
}

var membershipPrice = &Price{
	Gold:   2000,
	Silver: 1000,
}

// ^DUMMY MEMBERS
var memberships []MemberInfo = []MemberInfo{
	{Id: 1, Name: "Onkar", Membership: "Gold", StartDate: "2023-01-01", EndDate: "2023-03-01", MoneySubmitted: 12000, MonthlyPrice: 2000, Duration: 4, IsDeleted: false},
	{Id: 2, Name: "Aman", Membership: "Silver", StartDate: "2023-02-01", EndDate: "2023-05-01", MoneySubmitted: 6000, MonthlyPrice: 1000, Duration: 5},
}

func EnrollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data!!!")

	}

	var member MemberInfo
	_ = json.NewDecoder(r.Body).Decode(&member)
	// fmt.Println("data: ", member)
	member.Id = len(memberships) + 1
	//^ membership check
	if member.Membership == "Gold" {
		// monthly price and duration set krna hai
		member.Duration = member.MoneySubmitted / membershipPrice.Gold
		member.MonthlyPrice = membershipPrice.Gold
	} else if member.Membership == "Silver" {
		// monthly price and duration set krna hai
		member.Duration = member.MoneySubmitted / membershipPrice.Silver
		member.MonthlyPrice = membershipPrice.Silver
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

	// Handle reading logic here
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
	now := time.Now().Truncate(24 * time.Hour)
	for i, member := range memberships {
		id, _ := strconv.Atoi(params)
		if member.Id == id {
			// memberships = append(memberships[:index], memberships[index+1:]...)
  
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

			json.NewEncoder(w).Encode(member)

			break

		}
		// w.Write([]byte("You have successfully ended your membership"))
	}

}

func MembershipPriceShowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hy Welcome here are our Membership Prices")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(membershipPrice) // this
}

func MembershipPriceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Now you can update membership prices!!")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var data struct {
		Gold   float64 `json:"gold"`
		Silver float64 `json:"silver"`
	}

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// ^ sirf ek value change krni hai or dusri nhi
	if data.Gold == 0.0 {
		data.Gold = membershipPrice.Gold
	} else if data.Silver == 0.0 {
		data.Silver = membershipPrice.Silver
	}
	membershipPrice.Gold = data.Gold
	membershipPrice.Silver = data.Silver

	fmt.Fprint(w, "Membership prices updated")

}

func main() {
	fmt.Println("Implementing GYM membership Task...")
	r := http.NewServeMux()
	r.HandleFunc("/enroll", EnrollHandler)
	r.HandleFunc("/members", ShowMembersHandler)
	r.HandleFunc("/end", EndMembershipHandler)
	r.HandleFunc("/priceshow", MembershipPriceShowHandler)
	r.HandleFunc("/priceupdate", MembershipPriceUpdateHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
