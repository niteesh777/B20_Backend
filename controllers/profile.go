package controllers

import (
	"B20_Backend/models"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func EditProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	user := &models.User{}
	contactInfo := &models.ContactInfo{}
	err := json.NewDecoder(r.Body).Decode(user)

	var k []byte
	l := bytes.NewBuffer(k)
	json.NewEncoder(l).Encode(user)
	err = json.NewDecoder(l).Decode(contactInfo)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	if Db.Model(&contactInfo).Where("id = ?", id).Updates(&contactInfo).Error == nil {
		if Db.Model(&user).Where("user_id = ?", id).Updates(&user).Error == nil {
			var user models.User
			Db.Where("user_id = ?", id).Preload("User").Find(&user)
			json.NewEncoder(w).Encode(user)
		} else {

			json.NewEncoder(w).Encode("error")

		}
	} else {
		json.NewEncoder(w).Encode("error")
	}

}

func GetProfileInfoLocal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var user models.User
	Db.Where("user_id = ?", id).Preload("User").Find(&user)

	json.NewEncoder(w).Encode(user)

}
