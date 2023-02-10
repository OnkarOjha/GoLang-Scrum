package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("json !!")
	// encodeJson()
	decodeJson()
}

func encodeJson() {
	courses := []course{
		{"ReactJS ", 299, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"JS ", 599, "lco.in", "abc123", []string{"go"}},
		{"Angular ", 219, "lco.in", "abc123", nil},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
	"Name": "ReactJS ",
	"Price": 299,
	"Platform": "lco.in",
	"tags": ["web-dev","js"]
	}
	
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Printf("JSON was valid\n")
		json.Unmarshal(jsonDataFromWeb , &lcoCourse)
		fmt.Printf("%#v \n",lcoCourse)
	}else{
		fmt.Println("JSON WAS NOT VALID")
	}
}
