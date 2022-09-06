package main

import (
	"bugzilla/apis/models"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	byte_password := []byte(password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(byte_password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	DB.Create(&user)

	json.NewEncoder(w).Encode(user)

}
