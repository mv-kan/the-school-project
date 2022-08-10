package service

import (
	"testing"

	"github.com/mv-kan/the-school-project/entity"
	testingdb "github.com/mv-kan/the-school-project/testing-db"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvoiceService_Create(t *testing.T) {
	var amount = decimal.NewFromFloat(100)
	var invoice = entity.Invoice{}
	db := connectToDB()

	invoiceService := NewInvoice(db)
	invoice = entity.Invoice{
		PupilID:         testingdb.TestPupilInDB.ID,
		AmountOfMoney:   amount,
		TypeOfServiceID: testingdb.TestTypeOfServiceID,
	}
	invoice, err := invoiceService.Create(invoice)
	require.Nil(t, err)
	assert.Equal(t, amount.String(), invoice.AmountOfMoney.String())
}

func TestInvoiceService_Update(t *testing.T) {
	var amount = decimal.NewFromFloat(100)
	db := connectToDB()
	invoiceService := NewInvoice(db)

	// create
	invoice := entity.Invoice{
		PupilID:         testingdb.TestPupilInDB.ID,
		AmountOfMoney:   amount,
		TypeOfServiceID: testingdb.TestTypeOfServiceID,
	}
	invoice, err := invoiceService.Create(invoice)
	require.Nil(t, err)
	assert.Equal(t, amount, invoice.AmountOfMoney)

	// update
	updatedAmount := decimal.NewFromFloat(200)
	invoice.AmountOfMoney = updatedAmount
	err = invoiceService.Update(invoice)
	require.Nil(t, err)
	invoice, err = invoiceService.Find(invoice.ID)
	require.Nil(t, err)
	assert.Equal(t, updatedAmount.String(), invoice.AmountOfMoney.String())
}

func TestInvoiceService_Delete(t *testing.T) {
	var amount = decimal.NewFromFloat(100)
	var invoice = entity.Invoice{
		PupilID:         testingdb.TestPupilInDB.ID,
		AmountOfMoney:   amount,
		TypeOfServiceID: testingdb.TestTypeOfServiceID,
	}
	db := connectToDB()

	// create
	invoiceService := NewInvoice(db)
	invoice, err := invoiceService.Create(invoice)
	require.Nil(t, err)
	assert.Equal(t, amount, invoice.AmountOfMoney)

	// delete
	err = invoiceService.Delete(invoice.ID)
	require.Nil(t, err)
	_, err = invoiceService.Find(invoice.ID)
	assert.NotNil(t, err)
}

func TestInvoiceService_Find(t *testing.T) {
	// test find by id function
	db := connectToDB()
	invoiceService := NewInvoice(db)
	invoice, err := invoiceService.Find(testingdb.TestInvoiceInDB.ID)
	require.Nil(t, err)
	assert.Equal(t, testingdb.TestInvoiceInDB.AmountOfMoney.String(), invoice.AmountOfMoney.String())
	assert.Equal(t, testingdb.TestInvoiceInDB.Note.Note, invoice.Note.Note)
}

func TestInvoiceService_FindAll(t *testing.T) {
	// test find all function
	db := connectToDB()
	invoiceService := NewInvoice(db)
	_, err := invoiceService.FindAll()
	require.Nil(t, err)
}
