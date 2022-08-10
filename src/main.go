package main

import (
	"net/http"

	"github.com/mv-kan/the-school-project/web/router"
)

func main() {

	muxRouter := router.New()
	http.Handle("/", muxRouter)
	http.ListenAndServe(":8080", nil)
}
