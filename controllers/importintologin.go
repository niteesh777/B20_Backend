package controllers

import (
	"B20_Backend/models"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func ImportintoLogin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var users []models.ContactInfo
	result := Db.Find(&users)
	fmt.Println(result.RowsAffected)
	var info models.User

	for _, user := range users {
		response := Db.First(&info, user.Id)
		if response.RowsAffected == 0 {
			byte_password := []byte(user.Nick)

			// Hashing the password with the default cost of 10
			hashedPassword, err := bcrypt.GenerateFromPassword(byte_password, bcrypt.DefaultCost)
			if err != nil {
				panic(err)
			}

			userInfo := models.User{
				User:     user,
				Name:     user.Real_name,
				Email:    user.Email,
				Password: string(hashedPassword),
			}

			Db.Create(&userInfo)
		}
	}
	json.NewEncoder(w).Encode(users)
}
