package controllers

import (
	"B20_Backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetBugsThroughYear(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userId, _ := strconv.Atoi(r.URL.Query().Get("userId"))
	from, _ := strconv.Atoi(r.URL.Query().Get("from"))
	to, _ := strconv.Atoi(r.URL.Query().Get("to"))
	var bugs []time.Time

	Db.Model(&models.Bug{}).Select("Creation_time").Where("assigned_to_detail_id = ?", userId).Where("extract(year from creation_time) BETWEEN ? AND ?", from, to).Find(&bugs)

	var analytics []int

	for i := from; i <= to; i++ {
		analytics = append(analytics, 0)
	}

	for _, j := range bugs {

		analytics[j.Year()-from]++

	}

	json.NewEncoder(w).Encode(analytics)

}

func GetBugsByYear(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userId, _ := strconv.Atoi(r.URL.Query().Get("userId"))
	year, _ := strconv.Atoi(r.URL.Query().Get("year"))
	var bugs []time.Time

	Db.Model(&models.Bug{}).Select("Creation_time").Where("assigned_to_detail_id = ?", userId).Where("extract(year from creation_time) = ?", year).Find(&bugs)

	var analytics []int

	for i := 1; i <= 12; i++ {
		analytics = append(analytics, 0)
	}

	for _, j := range bugs {

		analytics[j.Month()-1]++

	}

	json.NewEncoder(w).Encode(analytics)

}

func GetBugsByMonth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userId, _ := strconv.Atoi(r.URL.Query().Get("userId"))
	year, _ := strconv.Atoi(r.URL.Query().Get("year"))
	month, _ := strconv.Atoi(r.URL.Query().Get("month"))
	var bugs []time.Time

	Db.Model(&models.Bug{}).Select("Creation_time").Where("assigned_to_detail_id = ?", userId).Where("extract(year from creation_time) = ?", year).Where("extract(month from creation_time) = ?", month).Find(&bugs)

	var analytics []int

	for i := 1; i <= 31; i++ {
		analytics = append(analytics, 0)
	}

	for _, j := range bugs {

		fmt.Println(j.Day())
		analytics[j.Day()-1]++

	}

	json.NewEncoder(w).Encode(analytics)

}

func GetBugsProgress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var bugs []models.Bug
	resultAll := Db.Where("assigned_to_detail_id = ?", id).Find(&bugs)
	resultVerified := Db.Where("assigned_to_detail_id = ?", id).Where("status = ?", "VERIFIED").Find(&bugs)
	resultResolved := Db.Where("assigned_to_detail_id = ?", id).Where("status = ?", "RESOLVED").Find(&bugs)

	progress := map[string]int64{
		"VERIFIED": resultVerified.RowsAffected,
		"RESOLVED": resultResolved.RowsAffected,
		"ALL":      resultAll.RowsAffected,
	}

	json.NewEncoder(w).Encode(progress)

}
