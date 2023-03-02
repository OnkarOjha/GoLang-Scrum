package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// type Info struct{
// 	// Weather WeatherInfo `json:"Weather`
// 	Name string `json:"name"`
// }

var Info string

type ResData struct {
	Temperature float64 `json:"tempr"`
	Message     string  `json:"message"`
}
type WeatherInfo struct {
	Name string  `json:"name"`
	Temp float64 `json:"temp"`
	Diff string  `json:"diff"`
}

func main() {
	http.HandleFunc("/weather", weatherHandler)
	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := "4bcee6ca2132a9808f9b4fe9f0290ea2"
	params := r.URL.Query()
	var cities []string
	ffd := make(map[string]ResData)
	for _, city := range params {
		cities = append(cities, city[0])
	}
	fmt.Println(cities)

	weatherInfos := make([]WeatherInfo, len(cities))
	// var finalMap map[string]WeatherInfo

	for i, city := range cities {
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		var data map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		temp := data["main"].(map[string]interface{})["temp"].(float64)

		weatherInfos[i] = WeatherInfo{
			Name: city,
			Temp: temp,
		}

	}

	for i, wi := range weatherInfos {
		diff := ""
		for j, wj := range weatherInfos {
			if i != j {
				tempDiff := wi.Temp - wj.Temp
				if wi.Temp > wj.Temp {
					diff += fmt.Sprintf("%s is %.1f°C hotter than %s, ", wi.Name, tempDiff, wj.Name)
				} else if wi.Temp < wj.Temp {
					diff += fmt.Sprintf("%s is %.1f°C cooler than %s, ", wi.Name, -tempDiff, wj.Name)

				}
			}
		}
		weatherInfos[i].Diff = strings.TrimSuffix(diff, ", ")
		fmt.Println("name: ", wi.Name)
		ffd[wi.Name] = ResData{
			Temperature: wi.Temp,
			Message:    diff,
		}
		// finalMap[wi.Name] = weatherInfos[i]
	}
	fmt.Println("ffd : " , ffd)
	// fmt.Println("weather infos: ", weatherInfos)
	// fmt.Println("final info: ", finalMap)

	// make a map

	// for i ,wi := range weatherInfos {
	// 	finalMap[wi.Name] = append(finalMap[wi.Name] , weatherInfos[i])
	// }

	responsJSON, err := json.Marshal(ffd)
	// err = json.NewEncoder(w).Encode(ffd)
	// fmt.Println("respoinse" , string(responseJSON))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responsJSON)
}

// func roundFloat(val float64, precision uint) float64 {
// 	ratio := math.Pow(10, float64(precision))
// 	return math.Round(val*ratio) / ratio
// }
