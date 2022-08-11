package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mv-kan/the-school-project/config"
	"github.com/mv-kan/the-school-project/logger"
	"github.com/mv-kan/the-school-project/service"
	testingdb "github.com/mv-kan/the-school-project/testing-db"
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
	// init request and handler
	t.Run("success", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/dorms/1", nil)
		require.Nil(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(invoiceCtrl.Get)
		// call handler
		handler.ServeHTTP(rr, req)
		// check response
		assert.Equal(t, http.StatusOK, rr.Code)
	})
	t.Run("not found", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/dorms/100", nil)
		require.Nil(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(invoiceCtrl.Get)
		// call handler
		handler.ServeHTTP(rr, req)
		// check response
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
