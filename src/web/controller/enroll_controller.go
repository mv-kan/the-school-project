package controller

import (
	"net/http"

	"github.com/mv-kan/the-school-project/service"
)

func NewEnroll(enrollServ service.IEnrollService) IEnrollController {
	return &enrollController{enrollServ: enrollServ}
}

type IEnrollController interface {
	Enroll(w http.ResponseWriter, r *http.Request)
}

type enrollController struct {
	enrollServ service.IEnrollService
}

func (c *enrollController) Enroll(w http.ResponseWriter, r *http.Request) {

}
