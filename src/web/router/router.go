package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
	"github.com/mv-kan/the-school-project/web/controller"
	"github.com/mv-kan/the-school-project/web/utils"
)

func getGenericRoutes(
	log logger.Logger,
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
		"dorms":             controller.New(log, dormServ),
		"invoices":          controller.New(log, invoiceServ),
		"pupils":            controller.New(log, pupilServ),
		"room-types":        controller.New(log, roomTypeServ),
		"rooms":             controller.New(log, roomServ),
		"school-classes":    controller.New(log, schoolClassServ),
		"supervisors":       controller.New(log, supervisorServ),
		"types-of-services": controller.New(log, typeOfServiceServ),
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
	log logger.Logger,
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
	routes := getGenericRoutes(log, dormServ, invoiceServ, pupilServ, roomTypeServ, roomServ, schoolClassServ, supervisorServ, typeOfServiceServ)

	// enroll routes
	enrollC := controller.NewEnroll(log, enrollServ)
	enrollRoutes := []route{
		{Name: "enroll-pupil", Method: http.MethodPost, Pattern: "/enroll", HandlerFunc: enrollC.Enroll},
	}

	// financial routes
	financialC := controller.NewFinancial(log, financialServ)
	financialRoutes := []route{
		{Name: "get-collected-money-for-month", Method: http.MethodGet, Pattern: "/financial/collected-money-for-month/{year:[0-9]+}/{month:[0-9]+}", HandlerFunc: financialC.CollectedMoneyForMonth},
		{Name: "get-all-lodging-debtors", Method: http.MethodGet, Pattern: "/financial/get-all-lodging-debtors", HandlerFunc: financialC.GetAllLodgingDebtors},
	}

	// room stat routes
	roomC := controller.NewRoom(log, roomServ, roomStatServ)
	roomStatRoutes := []route{
		{Name: "get-type-of-room", Method: http.MethodGet, Pattern: "/rooms/{id:[0-9]+}/type", HandlerFunc: roomC.GetRoomType},
		{Name: "get-all-residents-of-room", Method: http.MethodGet, Pattern: "/rooms/{id:[0-9]+}/all-residents", HandlerFunc: roomC.GetAllResidents},
		{Name: "get-available-space-of-room", Method: http.MethodGet, Pattern: "/rooms/{id:[0-9]+}/available-space", HandlerFunc: roomC.GetAvailableSpace},
	}

	// merge all routes into one slice
	routes = append(routes, enrollRoutes...)
	routes = append(routes, financialRoutes...)
	routes = append(routes, roomStatRoutes...)

	return routes
}

func New(
	log logger.Logger,
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
	routes := getRoutes(log, dormServ, invoiceServ, pupilServ, roomTypeServ, roomServ, schoolClassServ, supervisorServ, typeOfServiceServ, enrollServ, financialServ, roomStatServ)
	for _, r := range routes {
		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(r.HandlerFunc)
	}

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithError(w, http.StatusNotFound, "Endpoint is not found!")
	})
	return router
}
