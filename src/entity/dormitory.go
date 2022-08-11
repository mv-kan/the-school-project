package entity

type Dormitory struct {
	ID          int          `gorm:"primaryKey"`
	Name        string       `gorm:"notNull"`
	RoomTypes   []RoomType   `gorm:"foreignKey:DormitoryID"`
	Supervisors []Supervisor `gorm:"foreignKey:DormitoryID"`
}

func (d *Dormitory) GetID() int {
	return d.ID
}

func (d *Dormitory) SetID(id int) {
	d.ID = id
}
