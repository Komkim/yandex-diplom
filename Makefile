.PHONY: mock-gen migrate-generate migrate-up migrate-down

MIGRATION_DIR := "storage/postgres/migrations"
PG_DSN := "postgres://postgres:changeme@localhost:5432/yandex?sslmode=disable"
#PG_DSN := "user=postgres password=changeme dbname=yandex sslmode=disable"

migrate-generate:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) create $(name) go

migrate-up:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) postgres $(PG_DSN) up

migrate-down:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) postgres $(PG_DSN) down

generate: mock-gen

mock-gen:
	@rm -rf ./test/mocks/packages
	@go generate ./...