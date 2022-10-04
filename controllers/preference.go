package controllers

import (
	"B20_Backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func createPreference(userId int) {

	preference := models.BugPreference{}
	result := Db.Where("user_id", userId).Find(&preference)

	if result.RowsAffected == 0 {

		var value = models.BugPreference{
			UserID:        userId,
			Comment_count: true,
			// Deadline:         result.Path("bugs.deadline").Data().([]interface{})[0].(string),
			Type:               true,
			Status:             true,
			Priority:           true,
			Severity:           true,
			Summary:            true,
			Product:            true,
			Platform:           true,
			Resolution:         true,
			Target_milestone:   true,
			Classification:     true,
			Is_confirmed:       true,
			Is_open:            true,
			Last_change_time:   true,
			Creation_time:      true,
			Qa_contact:         true,
			Creator_detail:     true,
			Assigned_to_detail: true,
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

	Db.Model(&preference).Where("user_id = ?", id).Updates(body)

	json.NewEncoder(w).Encode(preference)

}
