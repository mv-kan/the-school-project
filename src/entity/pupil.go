package entity

type Pupil struct {
	ID            int    `gorm:"primaryKey"`
	Name          string `gorm:"notNull" validate:"required"`
	Surname       string `gorm:"notNull" validate:"required"`
	Email         *string
	Phonenumber   *string
	SchoolClassID int `gorm:"notNull" validate:"required"`
	SupervisorID  *int
	RoomID        *int
	Invoices      []Invoice `gorm:"foreignKey:PupilID" json:"-"`
}

func (p Pupil) GetID() int {
	return p.ID
}

func (p Pupil) SetID(id int) Base {
	p.ID = id
	return p
}
