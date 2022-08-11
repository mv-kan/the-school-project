package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Invoice struct {
	ID              int             `gorm:"primaryKey"`
	DateOfPayment   time.Time       `gorm:"notNull"`
	PaymentStart    time.Time       `gorm:"notNull" validate:"required"`
	PaymentDue      time.Time       `gorm:"notNull" validate:"required"`
	PupilID         int             `gorm:"notNull" validate:"required"`
	TypeOfServiceID int             `gorm:"notNull" validate:"required"`
	AmountOfMoney   decimal.Decimal `gorm:"type:numeric" validate:"required"`
	Note            *InvoiceNote    `gorm:"foreignKey:ID;references:id"`
}

func (i Invoice) GetID() int {
	return i.ID
}

func (i Invoice) SetID(id int) Base {
	i.ID = id
	return i
}
