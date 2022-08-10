package controller

import (
	"net/http"

	"github.com/mv-kan/the-school-project/service"
)

func New[T any](serv service.IService[T]) IController {
	return &controller[T]{service: serv}
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
}

func (c *controller[T]) Get(w http.ResponseWriter, r *http.Request) {

}

func (c *controller[T]) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (c *controller[T]) Delete(w http.ResponseWriter, r *http.Request) {

}

func (c *controller[T]) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *controller[T]) Update(w http.ResponseWriter, r *http.Request) {

}
