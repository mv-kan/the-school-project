package entity

type Room struct {
	ID         int     `gorm:"primaryKey"`
	RoomNumber string  `gorm:"notNull"`
	RoomTypeID int     `gorm:"notNull"`
	Pupils     []Pupil `gorm:"foreignKey:RoomID"`
}
