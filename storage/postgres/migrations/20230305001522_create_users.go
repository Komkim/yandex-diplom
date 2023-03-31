package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateUsers, downCreateUsers)
}

func upCreateUsers(tx *sql.Tx) error {
	//_, err := tx.Exec('create table test(	id serial);')
	_, err := tx.Exec(`
			create table users
			(
				id               uuid default gen_random_uuid() not null primary key,
				login	         varchar(40) not null,
				hashed_password  varchar(500) not null,
				create_at        timestamp with time zone default current_timestamp
			);
	`)
	return err
}

func downCreateUsers(tx *sql.Tx) error {
	_, err := tx.Exec("drop table users")
	return err
}