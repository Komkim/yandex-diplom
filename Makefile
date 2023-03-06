MIGRATION_DIR := "storage/postgres/migrations"
PG_DSN := "postgres://postgres:changeme@localhost:5432/yandex?sslmode=disable"

.PHONY: migrate-generate
migrate-generate:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) create $(name) go

.PHONY: migrate-up
migrate-up:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) postgres $(name) up

.PHONY: migrate-down
migrate-down:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) postgres $(PG_DSN) down