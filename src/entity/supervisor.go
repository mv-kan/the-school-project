package entity

type Supervisor struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"notNull"`
	Surname     string  `gorm:"notNull"`
	Email       string  `gorm:"notNull"`
	Phonenumber string  `gorm:"notNull"`
	DormitoryID int     `gorm:"notNull"`
	Pupils      []Pupil `gorm:"foreignKey:SupervisorID"`
}
