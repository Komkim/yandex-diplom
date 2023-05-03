package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateStatus, downCreateStatus)
}

func upCreateStatus(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TYPE status AS ENUM ('REGISTERED', 'INVALID', 'PROCESSING', 'PROCESSED');
	`)
	return err
}

func downCreateStatus(tx *sql.Tx) error {
	_, err := tx.Exec("drop type status")
	return err
}
