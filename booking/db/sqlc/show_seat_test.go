package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"lightsaid.com/go-movie/booking/utils"
)

func createRandomShowSeat(t *testing.T) ShowSeat {
	cinemaSeat := createRandomCinemaSeat(t)
	show := createRandomShow(t)
	book := createRandomBooking(t)
	arg := CreateShowSeatParams{
		CinemaSeatID: cinemaSeat.ID,
		ShowID:       show.ID,
		BookingID:    book.ID,
		Status:       int32(utils.RandomInt(1, 10)),
		Price:        utils.RandomPrice(),
	}
	ss, err := testQueries.CreateShowSeat(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ss)

	require.Equal(t, arg.CinemaSeatID, ss.CinemaSeatID)
	require.Equal(t, arg.ShowID, ss.ShowID)
	require.Equal(t, arg.BookingID, ss.BookingID)
	require.Equal(t, arg.Status, ss.Status)
	require.Equal(t, arg.Price, ss.Price)
	return ss
}

func TestCreateShowSeat(t *testing.T) {
	_ = createRandomShowSeat(t)
}

func TestGetShowSeat(t *testing.T) {
	ss1 := createRandomShowSeat(t)
	ss2, err := testQueries.GetShowSeat(context.Background(), ss1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, ss2)
	require.Equal(t, ss1, ss2)
}
