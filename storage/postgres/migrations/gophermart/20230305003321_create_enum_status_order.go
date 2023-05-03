package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateEnumStatusOrder, downCreateEnumStatusOrder)
}

func upCreateEnumStatusOrder(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TYPE status AS ENUM ('NEW', 'REGISTERED', 'INVALID', 'PROCESSING', 'PROCESSED');
	`)
	return err
}

func downCreateEnumStatusOrder(tx *sql.Tx) error {
	_, err := tx.Exec("drop type status")
	return err
}
