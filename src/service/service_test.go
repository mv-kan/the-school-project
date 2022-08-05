package service

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mv-kan/the-school-project/entity"
	"github.com/mv-kan/the-school-project/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// it is very important to keep this test values in sync with seeded values in "testing-db" folder
var testDorm = entity.Dormitory{ID: 1, Name: "Laura"}

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
	dorm, err := dormServ.Find(testDorm.ID)

	require.Nil(t, err)
	assert.Equal(t, testDorm, dorm)
}
