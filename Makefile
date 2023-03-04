MIGRATION_DIR := "storage/postgres/migrations"

.PHONY: migrate-generate
migrate-generate:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) create $(name) go

.PHONY: migrate-up
migrate-up:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) create $(name) up

.PHONY: migrate-down
migrate-down:
	$(GOPATH)/bin/goose -dir $(MIGRATION_DIR) create $(name) down