package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("We are handling API calls"))
}

// const url = "https://api.openweathermap.org/data/2.5/forecast?id=524901&appid=4bcee6ca2132a9808f9b4fe9f0290ea2"
const apiKey = "4bcee6ca2132a9808f9b4fe9f0290ea2"

func ApiHandlerChan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityInfo := vars["city"]
	fmt.Println("information", cityInfo)
	// url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", cityInfo, apiKey)
	// fmt.Println("url : ", url)
	errChan := make(chan error)
	dataChan := make(chan string, 1)
	readChan := make(chan []byte)
	dataChan <- fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", cityInfo, apiKey)
	// dataChan <- "http://api.openweathermap.org/data/2.5/forecast?id=524901&appid=4bcee6ca2132a9808f9b4fe9f0290ea2"

	go func() {

		res, err := http.Get(<-dataChan)
		if err != nil {
			errChan <- errors.New("Error while loading data in datachan")
		}
		defer res.Body.Close()

		read, err := ioutil.ReadAll(res.Body)
		if err != nil {
			errChan <- errors.New("Error while Reading  data in datachan")
		}
		// fmt.Println("data", string(read))
		readChan <- read
		// /fmt.Println("data channel: ", dataChan)
	}()

	select {
	case err := <-errChan:
		fmt.Println("There was an error: ", err)
	case data := <-readChan:
		w.Write([]byte(data))
	case <-time.After(1 * time.Millisecond):
		w.Write([]byte("timeout"))
	}

}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cityInfo := vars["city"]
	fmt.Println("information", cityInfo)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", cityInfo, apiKey)
	fmt.Println("url : ", url)
	// get request
	res, err := http.Get(url)
	if err != nil {
		http.Error(w, "error querrying the url "+url, http.StatusInternalServerError)
	}
	// fmt.Println("response", res.Body)

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "error reading response", http.StatusInternalServerError)
	}
	fmt.Println("data : ", string(data))
	w.Write([]byte(data))

	// err = json.NewDecoder(res.Body).Decode(&struct)
	// if err != nil {
	// 	panic(err)
	// }

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(strutc)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", ApiHandlerChan).Methods("GET")

	router.HandleFunc("/weather/{city}", ApiHandler).Methods("GET")
	// router.HandleFunc("/weather/{city}", ApiHandlerChan).Methods("GET")

	http.ListenAndServe(":8080", router)

}

// ^ API KEY -> 4bcee6ca2132a9808f9b4fe9f0290ea2
