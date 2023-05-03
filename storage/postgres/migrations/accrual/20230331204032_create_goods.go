package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateGoods, downCreateGoods)
}

func upCreateGoods(tx *sql.Tx) error {
	_, err := tx.Exec(`
			create table goods
			(
				id          uuid default gen_random_uuid() not null primary key,
				description varchar(50) not null,
				price       double precision,
				create_at   timestamp with time zone default current_timestamp
			);
	`)
	return err
}

func downCreateGoods(tx *sql.Tx) error {
	_, err := tx.Exec("drop table goods")
	return err
}
