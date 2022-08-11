package service

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mv-kan/the-school-project/config"
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
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

func TestMain(m *testing.M) {
	godotenv.Load("../.env")
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

func TestService_Find(t *testing.T) {
	db := connectToDB()

	dormRepo := repo.New[entity.Dormitory](db)
	dormServ := New(dormRepo)
	dorm, err := dormServ.Find(testingdb.TestDormInDB.ID)

	require.Nil(t, err)
	assert.Equal(t, testingdb.TestDormInDB, dorm)
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
		dorm          entity.Dormitory
		err           error
		idToOverwrite int = 1
	)

	t.Run("Save", func(t *testing.T) {
		dormRepo := repo.New[entity.Dormitory](db)
		dormServ := New(dormRepo)
		dorm = testingdb.TestDormToCreate

		// this id is not going to be saved in the db because in repo it is autoincrement and this values is overwritten by repo
		dorm.ID = idToOverwrite
		dorm, err = dormServ.Create(dorm)
		require.Nil(t, err)
		assert.Equal(t, testingdb.TestDormToCreate.Name, dorm.Name)
		assert.NotEqual(t, idToOverwrite, dorm.ID)
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
