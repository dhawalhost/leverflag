package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// DataAccess defines the interface for data access operations.
type DataAccess interface {
	Exec(query string, args ...any) (sql.Result, error)
	QueryRow(query string, args ...any) *sqlx.Row
	Query(query string, args ...any) (*sqlx.Rows, error)
}

// InitDB initializes the SQLite database and returns a sqlx DB instance.
func InitDB(dbPath string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create user table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        username TEXT NOT NULL,
        email TEXT NOT NULL
    )`)

	if err != nil {
		return nil, err
	}
	return db, nil
}
