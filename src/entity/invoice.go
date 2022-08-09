package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Invoice struct {
	ID              int             `gorm:"primaryKey"`
	DateOfPayment   time.Time       `gorm:"notNull"`
	PaymentStart    time.Time       `gorm:"notNull"`
	PaymentDue      time.Time       `gorm:"notNull"`
	PupilID         int             `gorm:"notNull"`
	TypeOfServiceID int             `gorm:"notNull"`
	AmountOfMoney   decimal.Decimal `gorm:"type:numeric"`
	Note            *InvoiceNote    `gorm:"foreignKey:ID;references:id"`
}
