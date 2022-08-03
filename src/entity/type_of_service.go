package entity

type TypeOfService struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"notNull"`
}
