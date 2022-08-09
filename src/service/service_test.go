package service

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// it is very important to keep this test values in sync with seeded values in "testing-db" folder
var testPupilInDB = entity.Pupil{
	ID:            1,
	Name:          "michael",
	Surname:       "lan",
	SchoolClassID: 3,
	RoomID:        sql.NullInt32{Int32: 2, Valid: true},
}

var testDormInDB = entity.Dormitory{ID: 1, Name: "Laura"}
var testDormToCreate = entity.Dormitory{Name: "Bartek"}
var testDorms = []entity.Dormitory{testDormInDB, {ID: 2, Name: "Laura"}}
var testDebptorID int = 3
var testCollectedMoney = decimal.NewFromFloat(760)
var testPupil = entity.Pupil{
	Name:          "new pupil name",
	Surname:       "new pupil surname",
	SchoolClassID: 1,
}
var testRoomID = 2
var testAvailableSpace = 1
var testResidentID = 2
var testRoomTypeID = 2
var testTypeOfServiceID = 1
var testInvoiceID = 2
var testInvoiceAmountOfMoney = decimal.NewFromFloat(380)
var testInvoiceInDB = entity.Invoice{
	ID:              1,
	AmountOfMoney:   decimal.NewFromFloat(0),
	PupilID:         1,
	TypeOfServiceID: 1,
	DateOfPayment:   time.Date(2022, time.August, 2, 0, 0, 0, 0, time.UTC),
	PaymentStart:    time.Date(2022, time.August, 2, 0, 0, 0, 0, time.UTC),
	PaymentDue:      time.Date(2022, time.September, 2, 0, 0, 0, 0, time.UTC),
	Note:            &entity.InvoiceNote{ID: 1, Note: "Bogdana did not pay because she have got help from our school for good grades"},
}

func connectToDB() *gorm.DB {
	var (
		DB_USER     = os.Getenv("DBUSER")
		DB_PASSWORD = os.Getenv("DBPASSWORD")
		DB_HOST     = os.Getenv("DBHOST")
		DB_PORT     = os.Getenv("DBPORT")
		DB_NAME     = os.Getenv("DBNAME")
	)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func TestMain(m *testing.M) {
	godotenv.Load("../.env")
	var (
		DB_USER     = os.Getenv("DBUSER")
		DB_PASSWORD = os.Getenv("DBPASSWORD")
		DB_NAME     = os.Getenv("DBNAME")
	)
	// Work out the path to the 'scripts' directory and set mount strings
	packageName := "service"
	workingDir, _ := os.Getwd()
	rootDir := strings.Replace(workingDir, packageName, "", 1)
	mountFrom := fmt.Sprintf("%s/testing-db/init.sql", rootDir)
	mountTo := "/docker-entrypoint-initdb.d/init.sql"

	// Create the Postgres TestContainer
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:14.4-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Mounts: testcontainers.ContainerMounts{
			{
				Source: testcontainers.GenericBindMountSource{HostPath: mountFrom},
				Target: testcontainers.ContainerMountTarget(mountTo)}},
		Env: map[string]string{
			"POSTGRES_USER":     DB_USER,
			"POSTGRES_PASSWORD": DB_PASSWORD,
			"POSTGRES_DB":       DB_NAME,
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections"),
	}

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		// Panic and fail since there isn't much we can do if the container doesn't start
		panic(err)
	}

	defer postgresC.Terminate(ctx)

	// Get the port mapped to 5432 and set as ENV
	p, _ := postgresC.MappedPort(ctx, "5432")
	os.Setenv("DBPORT", p.Port())

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestService_Find(t *testing.T) {
	db := connectToDB()

	dormRepo := repo.New[entity.Dormitory](db)
	dormServ := New(dormRepo)
	dorm, err := dormServ.Find(testDormInDB.ID)

	require.Nil(t, err)
	assert.Equal(t, testDormInDB, dorm)
}

// test for service find all
func TestService_FindAll(t *testing.T) {
	db := connectToDB()

	dormRepo := repo.New[entity.Dormitory](db)
	dormServ := New(dormRepo)
	dorms, err := dormServ.FindAll()

	require.Nil(t, err)
	assert.Equal(t, 2, len(dorms))
}

func TestService_CreateUpdateDelete(t *testing.T) {
	db := connectToDB()
	var (
		dorm entity.Dormitory
		err  error
	)

	t.Run("Save", func(t *testing.T) {
		dormRepo := repo.New[entity.Dormitory](db)
		dormServ := New(dormRepo)
		dorm, err = dormServ.Create(testDormToCreate)
		require.Nil(t, err)
		assert.Equal(t, testDormToCreate.Name, dorm.Name)
	})
	t.Run("Update", func(t *testing.T) {
		updatedName := "BartekUpdated"
		dormRepo := repo.New[entity.Dormitory](db)
		dormServ := New(dormRepo)
		dorm.Name = updatedName
		err := dormServ.Update(dorm)
		require.Nil(t, err)

		dorm, err = dormServ.Find(dorm.ID)
		require.Nil(t, err)

		dorm, err := dormServ.Find(dorm.ID)
		assert.Equal(t, updatedName, dorm.Name)
	})
	t.Run("Delete", func(t *testing.T) {
		dormRepo := repo.New[entity.Dormitory](db)
		dormServ := New(dormRepo)
		err := dormServ.Delete(dorm.ID)
		require.Nil(t, err)
		_, err = dormServ.Find(dorm.ID)
		assert.NotNil(t, err)
	})
}
