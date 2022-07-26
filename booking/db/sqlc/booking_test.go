package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"lightsaid.com/go-movie/booking/utils"
)

func createRandomBooking(t *testing.T, args ...CreateBookingParams) Booking {
	user := createRandomUser(t)
	show := createRandomShow(t)
	var arg CreateBookingParams
	if len(args) > 0 {
		arg = args[0]
	} else {
		arg = CreateBookingParams{
			UserID:     user.ID,
			ShowID:     show.ID,
			SeatNumber: utils.RandomSeatNumber(),
			CreatedAt:  time.Now(),
			Status:     int32(utils.RandomInt(1, 10)),
		}
	}
	book, err := testQueries.CreateBooking(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)

	require.Equal(t, arg.UserID, book.UserID)
	require.Equal(t, arg.ShowID, book.ShowID)
	require.Equal(t, arg.SeatNumber, book.SeatNumber)
	require.Equal(t, arg.Status, book.Status)
	require.WithinDuration(t, arg.CreatedAt, book.CreatedAt, time.Second)

	return book
}

func TestCreateBooking(t *testing.T) {
	_ = createRandomBooking(t)
}

func TestGetBooking(t *testing.T) {
	b1 := createRandomBooking(t)
	b2, err := testQueries.GetBooking(context.Background(), b1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, b2)
	require.Equal(t, b1, b2)
}

func TestGetBookingByUserID(t *testing.T) {
	book := createRandomBooking(t)
	fmt.Println(">> ", book.UserID)
	arg := CreateBookingParams{
		UserID: book.UserID,
		ShowID: book.ShowID,
	}
	for i := 0; i < 10; i++ {
		arg.SeatNumber = utils.RandomSeatNumber()
		arg.CreatedAt = time.Now()
		arg.Status = int32(utils.RandomInt(1, 10))
		fmt.Println(">> ", book.UserID)
		_ = createRandomBooking(t, arg)
	}
	param := GetBookingByUserIDParams{
		UserID: book.UserID,
		Limit:  5,
		Offset: 5,
	}
	list, err := testQueries.GetBookingByUserID(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, list)

	for _, v := range list {
		require.NotEmpty(t, v)
		require.Equal(t, v.UserID, book.UserID)
		require.Equal(t, v.ShowID, book.ShowID)
	}
}
