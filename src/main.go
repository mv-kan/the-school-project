package main

import (
	"net/http"

	"github.com/mv-kan/the-school-project/web/router"
	"gorm.io/gorm"
)

func connectToDB() (*gorm.DB, error) {

}

func main() {
	// TODO add connection to db
	// TODO New repoisotries
	// TODO New services
	// TODO New routers

	// TODO add logger +
	// TODO add documentation for all internal tools
	// TODO add documentation for all routes in swagger

	db, err := connectToDB()

	muxRouter := router.New()
	http.Handle("/", muxRouter)
	http.ListenAndServe(":8080", nil)
}
