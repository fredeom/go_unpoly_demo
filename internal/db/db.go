package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Db *sql.DB
}

func NewStore(dbName string) (Store, error) {
	Db, err := getConnection(dbName)
	if err != nil {
		return Store{}, err
	}

	if err := createMigrations(Db); err != nil {
		return Store{}, err
	}

	return Store{
		Db,
	}, nil
}

func getConnection(dbName string) (*sql.DB, error) {
	var (
		err error
		db  *sql.DB
	)

	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("🔥 failed to connect to the database: %s", err)
	}

	log.Println("🚀 Connected Successfully to the Database")

	return db, nil
}

func createMigrations(db *sql.DB) error {
	stmt := `CREATE TABLE IF NOT EXISTS company (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			address TEXT NOT NULL
		);`

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	stmt2 := `CREATE TABLE IF NOT EXISTS project (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			company_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			budget INTEGER NOT NULL
		);`

	_, err2 := db.Exec(stmt2)
	if err2 != nil {
		return err2
	}

	stmt3 := `CREATE TABLE IF NOT EXISTS task (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			done INTEGER NOT NULL
		);`

	_, err3 := db.Exec(stmt3)
	if err3 != nil {
		return err3
	}

	return nil
}
