package routes

import (
	"B20_Backend/controllers"
	"B20_Backend/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)
	// r.HandleFunc("/bug/{id}", controllers.GetBug).Methods("GET")
	r.HandleFunc("/importData/{start}/{range}", controllers.ImportData).Methods("GET")
	r.HandleFunc("/importLogin", controllers.ImportintoLogin).Methods("GET")
	r.HandleFunc("/bugPages", controllers.GetBugPages).Methods("GET")
	r.HandleFunc("/signup", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.ValidateLogin).Methods("POST")
	r.HandleFunc("/auth/bug/assigned/{userId}", controllers.GetAssignedBugs).Methods("GET")
	r.HandleFunc("/auth/bug/BugInfo/{id}", controllers.GetBugInfoLocal).Methods("GET")
	r.HandleFunc("/bug/{id}", controllers.GetBugLocalFormat).Methods("GET")
	r.HandleFunc("/auth/ProfileInfo/{id}", controllers.GetProfileInfoLocal).Methods("GET")
	r.HandleFunc("/auth/editprofile", controllers.GetProfileInfoLocal).Methods("PUT")
	// r.HandleFunc("/auth/editprofile/", controllers.EditProfile).Methods("PUT")

	//analytics
	r.HandleFunc("/auth/bugsProgress/{id}", controllers.GetBugsProgress).Methods("GET")
	r.HandleFunc("/auth/filterByYear", controllers.GetBugsThroughYear).Methods("GET")
	r.HandleFunc("/auth/filterByMonth", controllers.GetBugsByYear).Methods("GET")
	r.HandleFunc("/auth/filterByDays", controllers.GetBugsByMonth).Methods("GET")
	r.HandleFunc("/filterByDate", controllers.GetBugsByDate).Methods("GET")



	//preferences
	r.HandleFunc("/auth/createPreference", controllers.CreatePreferenceUsers).Methods("GET")
	r.HandleFunc("/auth/getPreference/{id}", controllers.GetPreference).Methods("GET")
	r.HandleFunc("/auth/editPreference/{id}", controllers.EditPreference).Methods("PUT")



	s := r.PathPrefix("/auth").Subrouter()
	s.Use(utils.JwtVerify)
	s.Use(CommonMiddleware)
	// s.HandleFunc("/bug/assigned/{userId}", controllers.GetAssignedBugs).Methods("GET")
	s.HandleFunc("/bug/created/{userId}", controllers.GetCreatedBug).Methods("GET")
	s.HandleFunc("/bug/qaRelated/{userId}", controllers.GetRelatedBug).Methods("GET")
	s.HandleFunc("/bug/all/{userId}", controllers.GetAllBugs).Methods("GET")
	// s.HandleFunc("/editprofile/", controllers.EditProfile).Methods("PUT")
	s.HandleFunc("/bug/editBug/", controllers.EditBug).Methods("PUT")
	// s.HandleFunc("/bug/BugInfo/{id}", controllers.GetBugInfoLocal).Methods("GET")
	// s.HandleFunc("/ProfileInfo/{id}", controllers.GetProfileInfoLocal).Methods("GET")
	//analytics
	// s.HandleFunc("/filterByYear", controllers.GetBugsThroughYear).Methods("GET")
	// s.HandleFunc("/filterByMonth", controllers.GetBugsByYear).Methods("GET")
	// s.HandleFunc("/filterByDays", controllers.GetBugsByMonth).Methods("GET")
	// r.HandleFunc("/bugshistory", GetBugHistory)

	// log.Fatal(http.ListenAndServe(":7000", r))

	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
