// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// // import "log"

// // // func main() {
// // // 	f, err := fmt.Println("errors")
// // // 	if err!= nil {
// // //         fmt.Print(err)
// // //     }
// // // 	fmt.Println("errors are:  ",f)
// // // }

// // // func main() {
// // // f, err := os.Create("names.txt")
// // // if err != nil {
// // // 	fmt.Println(err)
// // // }
// // // defer f.Close()

// // // msg := strings.NewReader("Helllo")
// // // io.Copy(f, msg)

// // // 	openfile(f)
// // // }
// // // func openfile(f *os.File){
// // // 	file , err := os.Open("names.txt")
// // // 	if err!=nil{
// // // 		fmt.Println(err)
// // // 		return
// // // 	}

// // // 	defer file.Close()

// // // 	readFileByte , err := io.ReadAll(file)
// // // 	if err!=nil{
// // // 		fmt.Println(err)
// // // 		return

// // // 	}
// // // 	fmt.Println(string(readFileByte))
// // // }

// // func main() {
// // 	defer foo()
// // 	file, err := os.Open("names.txt")
// // 	if err != nil {
// // 		log.Panicln(err)

// // 	}

// // 	defer file.Close()

// // 	readFileByte, err := io.ReadAll(file)
// // 	if err != nil {
// // 		log.Panicln(err)
// // 	}
// // 	fmt.Println(string(readFileByte))
// // 	// log.Println()
// // }
// // func foo(){
// // 	fmt.Println("ky pt a")
// // }

// // package main

// // import (
// //     "log"
// //     "os"
// // )

// func main() {
//     file, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//     if err != nil {
//         log.Fatal("Failed to open log file:", err)
//     }
//     defer file.Close()

//     log.SetOutput(file)

//     log.Println("This message will be written to the log file")
//     log.Println("This message will be written to the log file")
//     log.Println("This message will be written to the log file")

// }

package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
