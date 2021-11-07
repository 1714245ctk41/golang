package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	connection := GetDatabase()
	defer Closedatabase(connection)

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err error
		err2 := fmt.Errorf("Error in readin body: %g", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err2)
		return

	}
	var dbuser User
	connection.Where("email = ?", user.Email).First(&dbuser)

	//!check if email is already register or not
	if dbuser.Email != "" {
		var err error
		err2 := fmt.Errorf("Email already in use: %g", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err2)
		return
	}
	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}
	//* insert user details in database
	connection.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
