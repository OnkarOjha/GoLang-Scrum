// package main
// import "fmt"
// func main(){
// // 	var salary1 float32
// // 	var salary2 float64

// //   salary1 = 50000.503882901

// //   // can store decimals with greater precision
// //   salary2 = 50000.503882901

// //   fmt.Println(salary1) 
// //   fmt.Print(salary2)



// // currentSalary := 50000

// // fmt.Println("Hello")
// // fmt.Println("World!")
// // fmt.Println("Current Salary:", currentSalary)


// // taking input form users

// // var name string
// // var age int

// // fmt.Print("Print your name")
// // fmt.Scan(&name)
// // fmt.Println("Print your age")
// // fmt.Scan(&age)

// // fmt.Println("Your name is:", name)

// // fmt.Println("Your age is:", age)

// var temp float32
// var sunny bool

// //taking input
// fmt.Println("Enter the current temprature:")
// fmt.Scanf("%g" , &temp)

// fmt.Println("Is the day sunny?")
// fmt.Scanf("%t",  &sunny)

// fmt.Printf("Current temprature: %g\nIs the day Sunny? %t", temp , sunny)

// }

package main

import "fmt"

func main() {
    const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    fmt.Println("Println:")
    fmt.Println(sample)

    fmt.Println("Byte loop:")
    for i := 0; i < len(sample); i++ {
        fmt.Printf("%x ", sample[i])
    }
    fmt.Printf("\n")

    fmt.Println("Printf with %x:")
    fmt.Printf("%x\n", sample)

    fmt.Println("Printf with % x:")
    fmt.Printf("% x\n", sample)

    fmt.Println("Printf with %q:")
    fmt.Printf("%q\n", sample)

    fmt.Println("Printf with %+q:")
    fmt.Printf("%+q\n", sample)
}