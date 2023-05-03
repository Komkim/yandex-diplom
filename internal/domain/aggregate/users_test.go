package aggregate

import (
	"context"
	"testing"
	"yandex-diplom/internal/domain/entity"
	"yandex-diplom/test/fixtures"

	"github.com/stretchr/testify/require"
)

func TestRepositoryDB_SetOne(t *testing.T) {
	if testing.Short() {
		t.Skip(skipTestMessage)
	}

	var (
		req  = require.New(t)
		ctx  = context.Background()
		db   = getTestDB(req)
		repo = getTestUserRepo(db)
	)

	defer func() {
		db.Close()
	}()

	fixtures.ExecuteFixture(ctx, db, fixtures.CleanupFixture{})

	test := func(login, password string) func(t *testing.T) {
		return func(t *testing.T) {
			err := repo.SetOne(login, password)
			req.NoError(err)

			user := entity.Users{}
			err = db.QueryRow(ctx,
				`select id, login, hashed_password, create_at from users where login = $1;`,
				login,
			).Scan(&user.ID, &user.Login, &user.HashedPassword, &user.CreateAt)
			req.NoError(err)

			req.Equal(login, user.Login)
			req.Equal(password, user.HashedPassword)
		}
	}

	t.Run("insert user", test("user", "password"))
}

func TestRepositoryDB_GetOne(t *testing.T) {
	if testing.Short() {
		t.Skip(skipTestMessage)
	}

	var (
		req  = require.New(t)
		ctx  = context.Background()
		db   = getTestDB(req)
		repo = getTestUserRepo(db)
	)

	defer func() {
		db.Close()
	}()

	fixtures.ExecuteFixture(ctx, db, fixtures.CleanupFixture{})

	type u struct {
		Login    string
		Password string
	}

	users := []entity.Users{
		{Login: "login1", HashedPassword: "password1"},
		{Login: "login2", HashedPassword: "password2"},
	}

	for _, tempUser := range users {
		err := repo.SetOne(tempUser.Login, tempUser.HashedPassword)
		req.NoError(err)
	}

	test := func(login string, want *entity.Users, isError bool) func(t *testing.T) {
		return func(t *testing.T) {
			actual, err := repo.GetOne(login)
			if isError {
				req.Error(err)
			} else {
				req.NoError(err)
			}
			req.Equal(want.Login, actual.Login)
			req.Equal(want.HashedPassword, actual.HashedPassword)
		}
	}

	t.Run("get users", test(users[0].Login, &users[0], false))
	t.Run("get users", test(users[1].Login, &users[1], false))
}
