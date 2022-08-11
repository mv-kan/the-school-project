package entity

type Dormitory struct {
	ID          int          `gorm:"primaryKey"`
	Name        string       `gorm:"notNull" validate:"required"`
	RoomTypes   []RoomType   `gorm:"foreignKey:DormitoryID"`
	Supervisors []Supervisor `gorm:"foreignKey:DormitoryID"`
}

func (d Dormitory) GetID() int {
	return d.ID
}

func (d Dormitory) SetID(id int) Base {
	d.ID = id
	return d
}
