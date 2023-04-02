package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateConnectGoodsAndOrders, downCreateConnectGoodsAndOrders)
}

func upCreateConnectGoodsAndOrders(tx *sql.Tx) error {
	_, err := tx.Exec(`
			create table connect
			(
				id          uuid default gen_random_uuid() not null primary key,					
				number	    bigint not null,
				goods_id    uuid REFERENCES goods (id), 				
				create_at   timestamp with time zone default current_timestamp
			);
	`)
	return err
}

func downCreateConnectGoodsAndOrders(tx *sql.Tx) error {
	_, err := tx.Exec("drop table connect")
	return err
}
