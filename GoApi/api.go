package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// creaitng a model for courses

type Course struct{
	CourseId string `json :"courseid"`
	CourseName string `json :"coursename" `
	CoursePrice int `json: "-" `
	Author *Author  `json : "author"` // passing the reference as a pointer so that we are not copying exact same values
}


type Author struct{
	Fullname  string `json: "fullname`
	Website string `json :"website"`
}


//FAKE DB
// as we don't have the actual DB working so we create a slice instead
var courses []Course

// middlewares , helpers - File

func IsEmpty(c *Course) bool {
	// return c.CourseId =="" && c.CourseName==""
	return c.CourseName==""
}

func main(){

	fmt.Println("CRUD operation starts with API!!")
	r := mux.NewRouter()

	courses = append(courses , Course{CourseId : "2" , CourseName : "ReactJS",
	CoursePrice : 299 , Author: &Author{Fullname : "Onkar", Website : "onkar.in"}})

	courses = append(courses , Course{CourseId : "4" , CourseName : "MERN",
	CoursePrice : 499 , Author: &Author{Fullname : "Onkar Ojha", Website : "onkarojha.in"}})

	//routing

	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/create",createOneCourse).Methods("POST")
	r.HandleFunc("/update/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/delete/{id}",deleteOneCourse).Methods("DELETE")
	

	//listen to port
	log.Fatal(http.ListenAndServe(":3000",r))

}


//controllers - file

// serve home route so that anybody who is accessing the localhost should not encounter emtoy pages ri8]

// it is wrapped with a writer and a reader
// writer -> writing responses werever you feel to send responses
//reader ->  it is basically to read the responses from where is it coming from 
func serveHome(w http.ResponseWriter,r *http.Request){
	 // writing ontents on the home
	w.Write([]byte("<h1> Welcome to API onkar !!</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get all Courses!!")

	//method to setting headers
	// these all are key-value pairs
	w.Header().Set("Content-Type" , "application/json")
	json.NewEncoder(w).Encode(courses)// this will send response as a json value 

}

func getOneCourse(w http.ResponseWriter,r *http.Request){

	fmt.Println("Get one course")
	w.Header().Set("Content-Type" , "application/json")

	//grab id from request
	//holder of key-value pair
	params := mux.Vars(r)
	fmt.Printf("Params: \n Type is: %T",params)

	//loop throught he coursess
	// find matching ID and return the response
	for _ , course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given ID!!")
	return 


}

func createOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create one Course!!!")
	w.Header().Set("Content-Type","application-json")

	// what if entire body of the response is empty
	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data!!!")

	}

	//what about {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	// if course.IsEmpty() {
	// 	json.NewEncoder(w).Encode("No data inside JSON")
	// 	return
	// }
	//generate a unique id , string
	// append course into courses

	
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses , course)
	fmt.Println("successfully added new course")
	json.NewEncoder(w).Encode(course)
	return 


}

func UnixNano() {
	panic("unimplemented")
}

func updateOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Update Course according to courseID...")
	w.Header().Set("Content-Type","applicaition/json")

	// first -> grab id from request

	params := mux.Vars(r)

	// loop through the value for the id 
	//remove the index 
	//add the value again with my ID
	// boom!!! update done

	for index , course := range courses{
		if course.CourseId ==  params["id"]{
			courses = append(courses[:index] , courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO :-  send a response when id is not found

}


func deleteOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop find ID
	// perform remove operation
	// append index , index+1

	for index , course := range courses{
		if course.CourseId == params["id"]{
			//remove starts
			courses =append(courses[:index] , courses[index+1:]... )
			fmt.Println("element successfully deleted")
			break // it will break the entire loop
		}
	}
}