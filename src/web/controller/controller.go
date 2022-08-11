package controller

import (
	"encoding/json"
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

func New[T entity.Base](log logger.Logger, serv service.IService[T]) IController {
	return &controller[T]{log: log, service: serv}
}

type IController interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type controller[T entity.Base] struct {
	service service.IService[T]
	log     logger.Logger
}

// @Route: method GET /entities/{id}
// @Failure: 404 not found
// @Failure: 400 bad request
// @Failure: 500 internal server error
// @Success: http.StatusOK and the entity
func (c *controller[T]) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusBadRequest, err.Error())
	}
	entity, err := c.service.Find(id)
	if errors.Is(err, repo.ErrRecordNotFound) {
		utils.RespondWithErrorLog(c.log, w, http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, entity)
}

// @Route: method GET /entities
// @Failure: 404 not found
// @Failure: 500
// @Success: http.StatusOK and the entities
func (c *controller[T]) GetAll(w http.ResponseWriter, r *http.Request) {
	entities, err := c.service.FindAll()
	if errors.Is(err, repo.ErrRecordNotFound) {
		utils.RespondWithErrorLog(c.log, w, http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, entities)
}

// @Route: method DELETE /entities/{id}
// @Failure: 404 not found
// @Failure: 400 bad request
// @Failure: 500 internal server error
// @Success: http.StatusOK
func (c *controller[T]) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusBadRequest, err.Error())
	}
	err = c.service.Delete(id)
	if errors.Is(err, repo.ErrRecordNotFound) {
		utils.RespondWithErrorLog(c.log, w, http.StatusNotFound, err.Error())
		return
	} else if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, nil)
}

// @Route: method POST /entities
// @Failure: 404 not found
// @Failure: 500
// @Success: http.StatusOK and the entity
func (c *controller[T]) Create(w http.ResponseWriter, r *http.Request) {
	// get the entity from the request body
	entity := &T{}
	err := json.NewDecoder(r.Body, entity)
}

// @Route: method GET /entities
// @Failure: 404 not found
// @Failure: 500
// @Success: http.StatusOK and the entity
func (c *controller[T]) Update(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithErrorLog(c.log, w, http.StatusNotImplemented, utils.NotImplemtedMessage)
}
