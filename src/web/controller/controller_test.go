package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mv-kan/the-school-project/config"
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
	testingdb "github.com/mv-kan/the-school-project/testing-db"
	"github.com/mv-kan/the-school-project/web/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connStr string

func connectToDB() *gorm.DB {
	dsn := connStr
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// TODO write tests for invoice note
// generic type in this testing is Invoice because I want to test also getting child entities
func TestMain(m *testing.M) {
	godotenv.Load("../../.env")
	var (
		DB_USER     = os.Getenv("TEST_DB_USER")
		DB_PASSWORD = os.Getenv("TEST_DB_PASSWORD")
		DB_NAME     = os.Getenv("TEST_DB_NAME")
	)
	pconf := config.PostgresConfig{
		User:     DB_USER,
		Password: DB_PASSWORD,
		DBName:   DB_NAME,
	}
	ctx := context.Background()
	postgresC, err := testingdb.RunTestingDB(pconf)
	if err != nil {
		// Panic and fail since there isn't much we can do if the container doesn't start
		panic(err)
	}

	defer postgresC.Terminate(ctx)

	// Get the port mapped to 5432 and set as ENV
	connStr, err = testingdb.GetConnStringFromContainer(pconf, postgresC)
	if err != nil {
		panic(err)
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestController_GetAll(t *testing.T) {
	// init service and controller
	db := connectToDB()
	log := logger.New()
	invoiceServ := service.NewInvoice(db)
	invoiceCtrl := New(log, invoiceServ)
	// init request and handler
	t.Run("success", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/dorms", nil)
		require.Nil(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(invoiceCtrl.GetAll)
		// call handler
		handler.ServeHTTP(rr, req)
		// check response
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestController_Get(t *testing.T) {
	// init service and controller
	db := connectToDB()
	log := logger.New()
	invoiceServ := service.NewInvoice(db)
	invoiceCtrl := New(log, invoiceServ)

	path := "/dorms/{id}"
	r := *mux.NewRouter()
	r.HandleFunc(path, invoiceCtrl.Get).Methods(http.MethodGet)

	// init request and handler
	t.Run("success", func(t *testing.T) {
		// req
		req, err := http.NewRequest("GET", "/dorms/1", nil)
		require.Nil(t, err)
		// response
		rr := httptest.NewRecorder()

		// call handler
		r.ServeHTTP(rr, req)
		// check response
		assert.Equal(t, http.StatusOK, rr.Code)
	})
	t.Run("not found", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/dorms/100", nil)
		require.Nil(t, err)
		rr := httptest.NewRecorder()

		// call handler
		r.ServeHTTP(rr, req)
		// check response
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestController_Create(t *testing.T) {
	// init service and controller
	db := connectToDB()
	log := logger.New()
	invoiceServ := service.NewInvoice(db)
	invoiceCtrl := New(log, invoiceServ)
	invoice := entity.Invoice{
		PaymentStart:    time.Now(),
		PaymentDue:      time.Now(),
		PupilID:         testingdb.TestPupilInDB.ID,
		TypeOfServiceID: testingdb.TestTypeOfServiceIDsInDB[0],
		AmountOfMoney:   decimal.NewFromFloat32(100),
	}
	// CREATE
	invoiceJSON, err := json.Marshal(invoice)
	require.Nil(t, err)
	req, err := http.NewRequest("POST", "/invoice", bytes.NewReader(invoiceJSON))
	require.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(invoiceCtrl.Create)
	// call handler
	handler.ServeHTTP(rr, req)
	// check response
	assert.Equal(t, http.StatusCreated, rr.Code)
	invoiceRes, err := utils.ParseJSONFromBody[entity.Invoice](rr.Result().Body)
	require.Nil(t, err)
	assert.Equal(t, invoice.AmountOfMoney.String(), invoiceRes.AmountOfMoney.String())
}

func TestController_Delete(t *testing.T) {
	// init service and controller
	db := connectToDB()
	log := logger.New()
	invoiceServ := service.NewInvoice(db)
	invoiceCtrl := New(log, invoiceServ)
	invoice := entity.Invoice{
		PaymentStart:    time.Now(),
		PaymentDue:      time.Now(),
		PupilID:         testingdb.TestPupilInDB.ID,
		TypeOfServiceID: testingdb.TestTypeOfServiceIDsInDB[0],
	}
	// CREATE
	invoiceJSON, err := json.Marshal(invoice)
	require.Nil(t, err)
	req, err := http.NewRequest("POST", "/invoice", bytes.NewReader(invoiceJSON))
	require.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(invoiceCtrl.Create)
	// call handler
	handler.ServeHTTP(rr, req)
	// check response
	assert.Equal(t, http.StatusCreated, rr.Code)
	invoiceRes, err := utils.ParseJSONFromBody[entity.Invoice](rr.Result().Body)
	require.Nil(t, err)

	// DELETE
	path := "/invoices/{id}"
	r := *mux.NewRouter()
	r.HandleFunc(path, invoiceCtrl.Delete).Methods(http.MethodDelete)

	req, err = http.NewRequest("DELETE", "/invoices/"+fmt.Sprint(invoiceRes.ID), nil)
	require.Nil(t, err)
	rr = httptest.NewRecorder()

	// call handler
	r.ServeHTTP(rr, req)
	// check response
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestController_Update(t *testing.T) {
	// init service and controller
	db := connectToDB()
	log := logger.New()
	invoiceServ := service.NewInvoice(db)
	invoiceCtrl := New(log, invoiceServ)
	invoice := entity.Invoice{
		PaymentStart:    time.Now(),
		PaymentDue:      time.Now(),
		PupilID:         testingdb.TestPupilInDB.ID,
		TypeOfServiceID: testingdb.TestTypeOfServiceIDsInDB[0],
	}
	// CREATE
	invoiceJSON, err := json.Marshal(invoice)
	require.Nil(t, err)
	req, err := http.NewRequest("POST", "/invoice", bytes.NewReader(invoiceJSON))
	require.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(invoiceCtrl.Create)
	// call handler
	handler.ServeHTTP(rr, req)
	// check response
	assert.Equal(t, http.StatusCreated, rr.Code)
	invoiceRes, err := utils.ParseJSONFromBody[entity.Invoice](rr.Result().Body)
	require.Nil(t, err)

	// UPDATE
	updatedMoney := decimal.NewFromFloat32(200)
	path := "/invoices/{id}"
	r := *mux.NewRouter()
	r.HandleFunc(path, invoiceCtrl.Update).Methods(http.MethodPut)
	// update invoice and create invoice json body for request
	invoiceRes.AmountOfMoney = updatedMoney
	invoiceJSON, err = json.Marshal(invoiceRes)
	require.Nil(t, err)

	req, err = http.NewRequest("PUT", "/invoices/"+fmt.Sprint(invoiceRes.ID), bytes.NewReader(invoiceJSON))
	require.Nil(t, err)
	rr = httptest.NewRecorder()

	// call handler
	r.ServeHTTP(rr, req)
	// check response
	assert.Equal(t, http.StatusOK, rr.Code)
	invoiceRes, err = utils.ParseJSONFromBody[entity.Invoice](rr.Result().Body)
	require.Nil(t, err)
	assert.Equal(t, updatedMoney.String(), invoiceRes.AmountOfMoney.String())
}
