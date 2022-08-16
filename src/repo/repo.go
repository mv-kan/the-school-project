package repo

import (
	"errors"

	"github.com/mv-kan/the-school-project/entity"
	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("Record not found")

func New[T entity.Base](db *gorm.DB) IRepository[T] {
	return &Repository[T]{db: db}
}

type IRepository[T entity.Base] interface {
	FindAll() ([]T, error)
	Find(id int) (T, error)
	Delete(id int) error
	Create(T) (T, error)
	CreateWithID(T) (T, error)
	Update(T) error
	WithTx(tx *gorm.DB) IRepository[T]
}

type Repository[T entity.Base] struct {
	db *gorm.DB
}

func (repo Repository[T]) WithTx(tx *gorm.DB) IRepository[T] {
	return New[T](tx)
}

func (repo Repository[T]) FindAll() ([]T, error) {
	var entities []T
	err := repo.db.Find(&entities).Error
	return entities, err
}

func (repo Repository[T]) Find(id int) (T, error) {
	var entity T
	err := repo.db.Take(&entity, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity, ErrRecordNotFound
	} else if err != nil {
		return entity, err
	}
	return entity, nil
}

func (repo Repository[T]) Delete(id int) error {
	var entity T
	result := repo.db.Delete(&entity, id)
	err := result.Error
	if err != nil {
		return err
	} else if result.RowsAffected < 1 {
		return ErrRecordNotFound
	}
	return err
}

func (repo Repository[T]) Create(entity T) (T, error) {
	entity = entity.SetID(0).(T)
	err := repo.db.Create(&entity).Error
	return entity, err
}

func (repo Repository[T]) CreateWithID(entity T) (T, error) {
	err := repo.db.Create(&entity).Error
	return entity, err
}

func (repo Repository[T]) Update(entity T) error {
	// check if entity exists
	_, err := repo.Find(entity.GetID())
	if err != nil {
		return err
	}
	// and then update
	err = repo.db.Model(&entity).Updates(entity).Error
	return err
}
