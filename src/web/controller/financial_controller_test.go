package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
	testingdb "github.com/mv-kan/the-school-project/testing-db"
	"github.com/mv-kan/the-school-project/web/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFinancialController_GetAllLodgingDebtors(t *testing.T) {
	// init controller service
	db := connectToDB()
	log := logger.New()
	financialServ := service.NewFinancial(db)
	financialCtrl := NewFinancial(log, financialServ)

	// testing requesting
	path := "/financial/get-all-lodging-debtors"
	r := *mux.NewRouter()
	r.HandleFunc(path, financialCtrl.GetAllLodgingDebtors).Methods(http.MethodGet)

	req, err := http.NewRequest("GET", "/financial/get-all-lodging-debtors", nil)
	require.Nil(t, err)
	rr := httptest.NewRecorder()

	// call handler
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	debtors, err := utils.ParseJSONNoValidator[[]entity.Pupil](rr.Result().Body)
	require.Nil(t, err)
	assert.Equal(t, testingdb.TestDebtorID, debtors[0].ID)
}

func TestFinancialController_CollectedMoneyForMonth(t *testing.T) {
	// init controller service
	db := connectToDB()
	log := logger.New()
	financialServ := service.NewFinancial(db)
	financialCtrl := NewFinancial(log, financialServ)
	year, month := 2022, 4

	// testing requesting
	path := "/collected-money-for-month/{year}/{month}"
	r := *mux.NewRouter()
	r.HandleFunc(path, financialCtrl.CollectedMoneyForMonth).Methods(http.MethodGet)

	req, err := http.NewRequest("GET", fmt.Sprintf("/collected-money-for-month/%d/%d", year, month), nil)
	require.Nil(t, err)

	rr := httptest.NewRecorder()
	// call handler
	r.ServeHTTP(rr, req)
	// check response
	assert.Equal(t, http.StatusOK, rr.Code)
	amountOfMoneyMap, err := utils.ParseJSONNoValidator[map[string]decimal.Decimal](rr.Result().Body)
	require.Nil(t, err)
	amountOfMoney := amountOfMoneyMap["amountOfMoney"]
	assert.Equal(t, testingdb.TestAmountOfMoneyForForthMonth.String(), amountOfMoney.String())
}
