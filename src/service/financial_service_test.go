package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFinancialService_FindAllLodgingDebtors(t *testing.T) {
	db := connectToDB()
	financialService := NewFinancial(db)

	pupils, err := financialService.FindAllLodgingDebtors()
	require.Nil(t, err)

	assert.Equal(t, 1, len(pupils))
	assert.Equal(t, testDebptorID, pupils[0].ID)
}

func TestFinancialService_CollectedMoneyForMonth(t *testing.T) {
	db := connectToDB()

	financialService := NewFinancial(db)

	money, err := financialService.CollectedMoneyForMonth(2022, 4)
	require.Nil(t, err)

	assert.Equal(t, testCollectedMoney.String(), money.String())
}
