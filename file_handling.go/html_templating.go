package main

import (
	"fmt"
	"html/template"
	"os"
)

type Html struct{
	Header string
	Paragraph string
}
func main(){
	fmt.Println("HTML Templating")
	message := Html{Header: "Onkar" , Paragraph: "this is a paragraph"}

	templatePath := "C:/Users/hp/Desktop/go-workspace/Template.html"

	t , err := template.New("Template.html").ParseFiles(templatePath)
	if err!=nil{
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout,message )
	if err!=nil{
		fmt.Println(err)
	}
}