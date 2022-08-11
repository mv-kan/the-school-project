package entity

type TypeOfService struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"notNull"`
}

func (t *TypeOfService) GetID() int {
	return t.ID
}
