package entity

type SchoolClass struct {
	ID        int     `gorm:"primaryKey"`
	Name      string  `gorm:"notNull"`
	ClassYear int     `gorm:"notNull"`
	Pupils    []Pupil `gorm:"foreignKey:SchoolClassID"`
}

func (sc *SchoolClass) GetID() int {
	return sc.ID
}
