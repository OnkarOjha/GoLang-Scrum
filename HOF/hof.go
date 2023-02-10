package main
import (
	
	"fmt"
	// "math"
)

// func sphereVolume(vol func(radius float64)float64){
// 	fmt.Println("The volume of sphere is", vol(3))
// }
// func main(){
// 	resultSphere := func(radius float64) float64{
// 		result := 4/3 * math.Pi *radius *radius
// 		return result
// 	}

// 	sphereVolume(resultSphere)
// }
// TODO returning functions from functions
// func sphereVolume() func(radius float64) float64{
// 	result := func(radius float64) float64{
// 		r := 4/3 * math.Pi * radius * radius *radius
// 		return r
// 	}

// 	return result
// }
// func main(){
// 	volume := sphereVolume()
// 	fmt.Println("The volume of sphere: ",volume(3))
// }


// func sum(x , y int)int{
// 	return x+y
// }

// func partialSum(x int ) func(int) int{
// 	return func(y int)int{
// 		return sum(x,y)
// 	}
// }

// func main(){
// 	partial := partialSum(7)
// 	fmt.Println(partial(3))
// }

// func sum(x int) func(int) func(int) int{

// 	return func(y int) func(int) int{
// 		return func(z int )int{
// 			return x*x + y*y+z*z
// 		}
// 	}

// }

// func main(){
// 	fmt.Println(sum(6)(7)(8))
// }
// func print(x , y int , add func(int,int) int , substract func(int, int)int){
// 	fmt.Printf("Sum value of %d and %d is %d",x,y,add(x,y))
// 	fmt.Printf("Substraction value of %d and %d is %d",x,y,substract(x,y))
// }


// func Add_Substract() (func(int,int) int , func(int,int)int){
// 	add := func(x int , y int)int{
// 		return x+y
// 	}

// 	substract := func(x int, y int)int{
// 		return y-x
// 	}

// 	return add , substract
// }

// func main(){
// 	add ,substract := Add_Substract()
// 	print(5,6,add,substract)
// }


// func square(n []int) []int{
// 	for i , val := range n{
// 		n[i] = val*val
// 	}
// 	return n
// }
// func main(){
// 	fmt.Println("Square of numbers in slice")

// 	num := []int{1,2,3,4,5}

// 	fmt.Println("Numbers before squaring: ", num)

// 	square(num)
// 	fmt.Println("Numbers before squaring: ", num)
// }

func AddSub() (func(int,int)int, func(int,int)int){
	add := func(x , y int)int{
		return x+y
	}
	substract:= func(x , y int)int{
		return y-x
	}

	return add,substract
}

func apply(x int, y int,add func(int,int) int,substract func(int,int) int) (int , int){
	r1 := add(x,y)
	r2 := substract(x,y)

	return r1, r2
}

func main(){
	x := 3
	y := 4

	add , substract := AddSub()

	r1 , r2 := apply(x ,y , add , substract)

	fmt.Println("the result of addition is : ",r1)
	fmt.Println("the result of substraction is: ",r2)
}
