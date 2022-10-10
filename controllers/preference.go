package controllers

import (
	"B20_Backend/models"
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"github.com/gorilla/mux"
)

func createPreference(userId int) {

	preference := models.BugPreference{}
	result := Db.Where("user_id", userId).Find(&preference)

	if result.RowsAffected == 0 {

		var value = models.BugPreference{
			UserID:        userId,
			Comment_count: false,
			// Deadline:         result.Path("bugs.deadline").Data().([]interface{})[0].(string),
			Type:               false,
			Status:             false,
			Priority:           false,
			Severity:           false,
			Summary:            false,
			Product:            false,
			Platform:           false,
			Resolution:         false,
			Target_milestone:   false,
			Classification:     false,
			Is_confirmed:       false,
			Is_open:            false,
			Last_change_time:   false,
			Creation_time:      false,
			Qa_contact:         false,
			Creator_detail:     false,
			Assigned_to_detail: false,
		}

		Db.Create(&value)

	}

}

func CreatePreferenceUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var userIds []int

	Db.Model(&models.User{}).Select("user_id").Find(&userIds)

	for _, j := range userIds {
		createPreference(j)
	}

	json.NewEncoder(w).Encode("created default preferences for users")

}

func GetPreference(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	preference := models.BugPreference{}
	result := Db.Where("user_id = ?", id).Find(&preference)

	if result.RowsAffected == 0 {

		userid, _ := strconv.Atoi(id)
		createPreference(userid)
		Db.Where("user_id = ?", id).Find(&preference)

	}

	json.NewEncoder(w).Encode(preference)

}

func EditPreference(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	preference := &models.BugPreference{}

	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	fmt.Println(id);

	Db.Model(&preference).Where("user_id = ?", id).Updates(body)

	json.NewEncoder(w).Encode(preference)

}
