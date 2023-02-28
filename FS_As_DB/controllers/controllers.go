package cont

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "regexp"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name" validation:"required"`
	Gender  string `json:"gender" validation:"required"`
	Email   string `json:"email" validation:"required"`
	Address string `json:"address"`
}

var People []Person

//used for loading file information that what exactly it contains
func Loadfile() {
	//this function will load the file with all the json data
	fmt.Println("We are loading persons information...")
	file, err := os.Open("data/people.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// now read data from JSON file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data from JSON file: ", string(data))
}

//TODO applying validation to person struct
func (p Person) Validate() error {
	return validation.ValidateStruct(&p,
		// validation.Field(&p.Id , validation.Match(regexp.MustCompile(`^[0-9]{6}$`))),
		validation.Field(&p.Name, validation.Length(1, 255)),
		validation.Field(&p.Gender,  validation.In("male", "female")),
		validation.Field(&p.Email,  is.Email),
		validation.Field(&p.Address, validation.Required, validation.Length(1, 255)),
	)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are creating a new client")

	var newPerson Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//? validating the new person added
	err = newPerson.Validate()
	if err != nil {
		http.Error(w, "validation error "+err.Error(), http.StatusBadRequest)
	}

	// Add the new person to the list
	newPerson.Id = len(People) + 1
	People = append(People, newPerson)

	//? Now use fileserver to save the new person
	file, err := os.Open("data/people.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// now encode data as json
	data, err := json.Marshal(People)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("data/people.json", data, 0644)

	// Return the new person as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPerson)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are updating a client")
	// get persons id from URL path
	var id int
	fmt.Sscanf(r.URL.Path, "/update/%d", &id)
	fmt.Println("id:", id)

	var personToUpdate Person
	for i, person := range People {
		if person.Id == id {
			var updateMap map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&updateMap)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			//  now just make some if conditions to check which information you
			// exaclty wanna change like name , email , address
			if name, ok := updateMap["name"]; ok {
				person.Name = name.(string)
				//TODO
				validation.ValidateStruct(&person,validation.Field(&person.Name, validation.Required, validation.Length(1, 255)))
				
			}
			if gender, ok := updateMap["gender"]; ok {
				person.Gender = gender.(string)
				validation.ValidateStruct(&person,validation.Field(&person.Gender, validation.Required, validation.In("male", "female")))

			}
			if email, ok := updateMap["email"]; ok {
				person.Email = email.(string)
				validation.ValidateStruct(&person,validation.Field(&person.Email, validation.Required, is.Email))

			}
			if address, ok := updateMap["address"]; ok {
				person.Address = address.(string)
				validation.ValidateStruct(&person,validation.Field(&person.Address, validation.Required, validation.Length(1, 255)))

			}
			personToUpdate = person
			People[i] = person
		}
	}

	//?  now open the file and overwrite the updated user data into it
	file, err := os.OpenFile("data/people.json", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	data, err := json.Marshal(People)
	if err != nil {
		http.Error(w, "error marshalling", http.StatusInternalServerError)
		return
	}

	err = ioutil.WriteFile("data/people.json", data, 0644)
	if err != nil {
		http.Error(w, "error writing", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(personToUpdate)
	fmt.Fprintf(w, "Person Updated successfully")
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are deleting a client")
	var id int
	fmt.Sscanf(r.URL.Path, "/delete/%d", &id)
	fmt.Println("id:", id)
	// range through peoples and delete iten
	for i, p := range People {
		if p.Id == id {
			People = append(People[:i], People[i+1:]...)
			break
		}
	}

	// ? Now use fileserver to delete that data from that as well
	file, err := os.Open("data/people.json")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()
	data, err := json.Marshal(People)
	if err != nil {
		panic(err)
		return
	}
	err = ioutil.WriteFile("data/people.json", data, 0644)
	if err != nil {
		panic(err)
		return
	}
	// Return the new person as JSON
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(People)

	// w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Person deleted successfully")
}

func ReadClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are reading a client")
	// Return all people as JSON
	// Loadfile()

	//? reading data from file
	file, err := os.Open("data/people.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// now read data from JSON file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("data :", string(data))
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(string(data))
	//? byte data uthaya usko JSON form me convert kra and People me store kra

	json.Unmarshal(data, &People)
	// ? isne whi json format uthaya or usko byte slice mei convert kra
	p, _ := json.Marshal(People)

	w.Write([]byte(p))
}
