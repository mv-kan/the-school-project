DB_USER=postgres
DB_PASSWORD=postgrespw
DB_HOST=localhost
DB_PORT=49153
DB_NAME=postgres
DB_URL=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
MIGRATION_DIR=./db/migrate
SEEDERS_DIR=./db/seeders
CONTAINER_NAME=psql
# Default number of migration(s)
N = 1
# Migration schema version
V = 1

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq -digits 5 $(NAME)

migrate-up:
	migrate -source file://$(MIGRATION_DIR) -database $(DB_URL) up

migrate-down:
	migrate -source file://$(MIGRATION_DIR) -database $(DB_URL) down $(N)

migrate-test-up:
	migrate -source file://$(MIGRATION_DIR) -database $(TEST_DB_URL) up

migrate-test-down:
	migrate -source file://$(MIGRATION_DIR) -database $(TEST_DB_URL) down $(N)

seeders-create:
	migrate create -ext sql -dir $(SEEDERS_DIR) -seq -digits 5 $(NAME)

seeders-up:
	migrate -source file://$(SEEDERS_DIR) -database $(DB_URL) up

seeders-force:
	migrate -source file://$(SEEDERS_DIR) -database $(DB_URL) force $(V)

seeders-down:
	migrate -source file://$(SEEDERS_DIR) -database $(DB_URL) down

drop-db:
	migrate -source file://$(MIGRATION_DIR) -database $(DB_URL) drop -f

enter-db:
	docker exec -it $(CONTAINER_NAME) psql -U ${DB_USER}