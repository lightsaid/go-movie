package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"lightsaid.com/go-movie/booking/utils"
)

func createRandomCinemaHall(t *testing.T) CinemaHall {
	c := createRandomCinema(t)
	arg := CreateCinemaHallParams{
		CinemaID:   c.ID,
		Name:       utils.RandomString(6),
		TotalSeats: int32(utils.RandomInt(100, 1000)),
	}
	h, err := testQueries.CreateCinemaHall(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, h)

	require.Equal(t, arg.CinemaID, h.CinemaID)
	require.Equal(t, arg.Name, h.Name)
	require.Equal(t, arg.TotalSeats, h.TotalSeats)

	return h
}

func TestCreateCinemaHall(t *testing.T) {
	_ = createRandomCinemaHall(t)
}

func TestUpdateCinemaHall(t *testing.T) {
	c := createRandomCinema(t)
	h := createRandomCinemaHall(t)

	arg := UpdateCinemaHallParams{
		ID:         h.ID,
		CinemaID:   c.ID,
		Name:       utils.RandomString(6),
		TotalSeats: int32(utils.RandomInt(100, 1000)),
	}

	h2, err := testQueries.UpdateCinemaHall(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, h2)
	require.Equal(t, h.ID, h2.ID)
	require.Equal(t, arg.CinemaID, h2.CinemaID)
	require.Equal(t, arg.Name, h2.Name)
	require.Equal(t, arg.TotalSeats, h2.TotalSeats)
}

func TestListCinemaHall(t *testing.T) {
	rows := 10
	for i := 0; i < rows; i++ {
		_ = createRandomCinemaHall(t)
	}
	arg := GetCinemaHallListParams{
		Limit:  5,
		Offset: 5,
	}
	list, err := testQueries.GetCinemaHallList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, list)
	for _, v := range list {
		require.NotEmpty(t, v)
	}
}
