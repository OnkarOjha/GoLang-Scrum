// // // package main

// // // import (
// // // 	"fmt"
// // // 	"time"
// // // )

// // // type MyError struct {
// // // 	When time.Time
// // // 	What string
// // // }

// // // func (e *MyError) Error() string {
// // // 	return fmt.Sprintf("at %v, %s",
// // // 		e.When, e.What)
// // // }

// // // func run() error {
// // // 	return &MyError{
// // // 		time.Now(),
// // // 		"it didn't work",
// // // 	}
// // // }

// // // func main() {
// // // 	if err := run(); err != nil {
// // // 		// fmt.Println(err)
// // // 		panic(err)
// // // 	}
// // // }

// // package main

// // import (
// // 	//  "errors"
// // 	"fmt"
// // 	 "log"
// // )

// // type MyError struct{
// // 	extrainfo string
// // }

// // func (e *MyError) Error() string{
// // 	return fmt.Sprintf("got an error : %s", e.extrainfo)
// // }

// // func main(){
// // 	fmt.Println("Playing with error interface!!!!")

// // 	// var e error
// // 	// e = &MyError{extrainfo: "boom"}

// // 	// fmt.Println(e)

// // 	// // TODO second way

// // 	// f := errors.New( "Bhai ye khrnak hai")

// // 	// fmt.Println(f)

// // 	// //TODO 3rd way
// // 	// //checking an error type for an error interface

// // 	// var c *MyError // taking pointer to error interface
// // 	// check1 := errors.As(e , &c) // erros.As()
// // 	// fmt.Println(check1)

// // 	// //TODO 4th way
// // 	// //checking that error is wrapped  FOR ERROR CHAIN
// // 	// eWrapped := fmt.Errorf("Error: %w ", e) //Errorf formats according to a format specifier and returns the string as a value that satisfies error.
// // 	// // check2:= errors.Is(eWrapped , e)
// // 	// check2 := errors.Is(eWrapped , &MyError{extrainfo: "Bhaisahab"})
// // 	// fmt.Println(check2)

// // 	//TODO flow control

// // 	i , err := returnInt()
// // 	if err!=nil{
// // 		log.Fatal(err)
// // 	}
// // 	log.Print(i)
// // 	// fmt.Println(i)
// // 	// fmt.Println(log.Lmicroseconds)
// // 	// fmt.Println(log.Lmsgprefix)
// // 	// fmt.Println(log.LUTC)

// // 	s , err2 := returnString(i)
// // 	if err2!=nil{
// // 		// panic(err2)
// // 		log.Fatal(err2)
// // 	}
// // 	log.Print(s)

// // }

// // func returnInt() (int , error){
// // 	return 40, nil
// // }

// // func returnString(x int) (string , error){
// // 	if x>42{
// // 		return "That's fine" , nil
// // 	}
// // 		return "" ,&MyError{extrainfo: "haa haai 2222"}

// // }

// //TODO error exercise

// package main

// import (
// 	// "errors"
// 	"fmt"
// 	"math"
// )

// type ErrNegativeSqrt float64

// func (e ErrNegativeSqrt) Error() string{
// 	//srf print kro

// 	return fmt.Sprintf("cannot Sqrt negative number: %v \n",e)
// }


// func Sqrt(x float64) (float64, error) {
// 	//checking waala kaam
// 	if x<0{
// 		var err ErrNegativeSqrt = ErrNegativeSqrt(x)
// 		return 0 , err
// 	}else{
// 		return math.Sqrt(x) , nil
// 	}


	
	
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// 	fmt.Println(Sqrt(-2))
// }

// // package main

// // import (
// // 	"fmt"
// // 	"time"
// // )

// // type MyError struct {
// // 	When time.Time
// // 	What string
// // }

// // func (e *MyError) Error() string {
// // 	return fmt.Sprintf("at %v, %s",
// // 		e.When, e.What)
// // }

// // func run() error {
// // 	return &MyError{
// // 		time.Now(),
// // 		"it didn't work",
// // 	}
// // }

// // func main() {
// // 	if err := run(); err != nil {
// // 		// fmt.Println(err)
// // 		panic(err)
// // 	}
// // }

// package main

// import (
// 	//  "errors"
// 	"fmt"
// 	 "log"
// )

// type MyError struct{
// 	extrainfo string
// }

// func (e *MyError) Error() string{
// 	return fmt.Sprintf("got an error : %s", e.extrainfo)
// }

// func main(){
// 	fmt.Println("Playing with error interface!!!!")

// 	// var e error
// 	// e = &MyError{extrainfo: "boom"}

// 	// fmt.Println(e)

// 	// // TODO second way

// 	// f := errors.New( "Bhai ye khrnak hai")

// 	// fmt.Println(f)

// 	// //TODO 3rd way
// 	// //checking an error type for an error interface

// 	// var c *MyError // taking pointer to error interface
// 	// check1 := errors.As(e , &c) // erros.As()
// 	// fmt.Println(check1)

// 	// //TODO 4th way
// 	// //checking that error is wrapped  FOR ERROR CHAIN
// 	// eWrapped := fmt.Errorf("Error: %w ", e) //Errorf formats according to a format specifier and returns the string as a value that satisfies error.
// 	// // check2:= errors.Is(eWrapped , e)
// 	// check2 := errors.Is(eWrapped , &MyError{extrainfo: "Bhaisahab"})
// 	// fmt.Println(check2)

// 	//TODO flow control

// 	i , err := returnInt()
// 	if err!=nil{
// 		log.Fatal(err)
// 	}
// 	log.Print(i)
// 	// fmt.Println(i)
// 	// fmt.Println(log.Lmicroseconds)
// 	// fmt.Println(log.Lmsgprefix)
// 	// fmt.Println(log.LUTC)

// 	s , err2 := returnString(i)
// 	if err2!=nil{
// 		// panic(err2)
// 		log.Fatal(err2)
// 	}
// 	log.Print(s)

// }

// func returnInt() (int , error){
// 	return 40, nil
// }

// func returnString(x int) (string , error){
// 	if x>42{
// 		return "That's fine" , nil
// 	}
// 		return "" ,&MyError{extrainfo: "haa haai 2222"}

// }

//TODO error exercise

package main

import (
	// "errors"
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	//srf print kro

	return fmt.Sprintf("cannot Sqrt negative number: %v \n",int(e))
}


func Sqrt(x float64) (float64, error) {
	//checking waala kaam
	if x<0{
		var err ErrNegativeSqrt = ErrNegativeSqrt(x)
		return 0 , err
	}else{
		return math.Sqrt(x) , nil
	}


	
	
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}





