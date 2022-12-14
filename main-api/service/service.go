package service

import (
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
)

const (
	SERVICE_TYPE_LODGING       = 1
	SERVICE_TYPE_EDUCATION     = 1
	END_OF_SCHOOL_YEAR_DAY     = 27
	END_OF_SCHOOL_YEAR_MONTH   = 6
	START_OF_SCHOOL_YEAR_DAY   = 1
	START_OF_SCHOOL_YEAR_MONTH = 9
	CHARGE_FOR_EDUCATION       = 500
)

// func New[T entity.Base](db *gorm.DB) IService[T] {
// 	var instance T
// 	switch instance.(type) {
// 	case entity.Invoice:
// 		return NewInvoice()
// 	default:
// 		return &service[T]{r: r}
// 	}
// }

func New[T entity.Base](r repo.IRepository[T]) IService[T] {
	return &service[T]{r: r}
}

type IService[T entity.Base] interface {
	FindAll() ([]T, error)
	Find(id int) (T, error)
	Delete(id int) error
	Create(T) (T, error)
	Update(T) error
}

type service[T entity.Base] struct {
	r repo.IRepository[T]
}

// err not found is in repo package
func (s service[T]) FindAll() ([]T, error) {
	return s.r.FindAll()
}

// err not found is in repo package
func (s service[T]) Find(id int) (T, error) {
	return s.r.Find(id)
}
func (s service[T]) Delete(id int) error {
	return s.r.Delete(id)
}
func (s service[T]) Create(entity T) (T, error) {
	return s.r.Create(entity)
}
func (s service[T]) Update(entity T) error {
	return s.r.Update(entity)
}
