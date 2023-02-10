package main

import (
	"fmt"
	"time"
)

func main(){
	// start := time.Now()
	// t := time.Now()
	// fmt.Println(t.Format("02-01-2006 03:04:06 Monday"))
	// fmt.Println()
	// elapsed := t.Sub(start)
	// elap := time.Since(start)
	// fmt.Println()
	// fmt.Println(elap)

	// fmt.Println(elapsed)
	// fmt.Println()
	// tl , _ := time.LoadLocation("Local")
	// fmt.Println(t.In(tl))

	// timeNow := time.Now()
	// supposedTime := time.Now()

	// // fmt.Println(timeNow.Equal(supposedTime))
	// fmt.Println(timeNow == supposedTime)
	// _, month, day := time.Now().Date()
	// if month == time.February && day == 8 {
	// 	fmt.Println("Happy Go day!")
	// }else{
	// 	fmt.Println("Not a golang day")
	// }

	// start := time.Date(2023, 10, 31, 2, 0, 0, 0, time.Local)
	// // start := time.Now()
	// afterTenSeconds := start.Add(time.Second * 10)
	// afterTenMinutes := start.Add(time.Minute * 10)
	// afterTenHours := start.Add(time.Hour * 10)
	// afterTenDays := start.Add(time.Hour * 24 * 10)

	// onedaylater := start.AddDate(0,0,1)
	// occt31 := start.AddDate(0,1,0)

	// fmt.Printf("start = %v\n", start.Format("02-01-2006 03:04:06 Monday"))
	// fmt.Printf("start.Add(time.Second * 10) = %v\n", afterTenSeconds.Format("02-01-2006 03:04:06 Monday"))
	// fmt.Printf("start.Add(time.Minute * 10) = %v\n", afterTenMinutes.Format("02-01-2006 03:04:06 Monday"))
	// fmt.Printf("start.Add(time.Hour * 10) = %v\n", afterTenHours.Format("02-01-2006 03:04:06 Monday"))
	// fmt.Printf("start.Add(time.Hour * 24 * 10) = %v\n", afterTenDays.Format("02-01-2006 03:04:06 Monday"))
	// fmt.Printf("ek din add krne k baad  = %v\n", onedaylater.Format("02-01-2006 03:04:06 Monday"))

	// fmt.Printf("OCTOBER 31 mein 1 month add kra to = %v\n", occt31.Format("02-01-2006 03:04:06 Monday"))

		// Parse a time value from a string in the standard Unix format.
		t, err := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
		if err != nil { // Always check errors even if they should not happen.
			panic(err)
		}
	
		tz, err := time.LoadLocation("Asia/Shanghai")
		if err != nil { // Always check errors even if they should not happen.
			panic(err)
		}

		fmt.Println(tz)
	
		// time.Time's Stringer method is useful without any format.
		fmt.Println("default format:", t)
	
		// Predefined constants in the package implement common layouts.
		fmt.Println("Unix format:", t.Format(time.UnixDate))

}