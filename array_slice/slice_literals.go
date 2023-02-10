package main

import "fmt"

func main() {
	// q := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(q)

	// r := []bool{true, false, true, true, false, true}
	// fmt.Println(r)

	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }

	s := new([]int)
	fmt.Println(s)
	*s = append( *s, 1)
	fmt.Println(s)
	// printSlice(s)

	// a := append([]byte("hello"),"world" ...)
	// fmt.Println(a)
	// fmt.Printf("the byte slice is %x length= %d , capacity= %d \n",a,len(a),cap(a) )

	// a = append(a,"ojha" ...)
	// fmt.Println(a)
	// fmt.Printf("the byte slice is %x length= %d , capacity= %d \n",a,len(a),cap(a) )

	// b := make([]byte , 10)
	// fmt.Println(b)
	// fmt.Printf("the byte slice is %x length= %d , capacity= %d \n",b,len(b),cap(b) )

	// b =append(b,"ojha"...)
	// fmt.Println(b)
	// fmt.Printf("the byte slice is %x length= %d , capacity= %d \n",b,len(a),cap(a) )


	b :=append([]byte("Tech"),"geek"... )
	fmt.Println()
	fmt.Printf("the byte slice is %x length= %d , capacity= %d \n",b,len(b),cap(b) )

	// s = append([]int{0,10}, 1,2,3 )
	// printSlice(s)

	// s = append(s, 4,5,6)
	// printSlice(s)
}

func printSlice(s []int){
	fmt.Printf("len= %d cap= %d %d \n",len(s) , cap(s) ,s)
}