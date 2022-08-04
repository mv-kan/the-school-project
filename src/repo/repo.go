package repo

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("Record not found")

func New[T any](db *gorm.DB) IRepository[T] {
	return &Repository[T]{db: db}
}

type IRepository[T any] interface {
	FindAll() ([]T, error)
	Find(id int) (T, error)
	Delete(id int) error
	Save(T) (T, error)
	Update(T) error
	WithTx(tx *gorm.DB) IRepository[T]
}

type Repository[T any] struct {
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
	err := repo.db.Find(&entity, id).Error
	if errors.Is(err, sql.ErrNoRows) {
		return entity, ErrRecordNotFound
	} else if err != nil {
		return entity, err
	}
	return entity, nil
}

func (repo Repository[T]) Delete(id int) error {
	var entity T
	err := repo.db.Delete(&entity, id).Error
	if errors.Is(err, sql.ErrNoRows) {
		return ErrRecordNotFound
	}
	return err
}

func (repo Repository[T]) Save(entity T) (T, error) {
	err := repo.db.Create(&entity).Error
	return entity, err
}

func (repo Repository[T]) Update(entity T) error {
	err := repo.db.Save(&entity).Error
	return err
}
