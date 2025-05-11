package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	BookingTx(ctx context.Context, arg BookingTxParams, passengers []PassengerParams, payment PaymentParams) (BookingTxResult, error)
	CreateFlightWithSeats(ctx context.Context, params CreateFlightWithSeatsParams) (Flight, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	// Tuong tu
	connPool *pgxpool.Pool
	// Tranh copy toan bo struct Queries von chua nhieu du lieu. Dung 1 instance => it ro ri bo nho. Embeded struct truy cap truc tiep cac method duoc
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
