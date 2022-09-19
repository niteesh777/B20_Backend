package controllers

import (
	"B20_Backend/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func EditProfile(w http.ResponseWriter, r *http.Request) {
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

func GetProfile(w http.ResponseWriter, r *http.Request) {

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
