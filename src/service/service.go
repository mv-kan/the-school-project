package service

import "github.com/mv-kan/the-school-project/repo"

func New[T any](r repo.IRepository[T]) IService[T] {
	return &service[T]{r: r}
}

type IService[T any] interface {
	FindAll() ([]T, error)
	Find(id int) (T, error)
	Delete(id int) error
	Save(T) (T, error)
	Update(T) error
}

type service[T any] struct {
	r repo.IRepository[T]
}

func (s service[T]) FindAll() ([]T, error) {
	return s.r.FindAll()
}
func (s service[T]) Find(id int) (T, error) {
	return s.r.Find(id)
}
func (s service[T]) Delete(id int) error {
	return s.r.Delete(id)
}
func (s service[T]) Save(entity T) (T, error) {
	return s.r.Save(entity)
}
func (s service[T]) Update(entity T) error {
	return s.r.Update(entity)
}
