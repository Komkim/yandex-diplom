package model

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"yandex-diplom/internal/mistake"
)

type ModelConnect struct {
	ID       uuid.UUID `db:"id"`
	GoodsID  uuid.UUID `db:"goods_id"`
	Number   int64     `db:"number"`
	CreateAt time.Time `db:"create_at"`
}

type Connect struct {
	db *pgxpool.Pool
}

func NewConnect(db *pgxpool.Pool) *Connect {
	return &Connect{db: db}
}

func (c *Connect) SetConnect(number int64, goodsID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into connect (goods_id, number)
		values ($1, $2)
		returning id `
	var id uuid.UUID
	err := c.db.QueryRow(ctx, sqlStatement, goodsID, number).Scan(&id)
	if err != nil {
		return err
	}

	if id.ID() < 1 {
		return mistake.ErrDBID
	}

	return nil
}
