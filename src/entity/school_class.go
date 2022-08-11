package entity

type SchoolClass struct {
	ID        int     `gorm:"primaryKey"`
	Name      string  `gorm:"notNull" validate:"required"`
	ClassYear int     `gorm:"notNull" validate:"required"`
	Pupils    []Pupil `gorm:"foreignKey:SchoolClassID"`
}

func (s SchoolClass) GetID() int {
	return s.ID
}

func (s SchoolClass) SetID(id int) Base {
	s.ID = id
	return s
}
