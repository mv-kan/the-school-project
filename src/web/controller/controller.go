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
		return
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
		return
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
// @Failure: 400 bad request
// @Failure: 500 internal server error
// @Success: 201 and the entity
func (c *controller[T]) Create(w http.ResponseWriter, r *http.Request) {
	// get the entity from the request body
	entity, err := utils.ParseJSONFromBody[T](r.Body)
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusBadRequest, err.Error())
		return
	}

	// create the entity
	entityWithID, err := c.service.Create(entity)
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, entityWithID)
}

// @Route: method PUT /entities/{ID}
// @Failure: 404 not found
// @Failure: 400 bad request
// @Failure: 500 internal server error
// @Success: http.StatusOK and the entity
func (c *controller[T]) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusBadRequest, err.Error())
		return
	}
	// get the entity from the request body
	entity, err := utils.ParseJSONFromBody[T](r.Body)
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusBadRequest, err.Error())
		return
	}
	// set entity id to the id from the url
	entity = entity.SetID(id).(T)
	// create the entity
	err = c.service.Update(entity)
	if err != nil {
		utils.RespondWithErrorLog(c.log, w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, entity)
}
