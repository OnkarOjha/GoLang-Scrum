package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(times []string) string {

	Hour := 0
	Minute := 0
	Second := 0
	for _, time := range times {
		t1 := strings.Split(time, ":")
		// fmt.Println(t1)

		H, _ := strconv.Atoi(t1[0])
		Hour += H

		M, _ := strconv.Atoi(t1[1])
		Minute += M

		S, _ := strconv.Atoi(t1[2])
		Second += S
	}

	// fmt.Println(Hour) // hour addition values
	// fmt.Println(Minute) // minute addition values
	// fmt.Println(Second) /// second addition values

	// if 17 + 18 = 35
	// 35 -24 = 11 is the real time
	// minHand := 0
	// hourHand := 0

	if Hour > 24 {
		Hour = Hour - 24
	}
	// 55 + 58 = 113
	// 113 - 60 = 53 bachega baki 1 minute plus hoga
	//TODO for above 60 or 24 hours in terms of days
	if Second > 60 {
		Second = Second - 60
		Minute++
	}
	if Minute > 60 {
		Minute = Minute - 60
		Hour++
	}
	if Hour > 24 {
		Hour = Hour - 24
		fmt.Println("Day Changes Here!!")
	}

	// //TODO for above 120 0r 48 hrs in terms of hours
	// if Second > 120{
	//     Second = Second - 120
	//     Minute += 2
	// }
	// if Minute > 120{
	//     Minute = Minute -120
	//     Hour+=2
	// }
	// if Hour > 48{
	//     Hour = Hour - 48
	//     fmt.Println("Do din bad gye bhaisahab!!!")
	// }

	// agar hour  == 24 to 00 and day change
	if Hour == 24 {
		Hour = 0
		fmt.Println("Day badla!!")
	}

	var hourString string
	var minuteString string
	var secondString string

	if Hour < 10 {
		hourString = "0" + strconv.Itoa(Hour)
	}
	if Minute < 10 {
		minuteString = "0" + strconv.Itoa(Minute)
	}
	if Second < 10 {
		secondString = "0" + strconv.Itoa(Second)
	}
	// ye tbhi aa rhi inme valus jb 10 s km ha to
	// fmt.Println(hourString)
	// fmt.Println(minuteString)
	// fmt.Println(secondString)

	//check and print according to 10 se kam ya 10 se jyda values

	var result string

	if Hour < 10 {
		result += hourString + ":"
	} else {
		result += strconv.Itoa(Hour) + ":"
	}

	if Minute < 10 {
		result += minuteString + ":"
	} else {
		result += strconv.Itoa(Minute) + ":"
	}

	if Second < 10 {
		result += secondString
	} else {
		result += strconv.Itoa(Second)
	}

	return result

}
func Comparison(a, b string) bool {
	s1 := strings.Split(a, ":")
	s2 := strings.Split(b, ":")
	//  fmt.Println(s1 , s2)

	hour1, _ := strconv.Atoi(s1[0])
	hour2, _ := strconv.Atoi(s2[0])

	if hour1 == hour2 {
		minute1, _ := strconv.Atoi(s1[1])
		minute2, _ := strconv.Atoi(s2[1])

		if minute1 == minute2 {
			second1, _ := strconv.Atoi(s1[2])
			second2, _ := strconv.Atoi(s2[2])

			return second1 > second2
		}
		return minute1 > minute2
	}

	return hour1 > hour2
}

func Sub(a string, b string) string {
	if !Comparison(a, b) {
		a, b = b, a
	}

	t1 := strings.Split(a, ":")
	t2 := strings.Split(b, ":")

	var Hour1 int
	var Minute1 int
	var Second1 int

	Hour1, _ = strconv.Atoi(t1[0])
	Minute1, _ = strconv.Atoi(t1[1])
	Second1, _ = strconv.Atoi(t1[2])

	// fmt.Println(Hour1)
	// fmt.Println(Minute1)
	// fmt.Println(Second1)

	var Hour2 int
	var Minute2 int
	var Second2 int

	Hour2, _ = strconv.Atoi(t2[0])
	Minute2, _ = strconv.Atoi(t2[1])
	Second2, _ = strconv.Atoi(t2[2])

	// fmt.Println(Hour2)
	// fmt.Println(Minute2)
	// fmt.Println(Second2)

	// real substraction goes with borrow condition
	if Second2 > Second1 {
		Second1 += 60
		Minute1--
	}
	if Minute2 > Minute1 {
		Minute1 += 60
		Hour1--
	}

	hourInt := Hour1 - Hour2
	minuteInt := Minute1 - Minute2
	secondInt := Second1 - Second2

	// // hourInt := Hour1 - Hour2
	// // minuteInt :=

	var result string

	// if Hour <10{
	//     hourString  = "0" +strconv.Itoa(Hour)
	// }
	// if Minute<10{
	//     minuteString = "0" + strconv.Itoa(Minute)
	// }
	// if Second<10{
	//     secondString = "0" + strconv.Itoa(Second)
	// }

	if hourInt < 10 {
		result += "0" + strconv.Itoa(hourInt) + ":"
	} else {
		result += strconv.Itoa(hourInt) + ":"
	}
	if minuteInt < 10 {
		result += "0" + strconv.Itoa(minuteInt) + ":"
	} else {
		result += strconv.Itoa(minuteInt) + ":"
	}
	if secondInt < 10 {
		result += "0" + strconv.Itoa(secondInt)
	} else {
		result += strconv.Itoa(secondInt)
	}

	return result

}

func main() {
	// times1 := []string{"17:00:00" , "22:56:00","00:00:56"}
	// fmt.Println(Add(times1))
	//  fmt.Println(comparison("08:00:05" , "00:05:00"))
	fmt.Println(Sub("00:05:00", "08:00:05"))
	// Sub("17:00:56" , "22:56:00")

	// fmt.Println(t)
}

