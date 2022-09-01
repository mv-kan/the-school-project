package entity

type Supervisor struct {
	ID          int     `gorm:"primaryKey"`
	Name        string  `gorm:"notNull" validate:"required"`
	Surname     string  `gorm:"notNull" validate:"required"`
	Email       string  `gorm:"notNull" validate:"required,email"`
	Phonenumber string  `gorm:"notNull" validate:"required,numeric"`
	DormitoryID int     `gorm:"notNull" validate:"required"`
	Pupils      []Pupil `gorm:"foreignKey:SupervisorID" json:"-"`
}

func (s Supervisor) GetID() int {
	return s.ID
}

func (s Supervisor) SetID(id int) Base {
	s.ID = id
	return s
}
