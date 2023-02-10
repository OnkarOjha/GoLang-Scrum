package main
import "fmt"

func main(){
	// var floatValue float64 = 5.45
	// var intValue int = int(floatValue)

	// fmt.Printf("Float value is %g\n" , floatValue)
	// fmt.Printf("Integer value is %d" , intValue)

	// var intValue int  = 2
	// var floatValue float32 = float32(intValue)

	// fmt.Printf("Int value is %d\n" , intValue)
	// fmt.Printf("Float value is %f\n" , floatValue)

	var number1 int  = 10
	var number2 float32 = 5.6

	var sum float32

	sum = float32(number1) + number2

	fmt.Printf("Sum is %g", sum)

	

}