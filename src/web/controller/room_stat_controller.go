package controller

import (
	"net/http"

	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
	"github.com/mv-kan/the-school-project/web/utils"
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
	utils.RespondWithErrorLog(c.log, w, http.StatusNotImplemented, utils.NotImplemtedMessage)

}
func (c *roomStatController) GetAvailableSpace(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithErrorLog(c.log, w, http.StatusNotImplemented, utils.NotImplemtedMessage)

}
func (c *roomStatController) GetAllResidents(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithErrorLog(c.log, w, http.StatusNotImplemented, utils.NotImplemtedMessage)

}
