package controller

import (
	"net/http"

	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
	"github.com/mv-kan/the-school-project/web/utils"
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

// @Route: method POST /enroll
// @Failure: 404 not found
// @Failure: 400 bad request
// @Failure: 500 internal server error
// @Success: 200 and the entity
func (c *enrollController) Enroll(w http.ResponseWriter, r *http.Request) {
	// get the entity from the request body
	pupil, err := utils.ParseJSONFromBody[entity.Pupil](r.Body)
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusBadRequest, err.Error())
		return
	}

	pupil, err = c.enrollServ.Enroll(pupil)
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, pupil)
}
