package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	secretkey := "mysecretkey"
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			var err error
			err2 := fmt.Errorf("Password is incorect: %g", err)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err2)
			return
		}
		var mySigningKey = []byte(secretkey)
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, err) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err error
			err2 := fmt.Errorf("Your Token has been expired: %g", err)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err2)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return
			} else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			}

		}
		var reserr error
		reserr = fmt.Errorf("Your Token has been expired: %g", reserr)
		json.NewEncoder(w).Encode(err)
	}
}
