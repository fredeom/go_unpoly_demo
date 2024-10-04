package db

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase(dbPath string) (*sql.DB, error) {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	_, err = DB.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS company (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			address TEXT NOT NULL
		)`,
	)
	if err != nil {
		return nil, err
	}
	return DB, nil
}
