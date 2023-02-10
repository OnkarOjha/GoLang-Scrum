// package main

// import "fmt"

// func main() {
// 	names := [4]string{
// 		"John",
// 		"Paul",
// 		"George",
// 		"Ringo",
// 	}
// 	fmt.Println(names)

// 	a := names[0:2] // John , Paul
// 	b := names[1:3] // Paul , George
// 	fmt.Println(a, b)
// 	//TODO if we make changes to any copy of the slice then that change can be observed in other copies as well
// 	b[0] = "XXX"   //[John XXX] [XXX George]  Paul pura XXX ho gya 

// 	fmt.Println(a, b)
// 	fmt.Println(names)
// }

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Initialize a 2D slice of uint8
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
	}

	// Fill the 2D slice with values
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			pic[i][j] = uint8(i^j)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
