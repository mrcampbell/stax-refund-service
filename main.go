package main

import (
	"context"

	"github.com/mrcampbell/stax-refund-service/app"
	"github.com/mrcampbell/stax-refund-service/config"
	"github.com/mrcampbell/stax-refund-service/db"
	routes "github.com/mrcampbell/stax-refund-service/internal/http/routers"
	"github.com/mrcampbell/stax-refund-service/internal/mock"
	"github.com/mrcampbell/stax-refund-service/internal/sqlc"
	"github.com/mrcampbell/stax-refund-service/internal/sqlite"
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

	mockAuthService := mock.NewAuthServiceWithMockedMethods()
	paymentClient := sqlite.NewPaymentClient(queries)
	refundService := sqlite.NewRefundService(queries)

	serverResources := app.ServerResources{
		Queries:       queries,
		AuthService:   mockAuthService,
		PaymentClient: paymentClient,
		RefundService: refundService,
	}

	server := routes.NewServer(queries, serverResources)
	server.Run()
}
