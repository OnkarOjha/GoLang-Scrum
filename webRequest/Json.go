package main

import (
	"encoding/json"
	"fmt"
)

// if we keep name of variable lower case then wee know that we are neither exporting it
// nor we are making it public

type course struct{
	Name string 		`json:"CourseName"`
	Price int			`json:"CoursePrice"`
	Platform string 	`json:"CoursePlatform "`
	Password string		`json:"-"` // password won't be visible to anybody
	Tags  []string 		`json:"tags,omitempty"`				// slice of string
	//omitempty just says if the field is null just through empty value there
}

func main(){
	fmt.Println("More into JSON!!")
	// EncodeJson()
	DescodeJson()

}

func EncodeJson(){
	Courses := []course{
		{"React Js" , 299 , "lco.in", "abc@123" , []string{"web" , "dev" , "js"}},
		{"Angular Js" , 199 , "lco.in", "def@123" , nil},
	}

	//package this data as json data
	// finalJson , err := json.Marshal(Courses)
	// if err!=nil{
	// 	panic(err)
	// }

	// similar technique just to deal with indentation
	finalJson , err := json.MarshalIndent(Courses , "" , "\t")
	if err!=nil{
		panic(err)
	}


	fmt.Printf("%s \n",finalJson)
}

func DescodeJson(){
		jsonDataFromWeb := []byte(`
		{
			"CourseName": "React Js",
			"CoursePrice": 299,
			"CoursePlatform ": "lco.in",
			"tags": ["web","dev","js"]
		}
		`)


		var lcoCourse course
		// just to check data that we got is valid or not
		checkValid := json.Valid(jsonDataFromWeb)

		if checkValid{
			fmt.Println("JSON was valid!!!")
			json.Unmarshal(jsonDataFromWeb , &lcoCourse)
			fmt.Printf("%#v \n", lcoCourse)
		}else{
			fmt.Println("JSON s invlaid")
		}

		// some cases wher you jsut want to add data in key value pair
		// whenever in case of json the key is always there but value is not guaranteed as a key value pair it might
		// be array or object or anything so we are creating interaface
		var myOnlineData course
		json.Unmarshal(jsonDataFromWeb , &myOnlineData)
		// gives us output in the name of the alias  by creating interfaces
		fmt.Printf("%#v \n", myOnlineData)
		fmt.Printf("Type oi : %T \n", myOnlineData.Price)

		// looping through the data
		// for k,v := range myOnlineData{
		// 	fmt.Printf(" Key is %v adn value is %v and Type is: %T \n",k , v , v)
		// }
}