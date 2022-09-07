package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
