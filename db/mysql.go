package db

import "github.com/jmoiron/sqlx"

func CreateTable(db *sqlx.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS cities (
			id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(30) NOT NULL,
			state VARCHAR(30) NOT NULL
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
