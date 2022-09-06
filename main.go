package main

import (
	"bugzilla/apis/models"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Contact_Info struct {
	id        int
	real_name string
	name      string
	email     string
	nick      string
}

func intializeRouter() {

	r := mux.NewRouter()
	r.HandleFunc("/bug/{id}", GetBug).Methods("GET")
	//r.HandleFunc("/importData/{start}/{range}", importData).Methods("GET")
	// r.HandleFunc("/bugshistory", GetBugHistory)
	r.HandleFunc("/signup", Register).Methods("POST")
	r.HandleFunc("/login", ValidateLogin).Methods("POST")
	log.Fatal(http.ListenAndServe(":7000", r))
}

func initializeContactInfo(val map[string]interface{}) models.ContactInfo {

	var info = models.ContactInfo{

		Id:        int(val["id"].(float64)),
		Real_name: val["real_name"].(string),
		Name:      val["name"].(string),
		Email:     val["email"].(string),
		Nick:      val["nick"].(string),
	}

	return info

}

func intializeDbConnection() *gorm.DB {

	//dialect := os.Getenv("DIALECT")
	// host := os.Getenv("HOST")
	// dbport := os.Getenv("DBPORT")
	// user := os.Getenv("USER")
	// name := os.Getenv("NAME")
	// password := os.Getenv("PASSWORD")

	//dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, name, password, dbport)

	dbUri := "host=localhost user=postgres password=abc123 dbname=BugZilla port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Succesfully Connected to Data Base", DB)
	}

	DB.AutoMigrate(&models.Bug{})
	DB.AutoMigrate(&models.ContactInfo{})
	DB.AutoMigrate(&models.User{})

	return DB

}

func importData() {
	fmt.Println("Loading data.....")
	for i := 63142; i <= 63152; i++ {

		result, err := gabs.ParseJSON(getbugs(strconv.Itoa(i)))

		if err != nil {
			log.Fatal(err)
		}

		var qinfo models.ContactInfo
		var qck models.ContactInfo
		var x interface{}
		if len(result.Path("bugs").Data().([]interface{})) == 0 {
			continue
		}

		// fmt.Println(reflect.TypeOf(result.Path("bugs.qa_contact_detail").Data()))
		// fmt.Println(reflect.TypeOf(x))
		// fmt.Println(reflect.TypeOf(result.Path("bugs.qa_contact_detail").Data()) != reflect.TypeOf(x))
		// fmt.Println("=========================================")
		if reflect.TypeOf(result.Path("bugs.qa_contact_detail").Data()) != reflect.TypeOf(x) {
			qk := (result.Path("bugs.qa_contact_detail").Data().([]interface{})[0]).(map[string]interface{})
			qck = initializeContactInfo(qk)
			qa := DB.First(&qinfo, qck.Id)
			if qa.RowsAffected == 0 {
				DB.Create(&qck)
			}
		}
		// fmt.Println("=========================================")
		var cck models.ContactInfo
		var cinfo models.ContactInfo
		if reflect.TypeOf(result.Path("bugs.creator_detail").Data()) != reflect.TypeOf(x) {
			ck := (result.Path("bugs.creator_detail").Data().([]interface{})[0]).(map[string]interface{})
			cck = initializeContactInfo(ck)
			cd := DB.First(&cinfo, cck.Id)
			if cd.RowsAffected == 0 {
				DB.Create(&cck)
			}
		}
		// fmt.Println("=========================================")
		var ack models.ContactInfo
		var ainfo models.ContactInfo
		if reflect.TypeOf(result.Path("bugs.assigned_to_detail").Data()) != reflect.TypeOf(x) {
			ak := (result.Path("bugs.assigned_to_detail").Data().([]interface{})[0]).(map[string]interface{})
			ack = initializeContactInfo(ak)
			ad := DB.First(&ainfo, ack.Id)
			if ad.RowsAffected == 0 {
				DB.Create(&ack)
			}
		}
		// fmt.Println("=========================================")
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
			Qa_contact:         qck,
			Creator_detail:     cck,
			Assigned_to_detail: ack,
		}

		var buginfo models.Bug
		bug := DB.First(&buginfo, value.Id)
		if bug.RowsAffected == 0 {
			DB.Create(&value)
		}
		// fmt.Println("=========================================")
		// fmt.Println(value)

		// fmt.Println("=========================================")

		// k = (result.Path("bugs.cc_details").Data().([]interface{})[0].([]Contact_Info))
		// for i := range k {
		// 	ck = initializeContactInfo(i)
		// 	DB.Create(&ck)
		// }

		//res := result.Path("bugs.assigned_to").Data().([]interface{})[0].(string)
		// fmt.Println(res)
		fmt.Println(i, "done")

	}
	fmt.Println("done")
}
func main() {
	intializeDbConnection()
	intializeRouter()
}
