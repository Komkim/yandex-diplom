package model

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"yandex-diplom/internal/mistake"
)

type ModelOrders struct {
	ID       uuid.UUID `db:"id"`
	Number   int64     `db:"number"`
	Status   string    `db:"status"`
	Reward   float64   `db:"reward"`
	CreateAt time.Time `db:"create_at"`
}

type Orders struct {
	db *pgxpool.Pool
}

func NewOrders(db *pgxpool.Pool) *Orders {
	return &Orders{db: db}
}

func (o *Orders) SetOrders(number int64, status string, reward float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into orders (number, status, reward)
		values ($1, $2, $3)
		returning id `
	var id uuid.UUID
	err := o.db.QueryRow(ctx, sqlStatement, number, status, reward).Scan(&id)
	if err != nil {
		return err
	}

	if id.ID() < 1 {
		return mistake.ErrDBID
	}

	return nil
}

func (o *Orders) GetOrdersByNumber(number int64) (*ModelOrders, error) {

	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	orders := ModelOrders{}
	err := o.db.QueryRow(ctx,
		`select id, number, status, reward, create_at 
			 from orders where number = $1
    		 order by create_at desc limit 1;`,
		number,
	).Scan(
		&orders.ID,
		&orders.Number,
		&orders.Status,
		&orders.Reward,
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
