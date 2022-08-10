package controller

import (
	"net/http"

	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
)

func New[T any](log logger.Logger, serv service.IService[T]) IController {
	return &controller[T]{log: log, service: serv}
}

type IController interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type controller[T any] struct {
	service service.IService[T]
	log     logger.Logger
}

func (c *controller[T]) Get(w http.ResponseWriter, r *http.Request) {
	respondWithErrorLog(c.log, w, http.StatusNotImplemented, notImplemtedMessage)
}

func (c *controller[T]) GetAll(w http.ResponseWriter, r *http.Request) {
	respondWithErrorLog(c.log, w, http.StatusNotImplemented, notImplemtedMessage)
}

func (c *controller[T]) Delete(w http.ResponseWriter, r *http.Request) {
	respondWithErrorLog(c.log, w, http.StatusNotImplemented, notImplemtedMessage)
}

func (c *controller[T]) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *controller[T]) Update(w http.ResponseWriter, r *http.Request) {

}
