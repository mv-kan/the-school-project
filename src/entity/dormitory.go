package entity

type Dormitory struct {
	ID          int          `gorm:"primaryKey"`
	Name        string       `gorm:"notNull"`
	RoomTypes   []RoomType   `gorm:"foreignKey:DormitoryID"`
	Supervisors []Supervisor `gorm:"foreignKey:DormitoryID"`
}
