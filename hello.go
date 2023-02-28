func updatePersonHandler(w http.ResponseWriter, r *http.Request) {
	// Get the person ID from the URL
	id, err := strconv.Atoi(r.URL.Path[len("/person/"):])
	if err != nil {
	   http.Error(w, err.Error(), http.StatusBadRequest)
	   return
	}
 
	// Find the person with the given ID
	var updatedPerson Person
	for i, person := range people {
	   if person.ID == id {
		  // Parse the request body into a map
		  var updateMap map[string]interface{}
		  err = json.NewDecoder(r.Body).Decode(&updateMap)
		  if err != nil {
			 http.Error(w, err.Error(), http.StatusBadRequest)
			 return
		  }
 
		  // Update the person fields specified in the request
		  if firstName, ok := updateMap["first_name"].(string); ok {
			 person.FirstName = firstName
		  }
		  if lastName, ok := updateMap["last_name"].(string); ok {
			 person.LastName = lastName
		  }
		  if email, ok := updateMap["email"].(string); ok {
			 person.Email = email
		  }
 
		  // Validate the updated person
		  err = validatePerson(person)
		  if err != nil {
			 http.Error(w, err.Error(), http.StatusBadRequest)
			 return
		  }
 
		  // Update the person in the list
		  updatedPerson = person
		  people[i] = person
 
		  // Save the updated people to
 