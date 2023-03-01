package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Temp struct {
	T float64 `json:"temp"`
}

type weathyer struct {
	Main Temp `json:"main"`
}

const apiKey = "4bcee6ca2132a9808f9b4fe9f0290ea2"

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// cityInfo := vars["city"]
	city := []string{"mohali", "amritsar" , "delhi"}
	var url string
	for i , val := range city{
		fmt.Println("information", val)
		url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city[i], apiKey)
		fmt.Println("url : ", url)
		// get request
		res, err := http.Get(url)
		if err != nil {
			http.Error(w, "error querrying the url "+url, http.StatusInternalServerError)
		}
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			http.Error(w, "error reading response", http.StatusInternalServerError)
		}
		// fmt.Println("data : ", data)
		var temp weathyer
	
		json.Unmarshal(data, &temp)
		fmt.Println("temp:", temp.Main.T)
		w.Write([]byte(data))
	}

	w.Header().Set("Content-Type", "application/json")

}

func main() {
	fmt.Println("Working with multiple api request ....")

	router := mux.NewRouter()

	router.HandleFunc("/weather", ApiHandler).Methods("GET")
	// router.HandleFunc("/weather/{city}", ApiHandlerChan).Methods("GET")

	http.ListenAndServe(":8080", router)

}
