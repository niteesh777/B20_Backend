package routes

import (
	"B20_Backend/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/bug/{id}", controllers.GetBug).Methods("GET")
	r.HandleFunc("/importData/{start}/{range}", controllers.ImportData).Methods("GET")
	r.HandleFunc("/importLogin", controllers.ImportintoLogin).Methods("GET")
	r.HandleFunc("/bug/assigned/{userId}", controllers.GetAssignedBugs).Methods("GET")
	r.HandleFunc("/bug/created/{userId}", controllers.GetCreatedBug).Methods("GET")
	r.HandleFunc("/bug/qaRelated/{userId}", controllers.GetRelatedBug).Methods("GET")
	r.HandleFunc("/bug/all/{userId}", controllers.GetAllBugs).Methods("GET")
	r.HandleFunc("/bugPages", controllers.GetBugPages).Methods("GET")
	// r.HandleFunc("/bugshistory", GetBugHistory)
	r.HandleFunc("/signup", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.ValidateLogin).Methods("POST")
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
