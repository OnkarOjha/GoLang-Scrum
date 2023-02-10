package main

import "fmt"

//TODO function for wint

// func booyahDynamic2(b [3][3]int)string{
// 	//TODO handling rows
// 	for i:=0;i<3;i++{
// 		for j:=0;j<3;j++{

// 		}
// 	}
// }

func booyahDynamic(b [3][3]int) string {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			//TODO row check 0/1
			if b[i][j] == b[i][j+1] && b[i][j+1] == b[i][j+2] {
				if b[i][j+2] == 0 {
					return "0 wins"
				}
				return "1 wins"

			}
			if b[i+1][j] == b[i+1][j+1] && b[i+1][j+1] == b[i+1][j+2] {
				if b[i+1][j+2] == 0 {
					return "0 wins"
				}
				return "1 wins"
			}
			if b[i+2][j] == b[i+2][j+1] && b[i+2][j+1] == b[i+2][j+2] {
				if b[i+2][j+2] == 0 {
					return "0 wins"
				}
				return "1 wins"
			}

			//TODO column check 0/1
			if b[i][j] == b[i+1][j] && b[i+1][j] == b[i+2][j] {
				if b[i+2][j] == 0 {
					return "0 wins"
				}
				return "1 wins"
			}
			if b[i][j+1] == b[i+1][j+1] && b[i+1][j+1] == b[i+2][j+1] {
				if b[i+2][j+1] == 0 {
					return "0 wins"
				}
				return "1 wins"
			}
			if b[i][j+2] == b[i+1][j+2] && b[i+1][j+2] == b[i+2][j+2] {
				if b[i+2][j+2] == 0 {
					return "0 wins"
				}
				return "1 wins"
			}

			// TODO FOR 1 diagonal matching
			if b[0][0] == b[1][1] && b[1][1] == b[2][2] && b[2][2] == 1 {
				if b[2][2] == 0 {
					return "0 wins"
				}
				return "1 wins"
				
			}
			if b[0][2] == b[1][1] && b[1][1] == b[2][0] && b[2][0] == 1 {
				if b[2][0] == 0 {
					return "0 wins"
				}
				return "1 wins"
				
			}

			// // TODO FOR 0 diagonal matching
			// if b[0][0] == b[1][1] && b[1][1] == b[2][2] && b[2][2] == 0 {
			// 	return "0 won"
			// }
			// if b[0][2] == b[1][1] && b[1][1] == b[2][0] && b[2][0] == 0 {
			// 	return "0 won"
			// }

		}
	}

	return "MATCH DRAW"

}

//function for win

func booyah(b [3][3]int) string {
	// TODO FOR 1 row matching
	if b[0][0] == b[0][1] && b[0][1] == b[0][2] && b[0][2] == 1 {
		return "1 won"
	}
	if b[1][0] == b[1][1] && b[1][1] == b[1][2] && b[1][2] == 1 {
		return "1 won"
	}
	if b[2][0] == b[2][1] && b[2][1] == b[2][2] && b[2][2] == 1 {
		return "1 won"
	}

	// TODO FOR 0 row matching
	if b[0][0] == b[0][1] && b[0][1] == b[0][2] && b[0][2] == 0 {
		return "0 won"
	}
	if b[1][0] == b[1][1] && b[1][1] == b[1][2] && b[1][2] == 0 {
		return "0 won"
	}
	if b[2][0] == b[2][1] && b[2][1] == b[2][2] && b[2][2] == 0 {
		return "0 won"
	}

	// TODO FOR 1 column matching
	if b[0][0] == b[1][0] && b[1][0] == b[2][0] && b[2][0] == 1 {
		return "1 won"
	}
	if b[0][1] == b[1][1] && b[1][1] == b[2][1] && b[2][1] == 1 {
		return "1 won"
	}
	if b[0][2] == b[1][2] && b[1][2] == b[2][2] && b[2][2] == 1 {
		return "1 won"
	}

	// TODO FOR 0 column matching
	if b[0][0] == b[1][0] && b[1][0] == b[2][0] && b[2][0] == 0 {
		return "0 won"
	}
	if b[0][1] == b[1][1] && b[1][1] == b[2][1] && b[2][1] == 0 {
		return "0 won"
	}
	if b[0][2] == b[1][2] && b[1][2] == b[2][2] && b[2][2] == 0 {
		return "0 won"
	}

	// TODO FOR 1 diagonal matching
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] && b[2][2] == 1 {
		return "1 won"
	}
	if b[0][2] == b[1][1] && b[1][1] == b[2][0] && b[2][0] == 1 {
		return "1 won"
	}

	// TODO FOR 0 diagonal matching
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] && b[2][2] == 0 {
		return "0 won"
	}
	if b[0][2] == b[1][1] && b[1][1] == b[2][0] && b[2][0] == 0 {
		return "0 won"
	}

	return "MATCH DRAW	"
}

func main() {

	fmt.Println("Let's Play Tic Tac Toe!!")
	board := [3][3]int{
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 0},
	}

	fmt.Println(board)
	// fmt.Println("And the result is: ",booyah(board))
	fmt.Println("And the result is: ", booyahDynamic(board))

}
