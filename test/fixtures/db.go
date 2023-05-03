package fixtures

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Fixture interface {
	GetUserSQL() []string
	GetOrderSQL() []string
	GetBalanceSQL() []string
}

type CleanupFixture struct{}

func (cf CleanupFixture) GetUserSQL() []string {
	return []string{
		`TRUNCATE TABLE users RESTART IDENTITY CASCADE;`,
	}
}

func (cf CleanupFixture) GetOrderSQL() []string {
	return []string{
		`TRUNCATE TABLE orders RESTART IDENTITY CASCADE;`,
	}
}

func (cf CleanupFixture) GetBalanceSQL() []string {
	return []string{
		`TRUNCATE TABLE balance RESTART IDENTITY CASCADE;`,
	}
}

func ExecuteFixture(ctx context.Context, db *pgxpool.Pool, fixture Fixture) {
	for _, query := range fixture.GetUserSQL() {
		_, err := db.Exec(ctx, query)

		if err != nil {
			panic(err)
		}
	}

	for _, query := range fixture.GetOrderSQL() {
		_, err := db.Exec(ctx, query)

		if err != nil {
			panic(err)
		}
	}
	for _, query := range fixture.GetBalanceSQL() {
		_, err := db.Exec(ctx, query)

		if err != nil {
			panic(err)
		}
	}
}
