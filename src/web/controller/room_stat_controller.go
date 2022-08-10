package controller

import (
	"net/http"

	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
)

func NewRoomStat(log logger.Logger, roomStatServ service.IRoomStatService) IRoomStatController {
	return &roomStatController{log: log, roomStatServ: roomStatServ}
}

type IRoomStatController interface {
	GetRoomType(w http.ResponseWriter, r *http.Request)
	GetAvailableSpace(w http.ResponseWriter, r *http.Request)
	GetAllResidents(w http.ResponseWriter, r *http.Request)
}

type roomStatController struct {
	roomStatServ service.IRoomStatService
	log          logger.Logger
}

func (c *roomStatController) GetRoomType(w http.ResponseWriter, r *http.Request) {
	respondWithErrorLog(c.log, w, http.StatusNotImplemented, notImplemtedMessage)

}
func (c *roomStatController) GetAvailableSpace(w http.ResponseWriter, r *http.Request) {
	respondWithErrorLog(c.log, w, http.StatusNotImplemented, notImplemtedMessage)

}
func (c *roomStatController) GetAllResidents(w http.ResponseWriter, r *http.Request) {
	respondWithErrorLog(c.log, w, http.StatusNotImplemented, notImplemtedMessage)

}
