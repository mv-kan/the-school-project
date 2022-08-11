package entity

type SchoolClass struct {
	ID        int     `gorm:"primaryKey"`
	Name      string  `gorm:"notNull"`
	ClassYear int     `gorm:"notNull"`
	Pupils    []Pupil `gorm:"foreignKey:SchoolClassID"`
}

func (s SchoolClass) GetID() int {
	return s.ID
}

func (s SchoolClass) SetID(id int) Base {
	s.ID = id
	return s
}
