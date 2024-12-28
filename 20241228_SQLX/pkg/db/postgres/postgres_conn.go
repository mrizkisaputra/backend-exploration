package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConn() (*sqlx.DB, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/db_belajar?sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("NewPostgresConn.sqlx.Open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("NewPostgresConn.db.Ping: %w", err)
	}

	return db, nil
}
