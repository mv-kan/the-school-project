package entity

type SchoolClass struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"notNull"`
	ClassYear int     `gorm:"notNull"`
	Pupils    []Pupil `gorm:"foreignKey:SchoolClassID"`
}
