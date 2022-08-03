package entity

type Room struct {
	ID         uint    `gorm:"primaryKey"`
	RoomNumber string  `gorm:"notNull"`
	RoomTypeID int     `gorm:"notNull"`
	Pupils     []Pupil `gorm:"foreignKey:RoomID"`
}
