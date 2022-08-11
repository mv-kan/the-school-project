package entity

import "github.com/shopspring/decimal"

type RoomType struct {
	ID             int             `gorm:"primaryKey"`
	Price          decimal.Decimal `gorm:"notNull"`
	DormitoryID    int             `gorm:"notNull"`
	Name           string          `gorm:"notNull"`
	MaxOfResidents int             `gorm:"notNull"`
	Rooms          []Room          `gorm:"foreignKey:RoomTypeID"`
}

func (rt *RoomType) GetID() int {
	return rt.ID
}

func (rt *RoomType) SetID(id int) {
	rt.ID = id
}
