package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upPasswordsTable, downPasswordsTable)
}

func upPasswordsTable(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`CREATE TABLE passwords(
		id char(14) NOT NULL,
		user_id char(14) NOT NULL,
		website varchar(200) NOT NULL,
		email varchar(100) NOT NULL,
		password varchar(100) NOT NULL,
		created_at bigint NOT NULL,
		updated_at bigint NOT NULL,
		deleted_at bigint DEFAULT NULL,
		PRIMARY KEY (id)
	);`)

	if err != nil {
		return err
	}
	return nil
}

func downPasswordsTable(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE IF EXISTS passwords;`)
	if err != nil {
		return err
	}
	return nil
}
