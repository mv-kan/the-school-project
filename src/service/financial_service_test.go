package service

import (
	"testing"

	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDebptorID int = 3
var testCollectedMoney = decimal.NewFromFloat(1260)

func TestFinancialService_FindAllLodgingDebtors(t *testing.T) {
	db := connectToDB()
	pupilRepo := repo.New[entity.Pupil](db)
	invoiceRepo := repo.New[entity.Invoice](db)
	financialService := NewFinancial(invoiceRepo, pupilRepo)

	pupils, err := financialService.FindAllLodgingDebtors()
	require.Nil(t, err)

	assert.Equal(t, 1, len(pupils))
	assert.Equal(t, testDebptorID, pupils[0].ID)
}

func TestFinancialService_CollectedMoneyForMonth(t *testing.T) {
	db := connectToDB()
	pupilRepo := repo.New[entity.Pupil](db)
	invoiceRepo := repo.New[entity.Invoice](db)
	financialService := NewFinancial(invoiceRepo, pupilRepo)

	money, err := financialService.CollectedMoneyForMonth(2022, 8)
	require.Nil(t, err)

	assert.Equal(t, testCollectedMoney.String(), money.String())
}
