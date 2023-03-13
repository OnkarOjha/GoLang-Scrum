package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Notification struct {
	gorm.Model
	Title   string
	Message string
	UserID  uint
}
type DeletedUser struct {
	Name     string
	Email    string
	Password string
}

var DB *gorm.DB

func main() {
	fmt.Println("CRUD OPERATIONS")
	// initialize the db
	dsn := "host=localhost port=5432 user=postgres password=Onkar17082001 dbname=gorm sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error connecting to the database")
	}
	// defer db.Close()
	db.AutoMigrate(&User{})
	DB = db

	// Add hooks to save deleted user data
	// db.Callback().Delete().After("gorm:before_delete").Register("my_plugin:before_delete", beforeDelete)
	// db.Callback().Delete().After("gorm:after_delete").Register("my_plugin:after_delete", afterDelete)

	r := mux.NewRouter()
	r.HandleFunc("/create", createHandler).Methods("POST")
	r.HandleFunc("/read/{id}", readHandler).Methods("GET")
	r.HandleFunc("/update/{id}", updateHandler).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteHandler).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are creating a new user")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Println("test")

	// Insert user into database
	result := DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Println("test2")

	// Return created user with HTTP status code 201
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	// getting user by id
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("you are reading the user with id", id)

	var user User
	result := DB.First(&user, id)
	if result.Error != nil {

		http.Error(w, "User not found", http.StatusNotFound)

		return
	}

	// Return fetched user with HTTP status code 200
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
func updateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("you are updating the user with id", id)

	result := DB.First(&User{}, id)
	if result.Error != nil {

		http.Error(w, "User not found", http.StatusNotFound)

		return
	}
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// if err := DB.Save(&user).Error; err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }

	DB.Model(&User{}).Where("id = ?", id).Updates(user)
	json.NewEncoder(w).Encode(user)

}
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("you are deleting the user with id", id)
	DB.Delete(&User{}, id)
	w.Write([]byte("User deleted"))

}
