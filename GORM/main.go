package main

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string `gorm:"default:uuid_generate_v3()"` // db func
	Name     string
	Email    string
	Age      int
	Birthday time.Time
	Password string
	FullName  string `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
}

type DeletedUser struct {
	gorm.Model
	Name     string
	Email    string
	Age      int
	Birthday time.Time
	Password string
}

// Define a Notification struct
type Notification struct {
	gorm.Model
	Title   string
	Message string
	UserID  uint
}

// Define a BeforeCreate hook for the User struct
// func (user *User) BeforeCreate(tx *gorm.DB) error {
// 	// set default values
// 	// user.CreatedBy = "onkar"

// 	// hash the password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err!=nil{
// 		panic("failed to hash password")
// 	}
// 	user.Password = string(hashedPassword)
// 	return nil
// }

// Define a AfterCreate hook for the User struct
// func (user *User) AfterCreate(tx *gorm.DB) error {
// 	// Log the creation of the new user
// 	log.Println("User created:", user.ID)

// 	// Send a notification to the user
// 	notification := Notification{
// 		Title:   "Account Created",
// 		Message: "Your account has been successfully created.",
// 		UserID:  user.ID,
// 	}
// 	tx.Create(&notification)

// 	return nil

// }

// Create a new user and trigger the hooks
// func createUser(db *gorm.DB) {
// 	user := User{
// 		Name:  "hookswithpassword",
// 		Email: "hookswithpassword@example.com",
// 		Age:  323,
// 		Password: "hjsfjk",
// 	}

// 	// Create the user and trigger the hooks
// 	db.Create(&user)
// }

// Define an AfterDelete hook for the User struct
func (user *User) AfterDelete(tx *gorm.DB) error {
	// Save the deleted user data to the DeletedUser table
	deletedUser := DeletedUser{
		Name:     user.Name,
		Email:    user.Email,
		Age:      user.Age,
		Password: user.Password,
		Birthday: user.Birthday,
	}
	notification := Notification{
		Title:   "User deleted",
		Message: "Your user  has been successfully deleted.",
		UserID:  user.ID,
	}
	tx.Create(&notification)
	tx.Create(&deletedUser)

	return nil
}

func main() {
	dsn := "host=localhost user=postgres password=Onkar17082001 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Notification{}, &DeletedUser{})

	// Delete a user and trigger the AfterDelete hook
	user := User{}
	db.First(&user, 9) // Assume user with ID 1 exists
	db.Delete(&user)

	// // // Migrate the schema
	// err = db.AutoMigrate(&User{} , &Notification{})
	// if err != nil {
	// 	panic("failed to migrate database")
	// }
	// createUser(db)

	// // Create a new user
	// user := User{Name: "onkar", Email: "onkar@example.com"}
	// db.Create(&user)

	// // delete user by name
	// var userToDelete User
	// db.First(&userToDelete, "name =?", "onkar")
	// db.Delete(&userToDelete)

	// delete user by id
	// db.Delete(&User{},3)

	// // Find a user
	// var foundUser User
	// db.First(&foundUser, "name = ?", "onkar")
	// fmt.Printf("Found user: %v", foundUser)

	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	// result := db.Create(&user) // pass pointer of data to Create
	// fmt.Println("result: ", result)

	// // user.ID             // returns inserted data's primary key
	// // result.Error        // returns error
	// // result.RowsAffected // returns inserted records count
	// var users = []User{{Name: "tech"}, {Name: "aman"}, {Name: "prajwal"}}
	// db.Create(&users)

	// for _, user := range users {
	// 	user.ID // 1,2,3
	// }

}
