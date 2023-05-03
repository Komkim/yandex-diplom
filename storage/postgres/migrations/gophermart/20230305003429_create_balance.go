package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateBalance, downCreateBalance)
}

func upCreateBalance(tx *sql.Tx) error {
	_, err := tx.Exec(`
			create table balance
			(
				id        uuid default gen_random_uuid() not null primary key,
				user_id   uuid REFERENCES users (id),							
				sum       double precision,			
				create_at timestamp with time zone default current_timestamp
			);
	`)
	return err
}

func downCreateBalance(tx *sql.Tx) error {
	_, err := tx.Exec("drop table balance")
	return err
}
