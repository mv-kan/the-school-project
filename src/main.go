package main

import (
	"fmt"

	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	const (
		DB_USER     = "postgres"
		DB_PASSWORD = "secret"
		DB_HOST     = "localhost"
		DB_PORT     = 5432
		DB_NAME     = "postgres"
	)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	// sqlDB, err := sql.Open("postgres", dsn)
	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	pupil := entity.Pupil{}

	db.Preload("Invoices").First(&pupil, 2)
	for i, invoice := range pupil.Invoices {
		db.Model(&invoice).Preload("Note").Find(&invoice)
		pupil.Invoices[i] = invoice
	}

	fmt.Println(pupil)
	pupilRepo := repo.New[entity.Pupil](db)
	pupil, err = pupilRepo.Find(3)
	if err != nil {
		panic(err)
	}
	// find invoices
	invoiceRepo := repo.New[entity.Invoice](db)
	invoices, err := invoiceRepo.FindAll()
	if err != nil {
		panic(err)
	}
	pupilInvoices := make([]entity.Invoice, 0)
	for _, invoice := range invoices {
		if invoice.PupilID == pupil.ID {
			pupilInvoices = append(pupilInvoices, invoice)
		}
	}
	pupil.Invoices = pupilInvoices

	fmt.Println(pupil)
}
