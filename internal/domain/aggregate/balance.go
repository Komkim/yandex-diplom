package aggregate

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"yandex-diplom/internal/domain/entity"
	"yandex-diplom/internal/mistake"
)

type Balance struct {
	db *pgxpool.Pool
}

func NewBalanceRepo(db *pgxpool.Pool) *Balance {
	return &Balance{db: db}
}

func (b *Balance) GetByUser(userId uuid.UUID) (*entity.Balance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	balance := entity.Balance{}
	err := b.db.QueryRow(ctx,
		`with unic_balance as (
					select distinct upload_at as u
					from balance
				)
				select id, user_id, sum, create_at
				from balance,
					 unic_balance
				where upload_at in (select upload_at from balance where upload_at = unic_balance.u order by upload_at desc limit 1)
					and name=$1;`,
		userId,
	).Scan(&balance.Id, &balance.UserID, &balance.Sum, &balance.CreateAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &balance, nil
}

func (b *Balance) GetCurrentByUser(userID uuid.UUID) (*float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var current sql.NullFloat64
	err := b.db.QueryRow(ctx,
		`select sum(sum) from balance where user_id = $1 group by user_id;`,
		userID,
	).Scan(current)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &current.Float64, nil
}

func (b *Balance) GetWithdrawntByUser(userId uuid.UUID) (*float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var withdrawn sql.NullFloat64
	err := b.db.QueryRow(ctx,
		`select sum(sum) from balance where user_id = $1 and sum < 0 group by user_id;`,
		userId,
	).Scan(withdrawn)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &withdrawn.Float64, nil
}

func (o *Balance) SetSum(userID uuid.UUID, sum float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into balance (user_id, sum)
		values ($1, $2)
		returning id`
	var id uuid.UUID
	err := o.db.QueryRow(ctx, sqlStatement, userID, sum).Scan(&id)
	if err != nil {
		return err
	}

	if id.ID() < 1 {
		return mistake.DbIdError
	}

	return nil
}
