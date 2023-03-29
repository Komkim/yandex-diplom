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

type Orders struct {
	db *pgxpool.Pool
}

func NewOrdersRepo(db *pgxpool.Pool) *Orders {
	return &Orders{db: db}
}

func (o *Orders) SetOne(number int64, userId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into orders (user_id, number, status)
		values ($1, $2, $3)
		returning id `
	var id uuid.UUID
	err := o.db.QueryRow(ctx, sqlStatement, userId, number, valueobject.NEW).Scan(&id)
	if err != nil {
		return err
	}

	if id.ID() < 1 {
		return mistake.DbIdError
	}

	return nil
}

func (o *Orders) GetAllByUser(userid uuid.UUID) ([]entity.Orders, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var orders = []entity.Orders{}
	rows, err := o.db.Query(ctx,
		`select id, user_id, balance_id, number, status, accrual, withdraw, create_at 
			 from orders where user_id = $1
    		 order by create_at desc limit 100;`,
		userid,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := entity.Orders{}
		err := rows.Scan(
			&o.Id,
			&o.UserId,
			&o.BalanceId,
			&o.Number,
			&o.Status,
			&o.Accrual,
			&o.Withdraw,
			&o.CreateAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *Orders) GetAllByUserWithdrawals(userid string) ([]entity.Orders, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var orders = []entity.Orders{}
	rows, err := o.db.Query(ctx,
		`select id, user_id, balance_id, number, status, accrual, withdraw, create_at 
			 from orders where user_id = $1 and withdraw <> null
    		 order by create_at desc limit 100;`,
		userid,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := entity.Orders{}
		err := rows.Scan(
			&o.Id,
			&o.UserId,
			&o.BalanceId,
			&o.Number,
			&o.Status,
			&o.Accrual,
			&o.Withdraw,
			&o.CreateAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *Orders) GetByNumber(number int64) (*entity.Orders, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	orders := entity.Orders{}
	err := o.db.QueryRow(ctx,
		`select id, user_id, balance_id, number, status, accrual, withdraw, create_at 
			 from orders where number = $1
    		 order by create_at desc limit 1;`,
		number,
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
