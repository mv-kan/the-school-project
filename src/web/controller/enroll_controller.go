package controller

import (
	"net/http"

	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
)

func NewEnroll(log logger.Logger, enrollServ service.IEnrollService) IEnrollController {
	return &enrollController{log: log, enrollServ: enrollServ}
}

type IEnrollController interface {
	Enroll(w http.ResponseWriter, r *http.Request)
}

type enrollController struct {
	enrollServ service.IEnrollService
	log        logger.Logger
}

func (c *enrollController) Enroll(w http.ResponseWriter, r *http.Request) {

}
