package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomShow(t *testing.T) Show {
	m := createRandomMovie(t)
	h := createRandomCinemaHall(t)
	arg := CreateShowParams{
		Date:         time.Now(),
		StartTime:    time.Now(),
		EndTime:      time.Now(),
		CinemaHallID: h.ID,
		MovieID:      m.ID,
	}
	show, err := testQueries.CreateShow(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, show)
	require.WithinDuration(t, arg.Date, show.Date, time.Second)
	require.WithinDuration(t, arg.StartTime, show.StartTime, time.Second)
	require.WithinDuration(t, arg.EndTime, show.EndTime, time.Second)
	require.Equal(t, arg.CinemaHallID, show.CinemaHallID)
	require.Equal(t, arg.MovieID, show.MovieID)

	return show
}

func TestCreateShow(t *testing.T) {
	_ = createRandomShow(t)
}
