package db

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

//go:embed reset.sql
var resetDDL string

func connect(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(ctx context.Context, db *sql.DB) error {
	log.Println("migrating database")
	_, err := db.ExecContext(ctx, ddl)
	if err != nil {
		return err
	}

	return nil
}

func reset(ctx context.Context, db *sql.DB) error {
	log.Println("resetting database")
	_, err := db.ExecContext(ctx, resetDDL)
	if err != nil {
		return err
	}

	return nil
}

func NewInMemoryDB(ctx context.Context) (*sql.DB, error) {
	db, err := connect(":memory:")
	if err != nil {
		return nil, err
	}
	err = migrate(ctx, db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewSQLiteDB(ctx context.Context, path string) (*sql.DB, error) {
	db, err := connect(path)
	if err != nil {
		return nil, err
	}
	err = migrate(ctx, db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewEmptySQLiteDB(ctx context.Context, path string) (*sql.DB, error) {
	db, err := connect(path)
	if err != nil {
		return nil, err
	}
	err = reset(ctx, db)
	if err != nil {
		return nil, err
	}
	err = migrate(ctx, db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
