package aggregate

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"yandex-diplom/internal/domain/entity"
	"yandex-diplom/internal/domain/valueobject"
	"yandex-diplom/internal/mistake"
)

type OrdersRepo struct {
	db *pgxpool.Pool
}

func NewOrdersRepo(db *pgxpool.Pool) *OrdersRepo {
	return &OrdersRepo{db: db}
}

func (o *OrdersRepo) SetOne(number int64, userId, balanceId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into orders (user_id, balance_id, number, status)
		values ($1, $2, $3, $4)
		returning id `
	var id uuid.UUID
	err := o.db.QueryRow(ctx, sqlStatement, userId, balanceId, number, valueobject.NEW).Scan(&id)
	if err != nil {
		return err
	}

	if id.ID() < 1 {
		return mistake.DbIdError
	}

	return nil
}

func (o *OrdersRepo) GetAllByUser(userid string) (*entity.Orders, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	orders := entity.Orders{}
	err := o.db.QueryRow(ctx,
		`select id, user_id, balance_id, number, status, accrual, withdraw, create_at 
			 from orders where user_id = $1
    		 order by create_at desc limit 100;`,
		userid,
	).Scan(
		&orders.Id,
		&orders.UserId,
		&orders.BalanceId,
		&orders.Number,
		&orders.Status,
		&orders.Accrual,
		&orders.Withdraw,
		&orders.CreateAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (o *OrdersRepo) GetAllByUserWithdrawals(userid string) (*entity.Orders, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	orders := entity.Orders{}
	err := o.db.QueryRow(ctx,
		`select id, user_id, balance_id, number, status, accrual, withdraw, create_at 
			 from orders where user_id = $1 and withdraw <> null
    		 order by create_at desc limit 100;`,
		userid,
	).Scan(
		&orders.Id,
		&orders.UserId,
		&orders.BalanceId,
		&orders.Number,
		&orders.Status,
		&orders.Accrual,
		&orders.Withdraw,
		&orders.CreateAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &orders, nil
}
