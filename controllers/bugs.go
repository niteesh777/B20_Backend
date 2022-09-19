package controllers

import (
	"B20_Backend/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Url = "https://bugzilla.mozilla.org/rest/"

func GetBug(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	getUrl := Url + "bug?id=" + params["id"]

	fmt.Println(getUrl)

	response, err := http.Get(getUrl)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resmar any
	json.Unmarshal(responseData, &resmar)
	json.NewEncoder(w).Encode(resmar)

}

func GetAssignedBugs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]

	var bugs []models.Bug
	Db.Where("assigned_to_detail_id = ?", userId).Preload("Qa_contact").Preload("Creator_detail").Preload("Assigned_to_detail").Find(&bugs)

	json.NewEncoder(w).Encode(bugs)

}

func GetCreatedBug(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]

	var bugs []models.Bug
	Db.Where("creator_detail_id = ?", userId).Preload("Qa_contact").Preload("Creator_detail").Preload("Assigned_to_detail").Find(&bugs)

	json.NewEncoder(w).Encode(bugs)

}

func GetRelatedBug(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]

	var bugs []models.Bug
	Db.Where("qa_contact_id = ?", userId).Preload("Qa_contact").Preload("Creator_detail").Preload("Assigned_to_detail").Find(&bugs)

	json.NewEncoder(w).Encode(bugs)

}

func GetAllBugs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]

	var bugs []models.Bug
	Db.Where("assigned_to_detail_id = ?", userId).Or("creator_detail_id = ?", userId).Or("qa_contact_id = ?", userId).Preload("Qa_contact").Preload("Creator_detail").Preload("Assigned_to_detail").Find(&bugs)

	json.NewEncoder(w).Encode(bugs)

}

func GetBugPages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// page, _ := strconv.Atoi(request.page)
	// pageSize, _ := strconv.Atoi(request.pageSize)
	// sortBy := request.sortBy

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	sortBy := r.URL.Query().Get("sortBy")

	if sortBy == "" {
		sortBy = "id"
	}
	var bugs []models.Bug
	Db.Order(sortBy).Offset((page - 1) * pageSize).Limit(pageSize).Preload("Qa_contact").Preload("Creator_detail").Preload("Assigned_to_detail").Find(&bugs)

	json.NewEncoder(w).Encode(bugs)
}

func EditBug(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bug := &models.Bug{}
	err := json.NewDecoder(r.Body).Decode(bug)

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	Db.Model(&bug).Where("Id = ?", bug.Id).Updates(&bug)

	json.NewEncoder(w).Encode(bug)
}
