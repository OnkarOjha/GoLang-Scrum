package main

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct {
	Name   string
	Secret string
}

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	fmt.Println("plr")
	// secret := secret{Name: "Onkar", Secret: "A tech geek"}
	// // filename := "Template.txt"
	// // CreateFile(filename)
	// fmt.Println(secret)

	// templatePath := "C:/Users/hp/Desktop/go-workspace/Template.txt"
	// t , err := template.New("Template.txt").ParseFiles(templatePath)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// err = t.Execute(os.Stdout,secret)
	// if err!=nil{
	// 	fmt.Println(err)
	// }

	
	sweaters := Inventory{"wool", 17}
	 tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	// tmpl, err := template.New("test").Parse({{template "Inventory"}})
	// tmpl, err := template.New("test").Parse("{{if {{23 -}} < {{- 45}}  fmt.Println("It is True")   }}")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }
}

// func CreateFile(filename string) {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer file.Close()
// }
