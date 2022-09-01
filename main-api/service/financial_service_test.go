package service

import (
	"testing"

	testingdb "github.com/mv-kan/the-school-project/testing-db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFinancialService_FindAllLodgingDebtors(t *testing.T) {
	db := connectToDB()
	financialService := NewFinancial(db)

	pupils, err := financialService.FindAllLodgingDebtors()
	require.Nil(t, err)

	assert.Equal(t, 1, len(pupils))
	assert.Equal(t, testingdb.TestDebptorID, pupils[0].ID)
}

func TestFinancialService_CollectedMoneyForMonth(t *testing.T) {
	db := connectToDB()

	financialService := NewFinancial(db)

	money, err := financialService.CollectedMoneyForMonth(2022, 4)
	require.Nil(t, err)

	assert.Equal(t, testingdb.TestCollectedMoney.String(), money.String())
}
