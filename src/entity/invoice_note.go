package entity

type InvoiceNote struct {
	// ID is the foreign key to the Invoice ID column.
	ID   int `gorm:"primaryKey"`
	Note string
}
