package main

import (
	"fmt"
	

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// todo VALIDATION WITH STRUCT

// type User struct {
// 	Name     string
// 	Email    string
// 	Password string
	
// }
// type Address struct {
// 	Street  string
// 	State   string
// 	Country string
// 	ZipCode string
// }

// func (u User) Validate() error {
// 	return validation.ValidateStruct(&u,
// 		validation.Field(&u.Name, validation.Required),
// 		validation.Field(&u.Email, validation.Required , validation.Match(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`))),
// 		validation.Field(&u.Password, validation.Required, validation.Length(8, 0)),
// 	)
// }

// func (a Address) validateAddress() error {
// 	return validation.ValidateStruct(&a,
// 		// simple validation
// 		validation.Field(&a.Street, validation.Required),
// 		// validate state that it should contain at least 2 upper case letters
// 		validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[a-z]*[A-Z][a-z]*$"))),
// 		// validate country that it should contain at least 1 upper case letters
// 		validation.Field(&a.Country, validation.Required, validation.Match(regexp.MustCompile("^[a-z]*[A-Z][a-z]*$"))),
// 		//validate zip code that it should contain at least 6 digits
// 		validation.Field(&a.ZipCode, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{6}$"))),
// 	)
// }

// func main() {
// 	user := User{Name: "John", Email: "john@example.com", Password: "password"}
// 	address := Address{Street: "123 Main Street", State: "Jharkhand", Country: "India", ZipCode: "828135"}
// 	err := user.Validate()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// validating address
// 	err = address.validateAddress()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(user)
// 	fmt.Println(address)


// 	// checking internal error and differentiating it with validation error
// 	if err := user.Validate(); err != nil {
// 		if e, ok := err.(validation.InternalError); ok {
// 			// an internal error happened
// 			fmt.Println(e.InternalError())
// 		}
// 	}
// }

//TODO regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`) REGEX to match email addresses

// TODO validation with MAP



// func main(){
// 	fmt.Println("Validation with map")
	// c := map[string]interface{}{
	// 	"name" : "Onkar",
	// 	"email": "onkar@gmail.com",
	// 	"age" : 22,
	// 	"address" : map[string]interface{}{
	// 		"country" : "India",
	// 		"state" : "Jharkhand",
	// 		"pincode" : "828135",
	// 	},
	// }

	// // now lets validate the map	
	// err := validation.Validate(c,
	// 	validation.Map(
	// 		validation.Key("name",validation.Required , validation.Length(5,20)),
	// 		validation.Key("email" , validation.Required, validation.Match(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`))),
	// 		validation.Key("age", validation.Required, validation.Min(1) , validation.Max(100)),
	// 		validation.Key("address",
	// 		validation.Map(
	// 			validation.Key("country",validation.Required , validation.Match(regexp.MustCompile(`^[a-z]*[A-Z][a-z]*$`))),
	// 			validation.Key("state",validation.Required , validation.Match(regexp.MustCompile(`^[a-z]*[A-Z][a-z]*$`))),
	// 			validation.Key("pincode",validation.Required , validation.Match(regexp.MustCompile(`^[0-9]{6}$`))),


	// 		),
	// 	),
	// 	),
	// )
	// if err!=nil{
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("and we have c: ",c)


	// c := map[string]interface{}{
	// 	"Name":  "Onkar Ojha",
	// 	"Email": "q@gmail.com",
	// 	"Address": map[string]interface{}{
	// 		"State":  "Virginia",
	// 	},
	// }
	
	// err := validation.Errors{
	// 	"Name": validation.Validate("Name", validation.Required, validation.Length(5, 20)),
	// 	"Email": validation.Validate("Email", validation.Required, is.Email),
	// 	// "zip": validation.Validate(c.Address.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	// }.Filter()

	
	// fmt.Println(err)
	// fmt.Println("c is :",c)

// }


type Customer struct {
	Name    string
	Gender  string
	Email   string
	Address Address
}
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}


func (c Customer) Validate() error {
	return validation.ValidateStruct(&c,
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Field(&c.Name, validation.Required, validation.Length(5, 20)),
		// Gender is optional, and should be either "Female" or "Male".
		validation.Field(&c.Gender, validation.In("Female", "Male")),
		// Email cannot be empty and should be in a valid email format.
		validation.Field(&c.Email, validation.Required, is.Email),
		// Validate Address using its own validation rules
		validation.Field(&c.Address),
	)
}
func main(){

	c := Customer{
		Name:  "Qiang Xue",
		Gender: "Male",
		Email: "q@gmail.com",
		Address: Address{
			Street: "123 Main Street",
			City:   "Unknown",
			State:  "Virginia",
			Zip:    "12345",
		},
	}
	
	err := c.Validate()
	fmt.Println(err)

	// we can also apply conditional validation
	// result := validation.ValidateStruct(&a,
	// 	validation.Field(&a.Unit, validation.When(a.Quantity != "", validation.Required).Else(validation.Nil)),
	// 	validation.Field(&a.Phone, validation.When(a.Email == "", validation.Required.Error("Either phone or Email is required.")),
	// 	validation.Field(&a.Email, validation.When(a.Phone == "", validation.Required.Error("Either phone or Email is required.")),
	// ))
	
}