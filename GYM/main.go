package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	Duration       int     `json:"duration"`
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
	{Id: 1, Name: "Onkar", Membership: "Gold", StartDate: "2023-01-01", EndDate: "2023-04-01", MoneySubmitted: 8000, MonthlyPrice: 2000, Duration: 4},
	{Id: 2, Name: "Aman", Membership: "Silver", StartDate: "2023-02-01", EndDate: "2023-06-01", MoneySubmitted: 6000, MonthlyPrice: 1000, Duration: 5},
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
		member.Duration = int(member.MoneySubmitted / membershipPrice.Gold)
		member.MonthlyPrice = membershipPrice.Gold
	}else if member.Membership == "Silver" {
		// monthly price and duration set krna hai
		member.Duration = int(member.MoneySubmitted / membershipPrice.Silver)
		member.MonthlyPrice = membershipPrice.Silver
	}else{
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
	endTime := startTime.AddDate(0, 0, member.Duration)
	member.EndDate = endTime.Format("2006-01-02")

	// ^ membership ke paise agar km de rha hai to

	if member.MoneySubmitted < 1000{
		http.Error(w, "Money submitted is less than both Gold and silver membership", http.StatusMethodNotAllowed)
		return
	}else if member.MoneySubmitted > 1000 && member.MoneySubmitted< 2000{
		http.Error(w, "You can either submit 1000 or 2000 or more cause due to our Gold and Silver memerbership rules", http.StatusMethodNotAllowed)
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

	// Handle reading logic here
	json.NewEncoder(w).Encode(memberships) // this will send response as a json value

	fmt.Fprintf(w, "Reading successful")
}

func EndMembershipHandler(w http.ResponseWriter, r *http.Request) {

}

func MembershipPriceHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("Implementing GYM membership Task...")
	r := http.NewServeMux()
	r.HandleFunc("/enroll", EnrollHandler)
	r.HandleFunc("/members", ShowMembersHandler)
	r.HandleFunc("/end", EndMembershipHandler)
	r.HandleFunc("/price", MembershipPriceHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
