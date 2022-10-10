package main

import (
	"B20_Backend/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"github.com/rs/cors"
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

func main() {

	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(e)

	port := os.Getenv("DB_PORT")

	// Handle routes
	r := routes.Handlers()
	http.Handle("/", r)

	handler := cors.Default().Handler(r)

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
