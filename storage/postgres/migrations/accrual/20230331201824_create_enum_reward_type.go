package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateEnumRewardType, downCreateEnumRewardType)
}

func upCreateEnumRewardType(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TYPE reward_type AS ENUM ('%', 'pt');
	`)
	return err
}

func downCreateEnumRewardType(tx *sql.Tx) error {
	_, err := tx.Exec("drop type reward_type")
	return err
}
