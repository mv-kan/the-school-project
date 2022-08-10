package controller

import (
	"net/http"

	"github.com/mv-kan/the-school-project/service"
)

func NewRoomStat(roomStatServ service.IRoomStatService) IRoomStatController {
	return &roomStatController{roomStatServ: roomStatServ}
}

type IRoomStatController interface {
	GetRoomType(w http.ResponseWriter, r *http.Request)
	GetAvailableSpace(w http.ResponseWriter, r *http.Request)
	GetAllResidents(w http.ResponseWriter, r *http.Request)
}

type roomStatController struct {
	roomStatServ service.IRoomStatService
}

func (c *roomStatController) GetRoomType(w http.ResponseWriter, r *http.Request) {

}
func (c *roomStatController) GetAvailableSpace(w http.ResponseWriter, r *http.Request) {

}
func (c *roomStatController) GetAllResidents(w http.ResponseWriter, r *http.Request) {

}
