package entity

type RoomType struct {
	ID             int     `gorm:"primaryKey"`
	Price          float64 `gorm:"notNull"`
	DormitoryID    int     `gorm:"notNull"`
	Name           string  `gorm:"notNull"`
	MaxOfResidents int     `gorm:"notNull"`
	Rooms          []Room  `gorm:"foreignKey:RoomTypeID"`
}
