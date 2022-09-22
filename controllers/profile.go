package controllers

import (
	"B20_Backend/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func EditProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	Db.Model(&user).Where("user_id = ?", user.UserId).Updates(&user)

	json.NewEncoder(w).Encode(user)
}

func GetProfileInfoLocal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var user models.User
	Db.Where("user_id = ?", id).Find(&user)

	json.NewEncoder(w).Encode(user)

}
