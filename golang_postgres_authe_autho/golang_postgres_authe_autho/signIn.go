package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	connection := GetDatabase()
	defer Closedatabase(connection)

	var authdetails Authentication
	err := json.NewDecoder(r.Body).Decode(&authdetails)
	if err != nil {
		var err error
		err2 := fmt.Errorf("Error in reading body: %g", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err2)
		return
	}
	var authuser User
	connection.Where("email = ?", authdetails.Email).First(&authuser)
	if authuser.Email == "" {
		var err error
		err2 := fmt.Errorf("Username or Password is incorrect: %g", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err2)
		return
	}
	check := CheckPasswordHash(authdetails.Password, authuser.Password)

	if !check {
		var err error
		err2 := fmt.Errorf("Password is incorect: %g", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err2)
		return
	}
	validToken, err := GenerateJWT(authuser.Email, authuser.Role)
	if err != nil {
		var err error
		err2 := fmt.Errorf("Failed to generate token: %g", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err2)
		return
	}
	var token Token
	token.Email = authuser.Email
	token.Role = authuser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)

}
