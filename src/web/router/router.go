package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/service"
	"github.com/mv-kan/the-school-project/web/controller"
)

func NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not Implemented"))
}

func getGenericRoutes(
	dormServ service.IService[entity.Dormitory],
	invoiceServ service.IService[entity.Invoice],
	pupilServ service.IService[entity.Pupil],
	roomTypeServ service.IService[entity.RoomType],
	roomServ service.IService[entity.Room],
	schoolClassServ service.IService[entity.SchoolClass],
	supervisorServ service.IService[entity.Supervisor],
	typeOfServiceServ service.IService[entity.TypeOfService],
) []route {
	urlpathToController := map[string]controller.IController{
		"dorms":            controller.New[entity.Dormitory](dormServ),
		"invoices":         controller.New[entity.Invoice](invoiceServ),
		"pupils":           controller.New[entity.Pupil](pupilServ),
		"room-types":       controller.New[entity.RoomType](roomTypeServ),
		"rooms":            controller.New[entity.Room](roomServ),
		"school-classes":   controller.New[entity.SchoolClass](schoolClassServ),
		"supervisors":      controller.New[entity.Supervisor](supervisorServ),
		"type-of-services": controller.New[entity.TypeOfService](typeOfServiceServ),
	}
	genericRoutes := make([]route, 0, len(urlpathToController))
	for k, v := range urlpathToController {
		getRoute := route{Name: "get-" + k, Method: http.MethodGet, Pattern: "/" + k + "/{id:[0-9]+}", HandlerFunc: v.Get}
		getAllRoute := route{Name: "get-all-" + k, Method: http.MethodGet, Pattern: "/" + k, HandlerFunc: v.GetAll}
		deleteRoute := route{Name: "delete-" + k, Method: http.MethodDelete, Pattern: "/" + k + "/{id:[0-9]+}", HandlerFunc: v.Delete}
		createRoute := route{Name: "create-" + k, Method: http.MethodPost, Pattern: "/" + k, HandlerFunc: v.Create}
		updateRoute := route{Name: "update-" + k, Method: http.MethodPut, Pattern: "/" + k + "/{id:[0-9]+}", HandlerFunc: v.Update}
		genericRoutes = append(genericRoutes, getRoute, getAllRoute, deleteRoute, createRoute, updateRoute)
	}
	return genericRoutes
}

func getRoutes(
	dormServ service.IService[entity.Dormitory],
	invoiceServ service.IService[entity.Invoice],
	pupilServ service.IService[entity.Pupil],
	roomTypeServ service.IService[entity.RoomType],
	roomServ service.IService[entity.Room],
	schoolClassServ service.IService[entity.SchoolClass],
	supervisorServ service.IService[entity.Supervisor],
	typeOfServiceServ service.IService[entity.TypeOfService],
	enrollServ service.IEnrollService,
	financialServ service.IFinancialService,
	roomStatServ service.IRoomStatService) []route {
	routes := getGenericRoutes(dormServ, invoiceServ, pupilServ, roomTypeServ, roomServ, schoolClassServ, supervisorServ, typeOfServiceServ)
	return routes
}

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := getRoutes()
	for _, r := range routes {
		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(r.HandlerFunc)
	}
	return router
}
