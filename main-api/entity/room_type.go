package entity

import "github.com/shopspring/decimal"

type RoomType struct {
	ID             int             `gorm:"primaryKey"`
	Price          decimal.Decimal `gorm:"notNull" validate:"required"`
	DormitoryID    int             `gorm:"notNull" validate:"required"`
	Name           string          `gorm:"notNull" validate:"required"`
	MaxOfResidents int             `gorm:"notNull" validate:"required"`
	Rooms          []Room          `gorm:"foreignKey:RoomTypeID" json:"-"`
}

func (r RoomType) GetID() int {
	return r.ID
}

func (r RoomType) SetID(id int) Base {
	r.ID = id
	return r
}
