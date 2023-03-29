package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateOrders, downCreateOrders)
}

func upCreateOrders(tx *sql.Tx) error {
	_, err := tx.Exec(`
			create table orders
			(
				id          uuid default gen_random_uuid() not null primary key,
				user_id     uuid REFERENCES users (id),
				balance_id  uuid REFERENCES balance (id),				
				number	    bigint not null,
				status      status not null,
				accrual     double precision,
				withdraw    double precision,
				create_at   timestamp with time zone default current_timestamp
			);
	`)
	return err
}

func downCreateOrders(tx *sql.Tx) error {
	_, err := tx.Exec("drop table orders")
	return err
}
