package aggregate

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/vrischmann/envconfig"
	"yandex-diplom/config"
)

const skipTestMessage = "Skip test. please up local database for this test"

func getTestDB(req *require.Assertions) *pgxpool.Pool {
	ctx, _ := context.WithCancel(context.Background())

	type cfg struct {
		Dsn *config.Server
	}

	var c cfg
	err := envconfig.Init(&c)
	req.NoError(err)

	pool, err := pgxpool.New(ctx, c.Dsn.DatabaseDSN)
	req.NoError(err)

	return pool
}

func getTestUserRepo(db *pgxpool.Pool) UsersRepo {
	return NewUsersRepo(db)
}

func getTestOrderRepo(db *pgxpool.Pool) OrdersRepo {
	return NewOrdersRepo(db)
}
func getTestBalanceRepo(db *pgxpool.Pool) BalanceRepo {
	return NewBalanceRepo(db)
}
