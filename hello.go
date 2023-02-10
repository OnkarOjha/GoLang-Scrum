package main

// // import (
// // 	"fmt"
// // 	// "math"

// // )

// // type ks struct{

// // }

// // // func sum(s []int, c chan int){
// // // 	sum := 0

// // // 	for _ , v:=range s{
// // // 		sum+=v
// // // 	}

// // // 	c<-sum
// // // }

// // func main(){
// // 	// fmt.Println("This is for channel function")

// // 	// s := []int{11,2,4,5,6,7,9}
// // 	// c := make(chan int)
// // 	// go sum(s[:len(s)/2] ,c)
// // 	//  go sum(s[len(s)/2:] , c)

// // 	// x , y := <-c , <-c

// // 	// fmt.Println(x, y , x+y)

// // 	// fmt.Println("This is for Type casting ")

// // 	// greeting := "H there!!"
// // 	// fmt.Println([]byte(greeting))

// // 	//TODO :- TYPE CONVERSION

// // 	// fmt.Println("Type Conversion it is")
// // 	// var x , y int = 3 , 4
// // 	// var f float64 = math.Sqrt(float64(x*x + y*y))
// // 	// var z uint = uint(f)
// // 	// fmt.Println(x, y , z)
// // 	fmt.Println("dfsf")

// // }

// // package main

// import (
// 	"fmt"
// 	// "sort"
// )
// // func sayHello(names ...string) {
// //     for _, n := range names {
// //         fmt.Printf("Hello %s\n", n)
// //     }
// // }
// // func main() {

// // 	// sayHello()
// //     // sayHello("Rahul")
// //     // sayHello("Mohit", "Rahul", "Rohit", "Johny")

// // 	// seasons := []string{"Spring", "Summer", "Autumn", "Winter" , "sdfjgsdhfg"}
// // 	// printSlice(seasons)
// // 	// days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday","djbfgff" , "dfbsf" ,"djbfgff" , "dfbsf"}
// // 	// // printSlice(days)
// // 	// fmt.Println(days)

// // 	// fmt.Println("\nInterface")
// // 	// interfaceEx := []interface{}{"Australia", "Canada", "Japan"}
// // 	// fmt.Println(interfaceEx...)

// // 	// const placeOfInterest = `⌘`

// //     // fmt.Printf("plain string: ")
// //     // fmt.Printf("%s", placeOfInterest)
// //     // fmt.Printf("\n")

// //     // fmt.Printf("quoted string: ")
// //     // fmt.Printf("%+q", placeOfInterest)
// //     // fmt.Printf("\n")

// //     // fmt.Printf("hex bytes: ")
// //     // for i := 0; i < len(placeOfInterest); i++ {
// //     //     fmt.Printf("%x ", placeOfInterest[i])
// //     // }
// //     // fmt.Printf("\n")
// // 	// a:="😂"
// // 	// fmt.Println(a)

// // 	// s := []int{9,8,7,6,5}
// // 	// sort.Slice(s, func(i , j int) bool{
// // 	// 	return s[i] < s[j]
// // 	// })

// // 	// for _ , v := range s{
// // 	// 	fmt.Println(v)
// // 	// }

// // 	// people :=
// //  i:= "2"
// // 	fmt.Sprintf("%v", i)
// // }

// // func printSlice(s []string) {
// // 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// //

// // package main

// // import (
// // 	"fmt"
// // )

// // func main() {
// //   // declaring variables of different datatypes
// //   var message string = "Hello and welcome to"
// //   var year int = 2021

// // 	// printing out the declared variables following the format string
// // 	var complete_msg = fmt.Sprintf("%s educative in %d", message, year)
// //   fmt.Print(complete_msg)
// // }

// package main

// import (
// 	"fmt"
// 	// "strings"
// //   "math/rand"
// //  "time"
// )

// //FIXME EFWE
// var a string

// func init(){
//   fmt.Println("nefore deer")
//   fmt.Println("nefore deer1")
//   defer fmt.Println("init funciton")
//   a = "5"
// }

// // func sqrt(a int)int{
// //   i := 1
// //   res :=1
// //   for res<=a{
// //     i++
// //     res = i*i
// //   }

// //   return i-1
// // }

// func main(){
//   // var a int
//   // fmt.Println("Enter a number: ")
//   // fmt.Scanln(&a)
//   // fmt.Println("The squared number is: ",sqrt(a))
//   // isSkilled := true
//   // for isSkilled {
//   // fmt.Println("Ready for mission!" + "ssf")
//   // isSkilled = false
//   // }

//   // 	qs:= "onkar ojha"

// 	// fmt.Printf("%c",qs[0])

//  fmt.Println(a)
//   // var sb strings.Builder
//   // fmt.Println("This is a string builder",sb.String())

//   // sb.WriteString("onkar")
//   // fmt.Println(sb.Cap() , sb.Len())
//   //   sb.Grow(3)
//   //   sb.Grow(4)
//   //   sb.Grow(6)
//   //   fmt.Println(sb.Cap() , sb.Len())

//   // rand.Seed(time.Now().UnixNano())
// 	// fmt.Println(rand.Intn(100))
//   // fmt.Println(rand.Intn(100))

//   // fmt.Println(rand.Intn(100))
//   // fmt.Println(rand.Intn(100))
//   // fmt.Println(rand.Intn(100))

//   // i := 1
//   // withoutPointer(i)
//   // fmt.Println(i)

//   // fmt.Println("main funciton")
//   // defer fmt.Println("defer main funciton")
//   // defer fmt.Println(1)
//   // defer fmt.Println(2)

// //  a := [5]int{1,2,3,4,5}
// //  fmt.Println("a1",a[:])
// //  fmt.Println("a2",a[0:])
// //  fmt.Println("a1",a[:5])
// //  fmt.Println("a1",a[0:5])
// //  fmt.Println("a1",a[:])

// // fmt.Println(-13>>1)

//     // var a int32
//     // v1 := a + 'v'
//     // // v2 := a + "v"
//     // fmt.Println(v1)
//     // // fmt.Println(v2)

// }

// // func withoutPointer(i int){
// //   i = i+1
// // }

// var WhatIsThe = AnswerToLife()

// func AnswerToLife() int {
//     return 42
// }

//  func init() {
//     WhatIsThe = 0
// }

// func main() {
//     if WhatIsThe == 0 {
//         fmt.Println("It's all a lie.")
//     }
// }

// import (
//   "testing"
//    "github.com/stretchr/testify/assert"
// )

// func TestSomething(t *testing.T) {

//   var a string = "Hello"
//   var b string = "Hello"

//   assert.Equal(t, a, b, "The two words should be the same.")

// }

//  import (

//   "fmt"

//  )
// func main(){
//   fmt.Println("cards game!!")

//   card := deck{"spades" , "hearts"}
//   card.print()

// }

import (
	"fmt"
	"strconv"
	"strings"
	// "math/rand"
	// "time"
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
func comparison(a, b string) bool {
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

// [00 05 00] [08 00 05]

func Sub(a string, b string) string {
	if !comparison(a, b) {
		a, b = b, a
	}
	t1 := strings.Split(a, ":")
	t2 := strings.Split(b, ":")
	fmt.Println(t1, t2)

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

	var result string

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
	fmt.Println(Sub("08:00:05", "00:05:00"))
	// Sub("17:00:56" , "22:56:00")

	// fmt.Println(t)
}
