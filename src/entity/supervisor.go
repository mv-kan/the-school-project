package entity

type Supervisor struct {
	ID          int     `gorm:"primaryKey"`
	Name        string  `gorm:"notNull"`
	Surname     string  `gorm:"notNull"`
	Email       string  `gorm:"notNull"`
	Phonenumber string  `gorm:"notNull"`
	DormitoryID int     `gorm:"notNull"`
	Pupils      []Pupil `gorm:"foreignKey:SupervisorID"`
}

func (s Supervisor) GetID() int {
	return s.ID
}

func (s Supervisor) SetID(id int) Base {
	s.ID = id
	return s
}
