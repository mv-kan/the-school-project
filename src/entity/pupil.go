package entity

import "database/sql"

type Pupil struct {
	ID            int    `gorm:"primaryKey"`
	Name          string `gorm:"notNull"`
	Surname       string `gorm:"notNull"`
	Email         sql.NullString
	Phonenumber   sql.NullString
	SchoolClassID int `gorm:"notNull"`
	SupervisorID  sql.NullInt32
	RoomID        sql.NullInt32
	Invoices      []Invoice `gorm:"foreignKey:PupilID"`
}

func (p *Pupil) GetID() int {
	return p.ID
}

func (p *Pupil) SetID(id int) {
	p.ID = id
}
