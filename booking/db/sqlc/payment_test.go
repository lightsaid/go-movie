package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"lightsaid.com/go-movie/booking/utils"
)

func createRandomPayment(t *testing.T) Payment {
	b := createRandomBooking(t)
	arg := CreatePaymentParams{
		BookingID: b.ID,
		Amount:    utils.RandomPrice(),
	}
	p, err := testQueries.CreatePayment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, arg.BookingID, arg.BookingID)
	require.Equal(t, arg.Amount, arg.Amount)

	return p
}

func TestCreatePayment(t *testing.T) {
	_ = createRandomPayment(t)
}

func TestGetPayment(t *testing.T) {
	p := createRandomPayment(t)
	p2, err := testQueries.GetPayment(context.Background(), p.ID)

	require.NoError(t, err)
	require.NotEmpty(t, p2)
	require.Equal(t, p, p2)
}

func TestGetPaymentByBookingID(t *testing.T) {
	p := createRandomPayment(t)
	p2, err := testQueries.GetPaymentByBookingID(context.Background(), p.BookingID)

	require.NoError(t, err)
	require.NotEmpty(t, p2)
	require.Equal(t, p, p2)
}
