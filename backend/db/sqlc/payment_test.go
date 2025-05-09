package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomPayment(t *testing.T) Payment {
	booking := createRandomBooking(t) // đảm bảo bạn có hàm này

	amount := pgtype.Numeric{}
	err := amount.Scan("2500000.00")
	require.NoError(t, err)

	arg := CreatePaymentParams{
		Amount:        amount,
		Currency:      pgtype.Text{String: "VND", Valid: true},
		PaymentMethod: pgtype.Text{String: "CreditCard", Valid: true},
		Status:        pgtype.Text{String: "Success", Valid: true},
		BookingID:     pgtype.Text{String: booking.BookingID, Valid: true},
	}

	payment, err := testQueries.CreatePayment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, payment)

	require.Equal(t, arg.Currency.String, payment.Currency.String)
	require.Equal(t, arg.PaymentMethod.String, payment.PaymentMethod.String)
	require.Equal(t, arg.Status.String, payment.Status.String)
	require.Equal(t, arg.BookingID.String, payment.BookingID.String)

	return payment
}

func TestCreatePayment(t *testing.T) {
	createRandomPayment(t)
}

func TestGetPayment(t *testing.T) {
	p1 := createRandomPayment(t)

	p2, err := testQueries.GetPayment(context.Background(), p1.PaymentID)
	require.NoError(t, err)
	require.NotEmpty(t, p2)

	require.Equal(t, p1.PaymentID, p2.PaymentID)
	require.Equal(t, p1.BookingID, p2.BookingID)
}

func TestDeletePayment(t *testing.T) {
	p := createRandomPayment(t)

	err := testQueries.DeletePayment(context.Background(), p.PaymentID)
	require.NoError(t, err)

	_, err = testQueries.GetPayment(context.Background(), p.PaymentID)
	require.Error(t, err)
}

func TestListPayments(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomPayment(t)
	}

	arg := ListPaymentParams{
		Limit:  3,
		Offset: 0,
	}

	payments, err := testQueries.ListPayment(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, payments, 3)

	for _, payment := range payments {
		require.NotEmpty(t, payment)
	}
}
