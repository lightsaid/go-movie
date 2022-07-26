package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"lightsaid.com/go-movie/booking/utils"
)

func createRandomCinemaSeat(t *testing.T, halls ...CinemaHall) CinemaSeat {
	var h CinemaHall
	if len(halls) > 0 {
		h = halls[0]
	} else {
		h = createRandomCinemaHall(t)
	}
	arg := CreateCinemaSeatParams{
		CinemaHallID: h.ID,
		Type:         int32(utils.RandomInt(1, 10)),
		SeatNumber:   utils.RandomSeatNumber(),
	}
	c, err := testQueries.CreateCinemaSeat(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, c)
	require.Equal(t, arg.CinemaHallID, c.CinemaHallID)
	require.Equal(t, arg.Type, c.Type)
	require.Equal(t, arg.SeatNumber, c.SeatNumber)
	return c
}

func TestCreateCinemaSeat(t *testing.T) {
	_ = createRandomCinemaSeat(t)
}

func TestGetCinemaSeatByHallID(t *testing.T) {
	h := createRandomCinemaHall(t)
	cs := createRandomCinemaSeat(t, h)
	testQueries.GetCinemaSeatByHallID(context.Background(), h.ID)

	require.NotEmpty(t, cs)
	require.Equal(t, h.ID, cs.CinemaHallID)
}

func TestUpdateCinemaSeat(t *testing.T) {
	seat := createRandomCinemaSeat(t)
	arg := UpdateCinemaSeatParams{
		ID:         seat.ID,
		Type:       int32(utils.RandomInt(1, 10)),
		SeatNumber: utils.RandomSeatNumber(),
	}
	seat2, err := testQueries.UpdateCinemaSeat(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, seat2)

	require.Equal(t, seat.ID, seat2.ID)
	require.Equal(t, arg.Type, seat2.Type)
	require.Equal(t, arg.SeatNumber, seat2.SeatNumber)
}
