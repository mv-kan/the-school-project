package entity

import "database/sql"

type Pupil struct {
	ID            int    `gorm:"primaryKey"`
	Name          string `gorm:"notNull" validate:"required"`
	Surname       string `gorm:"notNull" validate:"required"`
	Email         sql.NullString
	Phonenumber   sql.NullString
	SchoolClassID int `gorm:"notNull" validate:"required"`
	SupervisorID  sql.NullInt32
	RoomID        sql.NullInt32
	Invoices      []Invoice `gorm:"foreignKey:PupilID"`
}

func (p Pupil) GetID() int {
	return p.ID
}

func (p Pupil) SetID(id int) Base {
	p.ID = id
	return p
}
