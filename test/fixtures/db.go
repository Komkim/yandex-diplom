package fixtures

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Fixture interface {
	GetUserSql() []string
	GetOrderSql() []string
	GetBalanceSql() []string
}

type CleanupFixture struct{}

func (cf CleanupFixture) GetUserSql() []string {
	return []string{
		`TRUNCATE TABLE users RESTART IDENTITY CASCADE;`,
	}
}

func (cf CleanupFixture) GetOrderSql() []string {
	return []string{
		`TRUNCATE TABLE orders RESTART IDENTITY CASCADE;`,
	}
}

func (cf CleanupFixture) GetBalanceSql() []string {
	return []string{
		`TRUNCATE TABLE balance RESTART IDENTITY CASCADE;`,
	}
}

func ExecuteFixture(ctx context.Context, db *pgxpool.Pool, fixture Fixture) {
	for _, query := range fixture.GetUserSql() {
		_, err := db.Exec(ctx, query)

		if err != nil {
			panic(err)
		}
	}

	for _, query := range fixture.GetOrderSql() {
		_, err := db.Exec(ctx, query)

		if err != nil {
			panic(err)
		}
	}
	for _, query := range fixture.GetBalanceSql() {
		_, err := db.Exec(ctx, query)

		if err != nil {
			panic(err)
		}
	}
}
