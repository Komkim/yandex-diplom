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
		`select id, user_id, number, status, sum, create_at 
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
			&o.Number,
			&o.Status,
			&o.Sum,
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

func (o *Orders) GetAllByUserWithdrawals(userid uuid.UUID) ([]entity.Orders, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var orders = []entity.Orders{}
	rows, err := o.db.Query(ctx,
		`select id, user_id, number, status, sum, create_at 
			 from orders where user_id = $1 and sum < 0
    		 order by create_at asc limit 100;`,
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
			&o.Number,
			&o.Status,
			&o.Sum,
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
		`select id, user_id, number, status, sum, create_at 
			 from orders where number = $1
    		 order by create_at desc limit 1;`,
		number,
	).Scan(
		&orders.Id,
		&orders.UserId,
		&orders.Number,
		&orders.Status,
		&orders.Sum,
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

func (o *Orders) SetSum(number int64, userID uuid.UUID, sum float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into orders (user_id, number, sum, status)
		values ($1, $2)
		returning id`
	var id uuid.UUID
	err := o.db.QueryRow(ctx, sqlStatement, number, userID, sum, valueobject.PROCESSED).Scan(&id)
	if err != nil {
		return err
	}

	if id.ID() < 1 {
		return mistake.DbIdError
	}

	return nil
}
