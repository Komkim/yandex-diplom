package aggregate

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"yandex-diplom/internal/domain/entity"
	"yandex-diplom/internal/mistake"
)

type Users struct {
	db *pgxpool.Pool
}

func NewUsersRepo(db *pgxpool.Pool) *Users {
	return &Users{db: db}
}

func (u *Users) GetOne(login string) (*entity.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	user := entity.Users{}
	err := u.db.QueryRow(ctx,
		`select id, login, hashed_password, create_at from users where login = $1;`,
		login,
	).Scan(&user.ID, &user.Login, &user.HashedPassword, &user.CreateAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *Users) SetOne(login, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into users (login, hashed_password)
		values ($1, $2)
		returning id `
	var id uuid.UUID
	err := u.db.QueryRow(ctx, sqlStatement, login, password).Scan(&id)
	if err != nil {
		return err
	}

	if id.ID() < 1 {
		return mistake.ErrDbId
	}

	return nil
}
