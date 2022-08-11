package entity

type Room struct {
	ID         int     `gorm:"primaryKey"`
	RoomNumber string  `gorm:"notNull"`
	RoomTypeID int     `gorm:"notNull"`
	Pupils     []Pupil `gorm:"foreignKey:RoomID"`
}

func (r Room) GetID() int {
	return r.ID
}

func (r Room) SetID(id int) Base {
	r.ID = id
	return r
}
