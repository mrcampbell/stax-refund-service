package main

import (
	"context"

	"github.com/mrcampbell/stax-refund-service/config"
	"github.com/mrcampbell/stax-refund-service/db"
	"github.com/mrcampbell/stax-refund-service/internal/sqlc"
	mock "github.com/mrcampbell/stax-refund-service/internal/sqlc/mocks"
)

func main() {
	ctx := context.Background()

	config.Load()

	conn, err := db.NewEmptySQLiteDB(ctx, "./db/files/stax.db")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queries := sqlc.New(conn)
	mock.PopulateSampleData(ctx, queries)
}
