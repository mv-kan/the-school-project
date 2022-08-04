package main

import (
	"fmt"
	"time"

	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
	"github.com/mv-kan/the-school-project/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func testServices(db *gorm.DB) {
	// test services
	invoiceRepo := repo.New[entity.Invoice](db)
	pupilRepo := repo.New[entity.Pupil](db)
	financialService := service.NewFinancial(invoiceRepo, pupilRepo)
	collectedMoney, err := financialService.CollectedMoneyForMonth(time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println("collected money for month", collectedMoney)

	debtors, err := financialService.FindAllLodgingDebtors()
	if err != nil {
		panic(err)
	}
	fmt.Println("debtors:")
	for _, pupil := range debtors {
		fmt.Println(pupil)
	}

	// test enroll

	enrollService := service.NewEnroll(db)
	pupilToEnroll := entity.Pupil{
		Name:          "john",
		Surname:       "johnson",
		SchoolClassID: 1,
	}
	err = enrollService.Enroll(pupilToEnroll)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success?")
	}
}

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
	_, err = pupilRepo.Save(entity.Pupil{
		Name:          "soemthing",
		Surname:       "back",
		SchoolClassID: 1,
	})
	if err != nil {
		panic(err)
	}
	testServices(db)
}
