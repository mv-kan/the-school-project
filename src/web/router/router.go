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
		"dorms":            controller.New(dormServ),
		"invoices":         controller.New(invoiceServ),
		"pupils":           controller.New(pupilServ),
		"room-types":       controller.New(roomTypeServ),
		"rooms":            controller.New(roomServ),
		"school-classes":   controller.New(schoolClassServ),
		"supervisors":      controller.New(supervisorServ),
		"type-of-services": controller.New(typeOfServiceServ),
	}
	genericRoutes := make([]route, 0, len(urlpathToController))
	for urlpath, controller := range urlpathToController {
		getRoute := route{Name: "get-" + urlpath, Method: http.MethodGet, Pattern: "/" + urlpath + "/{id:[0-9]+}", HandlerFunc: controller.Get}
		getAllRoute := route{Name: "get-all-" + urlpath, Method: http.MethodGet, Pattern: "/" + urlpath, HandlerFunc: controller.GetAll}
		deleteRoute := route{Name: "delete-" + urlpath, Method: http.MethodDelete, Pattern: "/" + urlpath + "/{id:[0-9]+}", HandlerFunc: controller.Delete}
		createRoute := route{Name: "create-" + urlpath, Method: http.MethodPost, Pattern: "/" + urlpath, HandlerFunc: controller.Create}
		updateRoute := route{Name: "update-" + urlpath, Method: http.MethodPut, Pattern: "/" + urlpath + "/{id:[0-9]+}", HandlerFunc: controller.Update}
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
	// get generic routes
	routes := getGenericRoutes(dormServ, invoiceServ, pupilServ, roomTypeServ, roomServ, schoolClassServ, supervisorServ, typeOfServiceServ)

	// enroll routes
	enrollC := controller.NewEnroll(enrollServ)
	enrollRoutes := []route{
		{Name: "enroll-pupil", Method: http.MethodPost, Pattern: "/enroll", HandlerFunc: enrollC.Enroll},
	}

	// financial routes
	financialC := controller.NewFinancial(financialServ)
	financialRoutes := []route{
		{Name: "get-collected-money-for-month", Method: http.MethodGet, Pattern: "/financial/collected-money-for-month/{year:[0-9]+}/{month:[0-9]+}", HandlerFunc: financialC.CollectedMoneyForMonth},
		{Name: "get-all-lodging-debtors", Method: http.MethodGet, Pattern: "/financial/get-all-lodging-debtors", HandlerFunc: financialC.GetAllLodgingDebtors},
	}

	// room stat routes
	roomStatC := controller.NewRoomStat(roomStatServ)
	roomStatRoutes := []route{
		{Name: "get-all-residents-of-room", Method: http.MethodGet, Pattern: "/rooms/{id:[0-9]+}/all-residents", HandlerFunc: roomStatC.GetAllResidents},
		{Name: "get-available-space-of-room", Method: http.MethodGet, Pattern: "/rooms/{id:[0-9]+}/available-space", HandlerFunc: roomStatC.GetAvailableSpace},
	}

	// merge all routes into one slice
	routes = append(routes, enrollRoutes...)
	routes = append(routes, financialRoutes...)
	routes = append(routes, roomStatRoutes...)

	return routes
}

func New(
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
	roomStatServ service.IRoomStatService) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := getRoutes(dormServ, invoiceServ, pupilServ, roomTypeServ, roomServ, schoolClassServ, supervisorServ, typeOfServiceServ, enrollServ, financialServ, roomStatServ)
	for _, r := range routes {
		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(r.HandlerFunc)
	}
	return router
}
