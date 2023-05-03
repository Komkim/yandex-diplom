package model

import (
	"context"
	"errors"
	"time"
	"yandex-diplom/internal/mistake"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ModelGoods struct {
	ID          uuid.UUID `db:"id"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	CreateAt    time.Time `db:"create_at"`
}

type ModelReward struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Reward     float64   `db:"reward"`
	RewardType string    `db:"reward_type"`
	CreateAt   time.Time `db:"create_at"`
}

type Goods struct {
	db *pgxpool.Pool
}

func NewGoods(db *pgxpool.Pool) *Goods {
	return &Goods{db: db}
}

func (g *Goods) SetGoods(description string, price float64) (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into goods (description, price)
		values ($1, $2)
		returning id `
	var id uuid.UUID
	err := g.db.QueryRow(ctx, sqlStatement, description, price).Scan(&id)
	if err != nil {
		return nil, err
	}

	if id.ID() < 1 {
		return nil, mistake.ErrDBID
	}

	return &id, nil
}

func (g *Goods) SetReward(name, rewardType string, reward int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into reward (name, reward, reward_type)
		values ($1, $2, $3)
		returning id `
	var id uuid.UUID
	err := g.db.QueryRow(ctx, sqlStatement, name, reward, rewardType).Scan(&id)
	if err != nil {
		return err
	}

	if id.ID() < 1 {
		return mistake.ErrDBID
	}

	return nil
}

func (g *Goods) GetGoods(description string) (*ModelGoods, error) {

	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	goods := ModelGoods{}
	err := g.db.QueryRow(ctx,
		`select id, description, price, create_at 
			 from orders where description = $1
    		 order by create_at desc limit 1;`,
		description,
	).Scan(
		&goods.ID,
		&goods.Description,
		&goods.Price,
		&goods.CreateAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &goods, nil
}

func (g *Goods) GetReward(description string) (*ModelReward, error) {

	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	reward := ModelReward{}
	err := g.db.QueryRow(ctx,
		`select id, name, reward, reward_type, create_at 
			 from reward where name = $1
    		 order by create_at desc limit 1;`,
		description,
	).Scan(
		&reward.ID,
		&reward.Name,
		&reward.Reward,
		&reward.RewardType,
		&reward.CreateAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &reward, nil
}
