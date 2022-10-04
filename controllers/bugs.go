package controllers

import (
	"B20_Backend/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

var Url = "https://bugzilla.mozilla.org/rest/"

func GetBug(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
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

func GetBugLocalFormat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	i := params["id"]
	result, err := gabs.ParseJSON(getbugs(i))

	if err != nil {
		log.Fatal(err)
	}

	if len(result.Path("bugs").Data().([]interface{})) == 0 {
		json.NewEncoder(w).Encode("No bug with that id")
		return
	}

	qa_contact_detail := checkData("bugs.qa_contact_detail", result)
	creator_detail := checkData("bugs.creator_detail", result)
	assigned_to_detail := checkData("bugs.assigned_to_detail", result)

	last_change_time := result.Path("bugs.last_change_time").Data().([]interface{})[0].(string)
	creation_time := result.Path("bugs.creation_time").Data().([]interface{})[0].(string)

	ltime, err := time.Parse(time.RFC3339, last_change_time)
	ctime, err := time.Parse(time.RFC3339, creation_time)

	var value = models.Bug{
		Id:            result.Path("bugs.id").Data().([]interface{})[0].(float64),
		Comment_count: result.Path("bugs.comment_count").Data().([]interface{})[0].(float64),
		// Deadline:         result.Path("bugs.deadline").Data().([]interface{})[0].(string),
		Type:               result.Path("bugs.type").Data().([]interface{})[0].(string),
		Status:             result.Path("bugs.status").Data().([]interface{})[0].(string),
		Priority:           result.Path("bugs.priority").Data().([]interface{})[0].(string),
		Severity:           result.Path("bugs.severity").Data().([]interface{})[0].(string),
		Summary:            result.Path("bugs.summary").Data().([]interface{})[0].(string),
		Product:            result.Path("bugs.product").Data().([]interface{})[0].(string),
		Platform:           result.Path("bugs.platform").Data().([]interface{})[0].(string),
		Resolution:         result.Path("bugs.resolution").Data().([]interface{})[0].(string),
		Target_milestone:   result.Path("bugs.target_milestone").Data().([]interface{})[0].(string),
		Classification:     result.Path("bugs.classification").Data().([]interface{})[0].(string),
		Is_confirmed:       result.Path("bugs.is_confirmed").Data().([]interface{})[0].(bool),
		Is_open:            result.Path("bugs.is_open").Data().([]interface{})[0].(bool),
		Last_change_time:   ltime,
		Creation_time:      ctime,
		Qa_contact:         qa_contact_detail,
		Creator_detail:     creator_detail,
		Assigned_to_detail: assigned_to_detail,
	}

	json.NewEncoder(w).Encode(value)
}

func GetBugInfoLocal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	params := mux.Vars(r)
	id := params["id"]

	var bug models.Bug
	Db.Where("id = ?", id).Preload("Qa_contact").Preload("Creator_detail").Preload("Assigned_to_detail").Find(&bug)

	json.NewEncoder(w).Encode(bug)

}

func GetAssignedBugs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	params := mux.Vars(r)
	userId := params["userId"]

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	sortBy := r.URL.Query().Get("sortBy")

	if sortBy == "" {
		sortBy = "id"
	}

	var bugs []models.Bug
	Db.Where("assigned_to_detail_id = ?", userId).Order(sortBy).Offset((page - 1) * pageSize).Limit(pageSize).Preload("Qa_contact").Preload("Creator_detail").Preload("Assigned_to_detail").Find(&bugs)

	json.NewEncoder(w).Encode(bugs)

}

func GetCreatedBug(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	params := mux.Vars(r)
	userId := params["userId"]

	var bugs []models.Bug
	Db.Where("creator_detail_id = ?", userId).Preload("Qa_contact").Preload("Creator_detail").Preload("Assigned_to_detail").Find(&bugs)

	json.NewEncoder(w).Encode(bugs)

}

func GetRelatedBug(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	params := mux.Vars(r)
	userId := params["userId"]

	var bugs []models.Bug
	Db.Where("qa_contact_id = ?", userId).Preload("Qa_contact").Preload("Creator_detail").Preload("Assigned_to_detail").Find(&bugs)

	json.NewEncoder(w).Encode(bugs)

}

func GetAllBugs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	params := mux.Vars(r)
	userId := params["userId"]

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	sortBy := r.URL.Query().Get("sortBy")

	if sortBy == "" {
		sortBy = "id"
	}

	var bugs []models.Bug
	Db.Where("assigned_to_detail_id = ?", userId).Order(sortBy).Offset((page-1)*pageSize).Limit(pageSize).Or("creator_detail_id = ?", userId).Or("qa_contact_id = ?", userId).Preload("Qa_contact").Preload("Creator_detail").Preload("Assigned_to_detail").Find(&bugs)

	json.NewEncoder(w).Encode(bugs)

}

func GetBugPages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
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
	enableCors(&w)

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
