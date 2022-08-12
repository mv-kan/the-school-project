package testingdb

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/mv-kan/the-school-project/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func mustGetPathToInitSQL() string {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	index := strings.LastIndex(workingDir, "src")
	return workingDir[:index] + strings.Join([]string{"src", "testing-db", "init.sql"}, string(os.PathSeparator))
}

func MustGetConnString(pconf config.PostgresConfig) string {
	container, err := RunTestingDB(pconf)
	if err != nil {
		panic(err)
	}
	connString, err := GetConnStringFromContainer(pconf, container)
	if err != nil {
		panic(err)
	}
	return connString
}

func GetConnStringFromContainer(pconf config.PostgresConfig, container testcontainers.Container) (string, error) {
	ctx := context.Background()

	p, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return "", err
	}
	pconf.Port = p.Port()
	pconf.Host, err = container.Host(ctx)
	if err != nil {
		return "", err
	}

	// create connection string
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pconf.User, pconf.Password, pconf.Host, pconf.Port, pconf.DBName)

	return dsn, nil
}

// returns connection string to postgres db, in pconf host and port are generated automatically
func RunTestingDB(pconf config.PostgresConfig) (testcontainers.Container, error) {
	var (
		DB_USER     = pconf.User
		DB_PASSWORD = pconf.Password
		DB_NAME     = pconf.DBName
	)
	// Work out the path to the 'scripts' directory and set mount strings
	mountFrom := mustGetPathToInitSQL()
	mountTo := "/docker-entrypoint-initdb.d/init.sql"

	// Create the Postgres TestContainer
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:14.4-bullseye",
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
		WaitingFor: wait.ForListeningPort("5432/tcp"),
		AutoRemove: true,
	}

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		// Panic and fail since there isn't much we can do if the container doesn't start
		return nil, err
	}

	return postgresC, nil
}
