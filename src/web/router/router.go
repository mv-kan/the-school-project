package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := []route{
		{Name: "Index", Method: http.MethodGet, Pattern: "/", HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		}},
	}
	for _, r := range routes {
		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(r.HandlerFunc)
	}
	return router
}
