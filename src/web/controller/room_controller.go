package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/repo"
	"github.com/mv-kan/the-school-project/service"
	"github.com/mv-kan/the-school-project/web/utils"
)

func NewRoom(log logger.Logger, roomServ service.IService[entity.Room], roomStatServ service.IRoomStatService) IRoomController {
	return &roomController{controller: controller[entity.Room]{log: log, service: roomServ}, roomStatServ: roomStatServ}
}

type IRoomController interface {
	IController
	GetRoomType(w http.ResponseWriter, r *http.Request)
	GetAvailableSpace(w http.ResponseWriter, r *http.Request)
	GetAllResidents(w http.ResponseWriter, r *http.Request)
}

type roomController struct {
	controller[entity.Room]
	roomStatServ service.IRoomStatService
}

// @Route: method GET /rooms/{id}/type
// @Failure: 404 not found
// @Failure: 400 bad request
// @Failure: 500 internal server error
// @Success: http.StatusOK and the room type entity
func (c *roomController) GetRoomType(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusBadRequest, err.Error())
		return
	}
	roomType, err := c.roomStatServ.FindRoomType(id)
	if errors.Is(err, repo.ErrRecordNotFound) {
		utils.RespondWithErrorLog(c.log, w, http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, roomType)
}

// @Route: method GET /rooms/{id}/available-space
// @Failure: 404 not found
// @Failure: 400 bad request
// @Failure: 500 internal server error
// @Success: http.StatusOK and available space in the room
func (c *roomController) GetAvailableSpace(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusBadRequest, err.Error())
		return
	}
	availableSpace, err := c.roomStatServ.FindAvailableSpace(id)
	if errors.Is(err, repo.ErrRecordNotFound) {
		utils.RespondWithErrorLog(c.log, w, http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]int{"availableSpace": availableSpace})
}

// @Route: method GET /rooms/{id}/all-residents
// @Failure: 404 not found
// @Failure: 400 bad request
// @Failure: 500 internal server error
// @Success: http.StatusOK and all residents(entity.Pupil) in the room
func (c *roomController) GetAllResidents(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusBadRequest, err.Error())
		return
	}
	residents, err := c.roomStatServ.FindAllResidents(id)
	if errors.Is(err, repo.ErrRecordNotFound) {
		utils.RespondWithErrorLog(c.log, w, http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, residents)
}
