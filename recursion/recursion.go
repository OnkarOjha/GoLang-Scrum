package main

import "fmt"
func factorial(i int)int{
	if(i<=1){
		return i
	}

	return i * factorial(i-1)
}

func fibonacci(i int)int{
	if i ==0 {
		return 0
	} 
	if i==1{
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main(){
	fmt.Println("recursion!!")
	a := 10
	fmt.Printf("Factorial of %d is %d",a , factorial(a))
	// fmt.Println(fibonacci(10))

	for i :=0;i<a;i++{
		fmt.Println(fibonacci(i))
	}

}