package entity

import (
	"time"
)

type Invoice struct {
	ID              uint        `gorm:"primaryKey"`
	DateOfPayment   time.Time   `gorm:"notNull"`
	PupilID         int         `gorm:"notNull"`
	TypeOfServiceID int         `gorm:"notNull"`
	AmountOfMoney   string      `gorm:"type:numeric"`
	Note            InvoiceNote `gorm:"foreignKey:InvoiceID;references:id"`
}
