package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {

		log.Println(err)
	}

	data := struct {
		Method string
		URL *url.URL
		Subm
	}{
		req.Form,
	}
	tql.Execute(w, "index.gohtml", data)
}

var tql *template.Template

func init(){
	tql = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8000",d)
}
