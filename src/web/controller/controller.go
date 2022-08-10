package controller

import (
	"net/http"

	"github.com/mv-kan/the-school-project/service"
)

func New[T any](serv service.IService[T]) IController {
	return &Controller[T]{service: serv}
}

type IController interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type Controller[T any] struct {
	service service.IService[T]
}

func (c *Controller[T]) Get(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller[T]) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller[T]) Delete(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller[T]) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller[T]) Update(w http.ResponseWriter, r *http.Request) {

}
