package entity

type InvoiceNote struct {
	InvoiceID int `gorm:"primaryKey"`
	Note      string
}
