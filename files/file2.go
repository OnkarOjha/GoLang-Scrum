package main

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {
	fmt.Println("File reading and writing part -2")

	file, err := os.Create("test.txt")

	if err != nil {
		panic(err)

	}
	defer file.Close()

	len, err := file.WriteString("Bhai dusri baar try kr rha yaar!!")

	if err != nil {
		panic(err)

	}

	stats, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Permission File Before: %s\n", stats.Mode())
	err = os.Chmod("test.txt", 0444)
	if err != nil {
		log.Fatal(err)
	}

	stats, err = os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Permission File After: %s\n", stats.Mode())

	fmt.Printf("\n File Name: %s", file.Name())
	fmt.Printf("\n Length of file is: %d bytes", len)

}

func ReadFile() {
	fmt.Println("Reading file bhaisahab!!")

	fileName := "test.txt"

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)

	}

	fmt.Println("File name : ", fileName)
	fmt.Println("\n Size : ", len(data))
	fmt.Printf("\n Data : %s", data)

}
func main() {
	CreateFile()
	ReadFile()
}
