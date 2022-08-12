package testingdb

import (
	"database/sql"
	"time"

	"github.com/mv-kan/the-school-project/entity"
	"github.com/shopspring/decimal"
)

// this looks like a pile of garbage, I will eventually refactor it to be neater (*maybe)
// it is very important to keep this test values in sync with seeded values in "testing-db" folder
var TestPupilInDB = entity.Pupil{
	ID:            1,
	Name:          "michael",
	Surname:       "lan",
	SchoolClassID: 3,
	RoomID:        sql.NullInt32{Int32: 2, Valid: true},
}
var TestInvoiceInDB = entity.Invoice{
	ID:              1,
	AmountOfMoney:   decimal.NewFromFloat(0),
	PupilID:         1,
	TypeOfServiceID: 1,
	DateOfPayment:   time.Date(2022, time.August, 2, 0, 0, 0, 0, time.UTC),
	PaymentStart:    time.Date(2022, time.August, 2, 0, 0, 0, 0, time.UTC),
	PaymentDue:      time.Date(2022, time.September, 2, 0, 0, 0, 0, time.UTC),
	Note:            &entity.InvoiceNote{ID: 1, Note: "Bogdana did not pay because she have got help from our school for good grades"},
}

var TestDormInDB = entity.Dormitory{ID: 1, Name: "Laura"}
var TestSchoolClassID = 1
var TestDormToCreate = entity.Dormitory{Name: "Bartek"}
var TestDorms = []entity.Dormitory{TestDormInDB, {ID: 2, Name: "Laura"}}
var TestDebptorID int = 3
var TestCollectedMoney = decimal.NewFromFloat(760)
var TestPupil = entity.Pupil{
	Name:          "new pupil name",
	Surname:       "new pupil surname",
	SchoolClassID: 1,
}
var TestRoomID = 2
var TestAvailableSpace = 1
var TestResidentID = 2
var TestRoomTypeID = 2
var TestTypeOfServiceID = 1
var TestInvoiceID = 2
var TestInvoiceAmountOfMoney = decimal.NewFromFloat(380)
var TestTypeOfServiceIDsInDB = []int{1, 2}
var TestAmountOfMoneyForForthMonth = decimal.NewFromFloat(760)
var TestDebtorID = 3
