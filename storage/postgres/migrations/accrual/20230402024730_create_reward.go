package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateReward, downCreateReward)
}

func upCreateReward(tx *sql.Tx) error {
	_, err := tx.Exec(`
			create table reward
			(
				id          uuid default gen_random_uuid() not null primary key,
				name        varchar(50) not null,
				reward      int,
				reward_type reward_type,
				create_at   timestamp with time zone default current_timestamp
			);
	`)
	return err
}

func downCreateReward(tx *sql.Tx) error {
	_, err := tx.Exec("drop table reward")
	return err
}
