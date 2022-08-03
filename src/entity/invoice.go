package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Invoice struct {
	ID              uint            `gorm:"primaryKey"`
	DateOfPayment   time.Time       `gorm:"notNull"`
	PupilID         int             `gorm:"notNull"`
	TypeOfServiceID int             `gorm:"notNull"`
	AmountOfMoney   decimal.Decimal `gorm:"type:numeric"`
	Note            InvoiceNote     `gorm:"foreignKey:InvoiceID;references:id"`
}
