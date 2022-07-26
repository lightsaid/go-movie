package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"lightsaid.com/go-movie/booking/utils"
)

func randomMovieParams() []byte {
	arg := Movie{
		Title:     utils.RandomString(8),
		ReleaseAt: time.Now(),
		CoverUrl:  fmt.Sprintf("%s.png", utils.RandomString(10)),
		Duration:  utils.NullInt32(utils.RandomInt(60, 300)).Int32(),
		Language:  utils.NullStr(utils.RandomLanguage()).String(),
		Genre:     utils.NullStr(utils.RandomGenres()).String(),
		Rating:    sql.NullFloat64{Float64: 4.5, Valid: true},
		Director:  utils.NullStr(utils.RandomString(4)).String(),
		Desc:      utils.NullStr(utils.RandomString(20)).String(),
		Status:    int32(utils.RandomInt(1, 10)),
		Star:      utils.NullStr(utils.RandomString(4)).String(),
		WishCount: utils.NullInt32(utils.RandomInt(0, 10000)).Int32(),
	}
	bye, _ := json.Marshal(&arg)
	return bye
}

func createRandomMovie(t *testing.T, args ...CreateMovieParams) Movie {
	var arg CreateMovieParams
	if len(args) > 0 {
		arg = args[0]
	} else {
		bye := randomMovieParams()
		err := json.Unmarshal(bye, &arg)
		require.NoError(t, err)
	}

	m, err := testQueries.CreateMovie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, m)

	require.Equal(t, arg.Title, m.Title)
	fmt.Println(">> ", arg.ReleaseAt)
	fmt.Println(">> ", m.ReleaseAt)
	require.WithinDuration(t, arg.ReleaseAt, m.ReleaseAt, time.Second)
	require.Equal(t, arg.CoverUrl, m.CoverUrl)
	require.Equal(t, arg.Duration, m.Duration)
	require.Equal(t, arg.Language, m.Language)
	require.Equal(t, arg.Genre, m.Genre)
	require.Equal(t, arg.Rating, m.Rating)
	require.Equal(t, arg.Director, m.Director)
	require.Equal(t, arg.Desc, m.Desc)
	require.Equal(t, arg.Star, m.Star)
	require.Equal(t, arg.Status, m.Status)
	require.Equal(t, arg.WishCount, m.WishCount)

	return m
}

func TestCreateMoive(t *testing.T) {
	_ = createRandomMovie(t)
}

func TestUpdateMovie(t *testing.T) {
	m := createRandomMovie(t)
	bye := randomMovieParams()
	var arg UpdateMoiveParams
	err := json.Unmarshal(bye, &arg)
	arg.ID = m.ID

	require.NoError(t, err)

	m2, err := testQueries.UpdateMoive(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, m2)
	require.Equal(t, m.ID, m2.ID)
	require.Equal(t, arg.CoverUrl, m2.CoverUrl)
	require.Equal(t, arg.Duration, m2.Duration)
	require.Equal(t, arg.Language, m2.Language)
	require.Equal(t, arg.Genre, m2.Genre)
	require.Equal(t, arg.Rating, m2.Rating)
	require.Equal(t, arg.Director, m2.Director)
	require.Equal(t, arg.Desc, m2.Desc)
	require.Equal(t, arg.Star, m2.Star)
	require.Equal(t, arg.Status, m2.Status)
	require.Equal(t, arg.WishCount, m2.WishCount)
}

func TestGetMovie(t *testing.T) {
	m := createRandomMovie(t)
	m2, err := testQueries.GetMovie(context.Background(), m.ID)
	require.NoError(t, err)
	require.NotEmpty(t, m2)

	require.Equal(t, m.ID, m2.ID)
	require.Equal(t, m.CoverUrl, m2.CoverUrl)
	require.Equal(t, m.Duration, m2.Duration)
	require.Equal(t, m.Language, m2.Language)
	require.Equal(t, m.Genre, m2.Genre)
	require.Equal(t, m.Rating, m2.Rating)
	require.Equal(t, m.Director, m2.Director)
	require.Equal(t, m.Desc, m2.Desc)
	require.Equal(t, m.Star, m2.Star)
	require.Equal(t, m.Status, m2.Status)
	require.Equal(t, m.WishCount, m2.WishCount)

}

func TestGetMovies(t *testing.T) {
	for i := 0; i < 10; i++ {
		_ = createRandomMovie(t)
	}
	arg := GetMoviesParams{
		Limit:  5,
		Offset: 2,
	}
	ms, err := testQueries.GetMovies(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ms)
	for _, v := range ms {
		require.NotEmpty(t, v)
	}
}

func TestGetMoviesByStatus(t *testing.T) {
	status := int32(utils.RandomInt(1, 10))

	for i := 0; i < 10; i++ {
		byte := randomMovieParams()
		var arg CreateMovieParams
		json.Unmarshal(byte, &arg)
		arg.Status = status
		_ = createRandomMovie(t, arg)
	}

	var count int32 = 5
	arg := GetMoviesByStatusParams{
		Status: status,
		Limit:  count,
		Offset: 5,
	}
	list, err := testQueries.GetMoviesByStatus(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, list)

	require.True(t, len(list) >= int(count))
	for _, v := range list {
		require.NotEmpty(t, v)
	}

}
