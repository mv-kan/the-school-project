package entity

type InvoiceNote struct {
	// ID is the foreign key to the Invoice ID column.
	ID   int `gorm:"primaryKey"`
	Note string
}

func (i *InvoiceNote) GetID() int {
	return i.ID
}

func (i *InvoiceNote) SetID(id int) {
	i.ID = id
}
