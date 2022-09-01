package entity

type Room struct {
	ID         int     `gorm:"primaryKey"`
	RoomNumber string  `gorm:"notNull" validate:"required"`
	RoomTypeID int     `gorm:"notNull" validate:"required"`
	Pupils     []Pupil `gorm:"foreignKey:RoomID" json:"-"`
}

func (r Room) GetID() int {
	return r.ID
}

func (r Room) SetID(id int) Base {
	r.ID = id
	return r
}
