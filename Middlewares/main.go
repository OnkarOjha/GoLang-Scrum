package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type NewsAggPage struct {
	Title string
	News  string
}

type Product struct {
	Id       string
	Name     string
	Price    float64
	Quantity int
	Status   bool
}

func main() {
	myHandler := http.HandlerFunc(homeHandler)
	http.HandleFunc("/", homeHandler) // to call the handler and assigning it the specific route in which we wanna serve
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/products", productHandler)
	http.HandleFunc("/html", htmlHandler)
	http.Handle("/middleware", middlewareHandler(myHandler))
	http.HandleFunc("/upper", templateHandler)

	http.ListenAndServe(":8000", nil) // to starting the server at the spcified port number

	//serving html files

}

// to create a to upper fucntion that will amke our txt to upper case
// adding functionalitis such as addition  , to upper , top lower , calculating the total amout , telling that the specific product is under stck or out of stock

func templateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing Template Handler")
	data := map[string]interface{}{
		"fullname": "biggy bee",
		"products": Product{
			Id:       "p01",
			Name:     "Product 1",
			Price:    400,
			Quantity: 10,
			Status:   true,
		},
	}
	tmp, _ := template.New("html/index.html").Funcs(template.FuncMap{
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
		"displayStatus": func(status bool) string {
			if status {
				return "Show Status"
			} else {
				return "Hide Status"
			}
		},
		"total": func(p Product) float64 {
			return p.Price + float64(p.Quantity)
		},
		"checkStock": func(p Product) string {
			if p.Quantity == 0 {
				return "Out of Stock"
			} else {
				return "In Stock"
			}
		},
	}).ParseFiles("html/index.html")

	tmp.Execute(w, data)
}

// func upperHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Executing the template handler!!")
// 	data := struct {
// 		Title string
// 		Body  string
// 	}{
// 		Title: "Title",
// 		Body:  "onkar",
// 	}
// 	tmp, err := template.New("index").Funcs(template.FuncMap{
// 		"upper": func(s string) string {
// 			return strings.ToLower(s)
// 		},
// 	}).Parse("<html><head><title>{{.Title}}</title></head><body>Full Name : {{upper .Body}}</body></html>")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	err = tmp.Execute(w, data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

func middlewareHandler(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running before main handler")
		w.Write([]byte("Hijacking Request "))
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Running after middleware handler")
	})
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	// p := NewsAggPage{Title: "this is the title",
	// 	News: "this is the news"}

	t, _ := template.ParseFiles("html/main.html")
	t.Execute(w, t)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing productHandler")
	fmt.Fprintf(w, "Hy we are handling products over here")
}


func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("executing userHandler")
	fmt.Fprintf(w, "Hy we are handling users over here")
}
//TODO HOME HANDLER
var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles(

		"html/mytemplate.html",
		"html/home.html",
	))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing the Home handler")
	templates.ExecuteTemplate(w, "layout", nil)
}
