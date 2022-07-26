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

func TestUpdateShow(t *testing.T) {
	show := createRandomShow(t)
	m := createRandomMovie(t)
	h := createRandomCinemaHall(t)

	arg := UpdateShowParams{
		ID:           show.ID,
		Date:         time.Now(),
		StartTime:    time.Now(),
		EndTime:      time.Now(),
		CinemaHallID: h.ID,
		MovieID:      m.ID,
	}
	show2, err := testQueries.UpdateShow(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, show2)

	require.Equal(t, arg.ID, show2.ID)
	require.WithinDuration(t, arg.Date, show2.Date, time.Second)
	require.WithinDuration(t, arg.StartTime, show2.StartTime, time.Second)
	require.WithinDuration(t, arg.EndTime, show2.EndTime, time.Second)
	require.Equal(t, arg.CinemaHallID, show2.CinemaHallID)
	require.Equal(t, arg.MovieID, show2.MovieID)
}

func TestGetShow(t *testing.T) {
	show := createRandomShow(t)
	s, err := testQueries.GetShow(context.Background(), show.ID)
	require.NoError(t, err)
	require.NotEmpty(t, s)

	require.Equal(t, show.ID, s.ID)
	require.Equal(t, show.MovieID, s.MovieID)
	require.Equal(t, show.CinemaHallID, s.CinemaHallID)
	require.Equal(t, show.EndTime, s.EndTime)
	require.Equal(t, show.StartTime, s.StartTime)
	require.Equal(t, show.Date, s.Date)
}

func TestGetShowList(t *testing.T) {
	rows := 10
	for i := 0; i < rows; i++ {
		_ = createRandomShow(t)
	}
	arg := GetShowListParams{
		Limit:  5,
		Offset: 10,
	}
	list, err := testQueries.GetShowList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, list)

	for _, ls := range list {
		require.NotEmpty(t, ls)
		require.True(t, ls.ID > 0)
	}
}
