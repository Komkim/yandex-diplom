package aggregate

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"yandex-diplom/internal/domain/entity"
)

type BalanceRepo struct {
	db *pgxpool.Pool
}

func NewBalanceRepo(db *pgxpool.Pool) *BalanceRepo {
	return &BalanceRepo{db: db}
}

func (b *BalanceRepo) GetByUser(userId string) (*entity.Balance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	balance := entity.Balance{}
	err := b.db.QueryRow(ctx,
		`with unic_balance as (
					select distinct upload_at as u
					from balance
				)
				select id, user_id, number, current, withdraw, upload_at
				from balance,
					 unic_balance
				where upload_at in (select upload_at from balance where upload_at = unic_balance.u order by upload_at desc limit 1)
					and name=$1;`,
		userId,
	).Scan(&balance.Id, &balance.User_id, &balance.Number, &balance.Current, &balance.Withdraw, &balance.UploadAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &balance, nil
}

func (o *BalanceRepo) SetOne(user_id string) error {
	//ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	//defer cancel()
	//
	//sqlStatement := `
	//	insert into orders (user_id, balance_id, number, status)
	//	values ($1, $2, $3, $4)
	//	returning id `
	//var id uuid.UUID
	//err := o.db.QueryRow(ctx, sqlStatement, userId, balanceId, number, valueobject.NEW).Scan(&id)
	//if err != nil {
	//	return err
	//}
	//
	//if id.ID() < 1 {
	//	return mistake.DbIdError
	//}

	return nil
}
